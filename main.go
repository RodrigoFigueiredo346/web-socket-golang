package main

import (
	"fmt"
	"log"
	"main/internal/config"
	"main/internal/db"
	"main/internal/mqtt"
	"main/internal/router"
	"main/internal/services"
	"net/http"

	"github.com/gorilla/handlers"
)

func main() {

	// Goro()

	config.Config()

	go mqtt.InitMqtt()

	db.InitDB()

	cfg := config.GetConfig()

	r := router.GeneratRoutes()

	//aqui podemos subir as informações pro cache ??? quantas serão???
	services.LoadPanelsInMemo()

	fmt.Printf("WebSocket running in ws://0.0.0.0:%s/ws\n", cfg.Port)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", cfg.Port), handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
	)(r)))

}
