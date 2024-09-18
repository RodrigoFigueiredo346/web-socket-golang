package main

import (
	"fmt"
	"log"
	"main/internal/handler"
	"main/internal/mqtt"
	"main/internal/services"

	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	go mqtt.InitMqtt()
	services.Config()

	cfg := services.GetConfig()

	r := mux.NewRouter()
	r.Use(enableCORS)
	r.HandleFunc("/ws", handler.WsHandler)
	r.HandleFunc("/api/sub", handler.ApiSub).Methods(http.MethodPost, http.MethodOptions)
	r.HandleFunc("/api/login", handler.Login).Methods(http.MethodPost, http.MethodOptions)
	r.HandleFunc("/api/create-user", handler.CreateUser).Methods(http.MethodPost, http.MethodOptions)
	r.HandleFunc("/api/create-panel", handler.CreatePanel).Methods(http.MethodPost, http.MethodOptions)

	fmt.Printf("WebSocket running in ws://0.0.0.0:%s/ws\n", cfg.Port)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", cfg.Port), r))
}

// Middleware para habilitar CORS
func enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*") // Permitir todas as origens
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		// Se for uma requisição OPTIONS (preflight), responde diretamente
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}
