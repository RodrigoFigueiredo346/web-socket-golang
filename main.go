package main

import (
	"fmt"
	"main/router"
	"net/http"
	"strconv"
)

const PORT = 8081

func main() {
	mux := router.Routes()

	fmt.Println("Starting server on port", PORT)
	err := http.ListenAndServe(":"+strconv.Itoa(PORT), mux)
	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}

	fmt.Println("Server is running on port", PORT)
}
