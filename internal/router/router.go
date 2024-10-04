package router

import (
	"main/internal/middleware"
	"main/internal/models"
	"net/http"

	"github.com/gorilla/mux"
)

func GeneratRoutes() *mux.Router {

	r := mux.NewRouter()
	r.Use(enableCORS)

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
