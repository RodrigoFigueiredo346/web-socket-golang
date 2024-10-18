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
