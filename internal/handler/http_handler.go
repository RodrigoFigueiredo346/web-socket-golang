package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"main/internal/models"
	"main/internal/services"
	"math/rand"
	"net/http"
	"strings"
)

type User struct {
	Username string
	Password string
	ID       int
}

func ApiSub(w http.ResponseWriter, r *http.Request) {

	contentType := r.Header.Get("Content-Type")

	if !strings.HasPrefix(contentType, "application/json") { //TODO splitar
		http.Error(w, "Content-Type must be 'application/json'", http.StatusUnsupportedMediaType)
		return
	}

	var reqBody models.JsonRpcRequest

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Cannot read body", http.StatusBadRequest)
		return
	}

	err = json.Unmarshal(body, &reqBody)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	paramsMap := reqBody.Params.(map[string]interface{})

	n1, ok1 := paramsMap["n1"].(float64)
	n2, ok2 := paramsMap["n2"].(float64)
	if !ok1 || !ok2 {
		http.Error(w, "Missing or invalid parameters", http.StatusBadRequest)
		return
	}

	result := n1 - n2

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf(`{"result": %.2f, "id": %d}`, result, reqBody.ID)))
}

func CreateUser(w http.ResponseWriter, r *http.Request) {

	var reqBody models.JsonRpcRequest

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Cannot read body", http.StatusBadRequest)
		return
	}

	err = json.Unmarshal(body, &reqBody)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	paramsMap := reqBody.Params.(map[string]interface{})

	user, ok1 := paramsMap["user"].(string)
	password, ok2 := paramsMap["password"].(string)
	if !ok1 || !ok2 {
		http.Error(w, "Missing or invalid parameters", http.StatusBadRequest)
		return
	}

	userCreated := User{
		Username: user,
		Password: password,
		ID:       rand.Intn(8999998) + 1000000,
	}

	c := services.NewCache()
	c.Set(string(userCreated.ID), userCreated)

	_, ok := c.Get(string(userCreated.ID))
	if !ok {
		http.Error(w, "Error getting ID cache", http.StatusInternalServerError)
		return
	}

	// userJSON, err := json.Marshal(u)
	// fmt.Println(string(userJSON))
	// if err != nil {
	// 	http.Error(w, "Error converting user to JSON", http.StatusInternalServerError)
	// 	return
	// }

	token, err := services.CreateToken(string(userCreated.ID))
	if err != nil {
		http.Error(w, fmt.Sprintf("Error creating token: %v", err), http.StatusInternalServerError)
		return
	}

	var userResponse models.JsonRpcResponse
	userResponse.Result = token
	userResponse.ID = reqBody.ID

	userResponseJSON, err := json.Marshal(userResponse)
	if err != nil {
		http.Error(w, "Error Marshal userResponse", http.StatusInternalServerError)
		return
	}
	w.Write([]byte(userResponseJSON))
}

func Login(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Login..."))
}
