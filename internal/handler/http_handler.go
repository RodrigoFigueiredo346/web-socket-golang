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

func ApiSub(w http.ResponseWriter, r *http.Request) {

	contentType := r.Header.Get("Content-Type")

	if !strings.HasPrefix(contentType, "application/json") { //TODO splitar
		http.Error(w, "Content-Type must be 'application/json'", http.StatusUnsupportedMediaType)
		return
	}

	var requestModel models.JsonRpcRequest

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Cannot read body", http.StatusBadRequest)
		return
	}

	err = json.Unmarshal(body, &requestModel)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	paramsMap := requestModel.Params.(map[string]interface{})

	n1, ok1 := paramsMap["n1"].(float64)
	n2, ok2 := paramsMap["n2"].(float64)
	if !ok1 || !ok2 {
		http.Error(w, "Missing or invalid parameters", http.StatusBadRequest)
		return
	}

	result := n1 - n2

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf(`{"result": %.2f, "id": %d}`, result, requestModel.ID)))
}

func CreateUser(w http.ResponseWriter, r *http.Request) {

	var requestModel models.JsonRpcRequest

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Cannot read body", http.StatusBadRequest)
		return
	}

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

	_, ok := cache.Get(string(user))
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

	//TODO - apesar de ser possível validar, o "method" perde sentido por ser REST

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

	cache := services.GetCache()

	userReq := paramsReq["user"]
	passReq := paramsReq["password"]
	userGet := fmt.Sprintf("%v", userReq)
	passGet := fmt.Sprintf("%v", passReq)

	userCache, ok := cache.Get(userGet)
	if !ok {
		http.Error(w, "User not registered", http.StatusInternalServerError)
		return
	}

	userGet = fmt.Sprintf("%v", userCache)
	userId := strings.Split(userGet, " ")[2]
	passCache := strings.Split(userGet, " ")[1]
	if passCache != passGet {
		http.Error(w, "Verify your credencials", http.StatusInternalServerError)
		return
	}

	token, err := services.CreateToken(userId)
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

func CreatePanel(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "aplication/json")

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

	parameters := requestModel.Params.(map[string]interface{})
	messages := parameters["messages"].([]interface{})
	name := parameters["name"].(string)
	brightMode := parameters["bright_mode"].(float64)
	var msgs []string
	for _, msg := range messages {
		msgs = append(msgs, fmt.Sprintf("%v", msg))
	}
	panelModel := models.PanelModel{
		ID:         123,
		Name:       name,
		Messages:   msgs,
		BrightMode: int(brightMode),
	}
	cache := services.GetCache()

	cache.Set(string(panelModel.ID), panelModel)
	_, ok := cache.Get(string(panelModel.ID))
	if !ok {
		http.Error(w, "Panel not registered", http.StatusInternalServerError)
		return
	}

	var response models.JsonRpcResponse

	response.Result = "Panel register succesfully"
	response.ID = requestModel.ID

	responseJSON, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Error Marshal userResponse", http.StatusInternalServerError)
		return
	}
	fmt.Println(fmt.Sprintf("Panel %d register succesfully", string(panelModel.ID)))

	w.Write([]byte(responseJSON))
}
