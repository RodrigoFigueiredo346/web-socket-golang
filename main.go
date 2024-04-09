package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/websocket"
)

var (
	clients   = make(map[*websocket.Conn]bool) // mapa para armazenar conexões de clientes
	broadcast = make(chan Message)             // canal de transmissão de mensagens
	upgrader  = websocket.Upgrader{}
)

// Message estrutura para representar uma mensagem
type Message struct {
	Sender    string `json:"sender"`
	Recipient string `json:"recipient"`
	Content   string `json:"content"`
}

func main() {
	// Configurar rota para lidar com atualizações de WebSocket
	http.HandleFunc("/ws", handleConnections)

	// Iniciar goroutine para manipular mensagens recebidas e enviá-las para o canal de transmissão
	go handleMessages()

	// Configurar servidor para servir arquivos estáticos e iniciar servidor HTTP
	fs := http.FileServer(http.Dir("html"))
	http.Handle("/", fs)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Println("Servidor iniciado na porta", port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal("Erro ao iniciar servidor: ", err)
	}
}

// handleConnections lida com as conexões WebSocket dos clientes
func handleConnections(w http.ResponseWriter, r *http.Request) {
	// Upgrade da conexão HTTP para WebSocket
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}
	// Garantir que a conexão seja fechada ao sair da função
	defer conn.Close()

	// Adicionar nova conexão ao mapa de clientes
	clients[conn] = true

	// Loop para ler mensagens do cliente
	for {
		var msg Message
		// Ler a mensagem JSON e armazenar em msg
		err := conn.ReadJSON(&msg)
		if err != nil {
			log.Printf("Erro ao ler mensagem: %v", err)
			delete(clients, conn)
			break
		}
		// Enviar a mensagem para o canal de transmissão
		broadcast <- msg
	}
}

// handleMessages transmite mensagens recebidas para todos os clientes conectados
func handleMessages() {
	for {
		// Receber a próxima mensagem do canal de transmissão
		msg := <-broadcast
		// Enviar mensagem para todos os clientes
		for client := range clients {
			err := client.WriteJSON(msg)
			if err != nil {
				log.Printf("Erro ao enviar mensagem para o cliente: %v", err)
				client.Close()
				delete(clients, client)
			}
		}
	}
}
