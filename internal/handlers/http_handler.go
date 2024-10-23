package handlers

import (
	"context"
	"database/sql"
	"encoding/json"

	"io"
	"main/internal/db"
	"main/internal/models"
	"main/internal/services"
	"main/sqlc"
	"net/http"
)

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
		if !ok || len(paramsList) == 0 {
			return models.JsonRpcResponse{
				Result: nil,
				Error: &models.Error{
					Code: -109,
				},
				ID: id,
			}
		}

		paramsData, ok := paramsList[0].(map[string]interface{})
		if !ok {
			return models.JsonRpcResponse{
				Result: nil,
				Error:  &models.Error{Code: -505},
				ID:     id,
			}
		}

		userId, err := userLogin(paramsData["email"].(string), paramsData["password"].(string))
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
				Error:  &models.Error{Code: -505},
				ID:     id,
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
			Error:  &models.Error{Code: -505},
			ID:     id,
		}

	}
}

func userLogin(email string, password string) (string, models.Error) {

	dt := sqlc.New(db.DB)
	ctx := context.Background()

	user, err := dt.GetUserByLoginAndPassword(ctx, sqlc.GetUserByLoginAndPasswordParams{
		Login: email,
		Pass:  password,
	})

	if err != nil {
		if err == sql.ErrNoRows {
			return "", models.Error{
				Code: -109,
			}
		} else {
			return "", models.Error{
				Code: -505,
			}
		}
	}
	return user.Iduser, models.Error{
		Code: 0,
	}
}

func createUser() int {
	return 0
}
