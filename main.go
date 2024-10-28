package main

import (
	"fmt"
	"log"
	"main/internal/config"
	"main/internal/db"
	"main/internal/mqtt"
	"main/internal/router"
	"net/http"
)

func main() {

	// Goro()

	config.Config()

	go mqtt.InitMqtt()

	db.InitDB()

	cfg := config.GetConfig()

	r := router.GeneratRoutes()

	fmt.Printf("WebSocket running in ws://0.0.0.0:%s/ws\n", cfg.Port)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", cfg.Port), r))

}

// package main

// import (
// 	"fmt"
// 	"net/http"

// 	"github.com/gorilla/mux"
// )

// // Middleware para habilitar CORS
// func enableCORS(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		// Permitir todas as origens
// 		w.Header().Set("Access-Control-Allow-Origin", "*")
// 		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
// 		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

// 		// Se for uma requisição OPTIONS (preflight), responde diretamente
// 		if r.Method == http.MethodOptions {
// 			w.WriteHeader(http.StatusNoContent) // Retorna 204 No Content
// 			return
// 		}

// 		next.ServeHTTP(w, r)
// 	})
// }

// func main() {
// 	r := mux.NewRouter()

// 	// Usar o middleware CORS
// 	r.Use(enableCORS)

// 	// Rota de exemplo
// 	r.HandleFunc("/api/user", func(w http.ResponseWriter, r *http.Request) {
// 		fmt.Fprintln(w, "Hello, User!")
// 	}).Methods(http.MethodGet)

// 	// Adicionando suporte para OPTIONS na rota
// 	r.HandleFunc("/api/user", func(w http.ResponseWriter, r *http.Request) {
// 		w.WriteHeader(http.StatusNoContent) // Retorna 204 No Content
// 	}).Methods(http.MethodOptions)

// 	// Iniciar o servidor
// 	http.ListenAndServe("192.168.56.22:5000", r)
// }
