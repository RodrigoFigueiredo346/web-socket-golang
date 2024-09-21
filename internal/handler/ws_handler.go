package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"main/internal/models"
	"main/internal/mqtt"
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

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()

	clientsMutex.Lock()
	clients[conn] = true
	clientsMutex.Unlock()

	go once.Do(func() { mqtt.BroadcastTime(clients, &clientsMutex) })

	for {
		_, message, err := conn.ReadMessage()
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
			conn.WriteMessage(websocket.TextMessage, []byte("Invalid request"))
			continue
		}

		if request.Method == "sum" {
			// topic := "server"
			// payload := request.Params.(string)
			// mqtt.Publish(topic, payload)

			sumParams, ok := request.Params.(map[string]interface{})
			if !ok {
				log.Println("Error converting Params para map[string]interface{}")
				return
			}

			n1, ok1 := sumParams["n1"].(float64) // JSON decodifica números como float64
			n2, ok2 := sumParams["n2"].(float64)

			if !ok1 || !ok2 {
				log.Println("Error converting parameters to integers")
				return
			}

			sum := n1 + n2

			if err := conn.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf("%2.f", sum))); err != nil {
				log.Println(err)
				return
			}

		}

	}
}
