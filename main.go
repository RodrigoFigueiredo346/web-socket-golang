package main

import (
	"fmt"
	"log"
	"main/internal/mqtt"
	"main/internal/router"
	"main/internal/services"

	"net/http"
)

func main() {

	go mqtt.InitMqtt()
	services.Config()

	cfg := services.GetConfig()

	r := router.GeneratRoutes()

	fmt.Printf("WebSocket running in ws://0.0.0.0:%s/ws\n", cfg.Port)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", cfg.Port), r))
}
