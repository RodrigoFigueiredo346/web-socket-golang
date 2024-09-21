package mqtt

import (
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/gorilla/websocket"
	"golang.org/x/exp/rand"
)

func BroadcastTime(clients map[*websocket.Conn]bool, clientsMutex *sync.Mutex) {
	fmt.Println("Starting broadcast...")
	id := 10000 + rand.Intn(99999)
	ticker := time.NewTicker(1 * time.Second) // Envia a cada 1 segundo
	defer ticker.Stop()

	for {
		<-ticker.C

		now := time.Now().Format("15:04:05")
		response := fmt.Sprintf(`{"result":"%s","id":%d}`, now, id)

		// Trava o mutex para garantir acesso seguro ao mapa de clientes
		clientsMutex.Lock()
		for client := range clients {
			err := client.WriteMessage(websocket.TextMessage, []byte(response))
			if err != nil {
				log.Printf("Error publishing message: %v\n", err)
				client.Close()
				delete(clients, client) // Remove cliente desconectado
			}
		}
		clientsMutex.Unlock()
	}
}
