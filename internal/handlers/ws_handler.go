package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"main/internal/errors"
	"main/internal/models"
	"main/internal/mqtt"
	"main/internal/services"
	"sync"

	"net/http"
	"strings"

	"github.com/gorilla/websocket"
)

var (
	clients      = make(map[*websocket.Conn]bool) // mapa global de clientes
	clientsMutex sync.Mutex                       // bloqueia a variavel para garantir a manipulação segura
	once         sync.Once                        // garante que o broadcast só seja iniciado uma vez
)

func WsHandler(w http.ResponseWriter, r *http.Request) {

	var upgrader = websocket.Upgrader{
		ReadBufferSize:  2048,
		WriteBufferSize: 2048,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}

	// err := services.VerifyToken(r)
	// if err != nil {
	// 	http.Error(w, "Unauthorized: Invalid token", http.StatusUnauthorized)
	// 	return
	// }

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()

	ws := services.NewWsService()

	clientsMutex.Lock()
	clients[conn] = true
	clientsMutex.Unlock()

	go once.Do(func() { mqtt.BroadcastTime(clients, &clientsMutex) })

	for {
		_, message, err := conn.ReadMessage()
		//start := time.Now().Format("15:04:05.000")
		//fmt.Println("start:xxxxxxxxxxxxxxxx ", start)

		if err != nil {
			log.Println(err)
			clientsMutex.Lock()
			delete(clients, conn)
			clientsMutex.Unlock()
			return
		}

		messageStr := strings.ReplaceAll(string(message), "“", "\"")
		messageStr = strings.ReplaceAll(messageStr, "”", "\"")

		fmt.Printf("Received message: %s\n", messageStr)

		var request models.JsonRpcRequest
		if err := json.Unmarshal([]byte(messageStr), &request); err != nil {
			log.Printf("Error decoding JSON-RPC: %v\n", err)
			return
		}
		res, err := ws.HandleRequest(request.Method, request.Params, request.ID)

		var response models.JsonRpcResponse
		if err != nil {
			response = models.JsonRpcResponse{
				Result: nil,
				Error:  &models.Error{Code: errors.ErrorConvertingJSON},
				ID:     request.ID,
			}
		} else {
			response = models.JsonRpcResponse{
				Result: res,
				Error:  nil,
				ID:     request.ID,
			}
		}

		jsonResponse, err := json.Marshal(response)
		fmt.Println("panelModel: ", string(jsonResponse))
		if err != nil {
			log.Println("Error encoding response JSON:", err)
			return
		}

		err = conn.WriteMessage(websocket.TextMessage, jsonResponse)
		if err != nil {
			log.Println(err)
			return
		}

	}
}
