package handlers

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"

	"io"
	"main/internal/db"
	"main/internal/errors"
	"main/internal/models"
	"main/internal/services"
	"main/sqlc"
	"net/http"
	"net/mail"
)

func isValidEmail(email string) models.Error {
	_, err := mail.ParseAddress(email)
	if err != nil {
		return models.Error{
			Code: errors.InvalidEmailFormat,
		}
	}

	return models.Error{
		Code: errors.NoError,
	}
}

func User(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Cannot read body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	var requestModel models.JsonRpcRequest
	err = json.Unmarshal(body, &requestModel)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	response := processJsonRpc(requestModel)

	responseJSON, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Error parsing response", http.StatusInternalServerError)
	}

	w.Write([]byte(responseJSON))
}

func processJsonRpc(request models.JsonRpcRequest) models.JsonRpcResponse {
	method := request.Method
	id := request.ID

	switch method {
	case "login":
		paramsList, ok := request.Params.([]interface{})
		fmt.Println(paramsList)
		if !ok || len(paramsList) == 0 {

			return models.JsonRpcResponse{
				Result: nil,
				Error: &models.Error{
					Code: errors.EmptyParameters,
				},
				ID: id,
			}
		}

		paramsData, ok := paramsList[0].(map[string]interface{})
		if !ok {
			return models.JsonRpcResponse{
				Result: nil,
				Error: &models.Error{
					Code: errors.InvalidRequestFormat,
				},
				ID: id,
			}
		}

		if len(paramsData) != 2 {
			return models.JsonRpcResponse{
				Result: nil,
				Error: &models.Error{
					Code: errors.InvalidCredentials,
				},
			}
		}

		userEmail := paramsData["email"].(string)
		emailErr := isValidEmail(userEmail)
		if emailErr.Code != 0 {
			return models.JsonRpcResponse{
				Result: nil,
				Error:  &emailErr,
				ID:     id,
			}
		}

		userPass := paramsData["password"].(string)

		userId, err := readUserData(userEmail, userPass)
		if err.Code != 0 {
			return models.JsonRpcResponse{
				Result: nil,
				Error:  &err,
				ID:     id,
			}
		}

		token, tokenErr := services.CreateToken(userId)
		if tokenErr != nil {
			return models.JsonRpcResponse{
				Result: nil,
				Error: &models.Error{
					Code: errors.InternalServerError,
				},
				ID: id,
			}
		}

		return models.JsonRpcResponse{
			Result: token,
			Error:  nil,
			ID:     id,
		}

	default:
		return models.JsonRpcResponse{
			Result: nil,
			Error: &models.Error{
				Code: errors.InvalidMethod,
			},
			ID: id,
		}

	}
}

func readUserData(email string, password string) (string, models.Error) {

	dt := sqlc.New(db.DB)
	ctx := context.Background()

	user, err := dt.GetUserByLoginAndPassword(ctx, sqlc.GetUserByLoginAndPasswordParams{
		Login: email,
		Pass:  password,
	})

	if err != nil {
		if err == sql.ErrNoRows {
			return "", models.Error{
				Code: errors.UserNotFound,
			}
		} else {
			return "", models.Error{
				Code: errors.DatabaseConnectionError,
			}
		}
	}
	return user.Iduser, models.Error{
		Code: errors.NoError,
	}
}
