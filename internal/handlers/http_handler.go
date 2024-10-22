package handlers

import (
	"context"
	"database/sql"
	"encoding/json"
	"log"

	"fmt"
	"io"
	"main/internal/db"
	"main/internal/models"
	"main/internal/services"
	"main/sqlc"
	"math/rand"
	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {

	var requestModel models.JsonRpcRequest

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Cannot read body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	err = json.Unmarshal(body, &requestModel)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	paramsMap := requestModel.Params.(map[string]interface{})

	user, ok1 := paramsMap["user"].(string)
	password, ok2 := paramsMap["password"].(string)
	if !ok1 || !ok2 {
		http.Error(w, "Missing or invalid parameters", http.StatusBadRequest)
		return
	}

	cache := services.GetCache()

	_, ok := cache.Get(user)
	if ok {
		http.Error(w, "User already registered ", http.StatusInternalServerError)
		return
	}

	userCreated := models.User{
		Username: user,
		Password: password,
		ID:       rand.Intn(8999998) + 1000000,
	}

	cache.Set(string(userCreated.Username), userCreated)
	_, ok = cache.Get(string(userCreated.Username))
	if !ok {
		http.Error(w, "Error setting user in cache", http.StatusInternalServerError)
		return
	}

	var userResponse models.JsonRpcResponse
	userResponse.Result = "User create successfully"
	userResponse.ID = requestModel.ID

	userResponseJSON, err := json.Marshal(userResponse)
	if err != nil {
		http.Error(w, "Error Marshal userResponse", http.StatusInternalServerError)
		return
	}
	w.Write([]byte(userResponseJSON))
}

func Login(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "aplication/json")
	ctx := context.Background()

	//todo controlar numero de tentativas de cadastro por IP????
	// 	ip := r.Header.Get("X-Forwarded-For")
	// 	ip1 := r.RemoteAddr

	var requestModel models.JsonRpcRequest

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Cannot read body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	err = json.Unmarshal(body, &requestModel)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	paramsReq := requestModel.Params.(map[string]interface{})

	// transformando interface em string
	userGet := fmt.Sprintf("%v", paramsReq["user"])
	passGet := fmt.Sprintf("%v", paramsReq["password"])

	dt := sqlc.New(db.DB)

	user, err := dt.GetUserByLoginAndPassword(ctx, sqlc.GetUserByLoginAndPasswordParams{
		Login: userGet,
		Pass:  passGet,
	})

	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("Login failed: user not found or incorrect password")
		} else {
			log.Fatal(err)
		}
	}

	// crio token

	//devolvo token na resposta

	/*
		cache := services.GetCache()
		userCache, ok := cache.Get(userGet)
		if !ok {
			http.Error(w, "User not registered", http.StatusNotFound)
			return
		}
		userGet = fmt.Sprintf("%v", userCache)
		fmt.Println(userGet)
		userId := strings.Split(userGet, " ")[2]
		passCache := strings.Split(userGet, " ")[1]
		if passCache != passGet {
			http.Error(w, "Verify your credencials", http.StatusForbidden)
			return
		}
	*/

	token, err := services.CreateToken(user.Iduser)

	if err != nil {
		http.Error(w, fmt.Sprintf("Error creating token: %v", err), http.StatusInternalServerError)
		return
	}
	var response models.JsonRpcResponse

	response.Result = token
	response.ID = requestModel.ID

	responseJSON, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Error Marshal userResponse", http.StatusInternalServerError)
		return
	}

	w.Write([]byte(responseJSON))
}

// func GetPanels(w http.ResponseWriter, r *http.Request) {

// 	w.Header().Set("Content-Type", "application/json")

// 	cache := services.GetCache()
// 	allCache := cache.GetAll()

// 	response := models.JsonRpcResponse{
// 		Result: allCache,
// 		ID:     999999,
// 	}

// 	responseJSON, err := json.Marshal(response)
// 	if err != nil {
// 		http.Error(w, "Error marshalling panels", http.StatusInternalServerError)
// 		return
// 	}
// 	w.Write(responseJSON)

// }
