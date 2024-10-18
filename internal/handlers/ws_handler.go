package handlers

import (
	"encoding/json"
	"fmt"
	"log"
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

		messageStr := string(message)
		messageStr = strings.ReplaceAll(messageStr, "“", "\"")
		messageStr = strings.ReplaceAll(messageStr, "”", "\"")

		fmt.Printf("Received message: %s\n", messageStr)

		var request models.JsonRpcRequest

		if err := json.Unmarshal([]byte(messageStr), &request); err != nil {
			log.Printf("Error decoding JSON-RPC: %v\n", err)
			conn.WriteMessage(websocket.TextMessage, []byte("Invalid request..."))
			continue
		}
		var response string
		res, err := ws.HandleRequest(request.Method, request.Params, request.ID)
		if err != nil {
			log.Println(err)
			response = fmt.Sprintf(`{"error":"%v","id":%d}`, err, request.ID)

		} else {
			response = fmt.Sprintf(`{"result":"%v","id":%d}`, res, request.ID)
		}

		err = conn.WriteMessage(websocket.TextMessage, []byte(response))

		if err != nil {
			log.Println(err)
			return
		}

	}
}
