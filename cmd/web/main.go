package main

import (
	"log"
	"main/router"
	"net/http"
	"strconv"
)

const PORT = 8080

func main() {
	mux := router.Routes()

	log.Println(PORT)

	_ = http.ListenAndServe(":"+strconv.Itoa(PORT), mux)
}
