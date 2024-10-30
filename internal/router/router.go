package router

import (
	"fmt"
	"main/internal/middleware"
	"main/internal/models"

	"github.com/gorilla/mux"
)

func GeneratRoutes() *mux.Router {
	fmt.Println("Generating routes...")

	r := mux.NewRouter()

	var newRoutes = []models.Routes{}

	newRoutes = append(newRoutes, Routes...)

	for _, route := range newRoutes {

		if route.RequireAuthentication {
			r.HandleFunc(route.URI, middleware.Logger(middleware.Authentication(route.Function))).Methods(route.Method...)

		} else {
			r.HandleFunc(route.URI, middleware.Logger(route.Function)).Methods(route.Method...)

		}

	}
	return r

}
