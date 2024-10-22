package router

import (
	"main/internal/handlers"
	"main/internal/models"
	"net/http"
)

//

var Routes = []models.Routes{

	{
		URI:                   "/api/login",
		Function:              handlers.Login,
		Method:                []string{http.MethodPost},
		RequireAuthentication: false,
	},
	{
		URI:                   "/ws",
		Function:              handlers.WsHandler,
		Method:                []string{http.MethodGet},
		RequireAuthentication: true,
	},

	{
		URI:                   "/api/create-user",
		Function:              handlers.CreateUser,
		Method:                []string{http.MethodPost},
		RequireAuthentication: true,
	},
	// {
	// 	URI:                   "/api/create-panel",
	// 	Function:              handlers.CreatePanel,
	// 	Method:                []string{http.MethodPost},
	// 	RequireAuthentication: true,
	// },
}

/*
r.HandleFunc("/ws", handlers.WsHandler)
r.HandleFunc("/api/sub", handlers.ApiSub).Methods(http.MethodPost, http.MethodOptions)
r.HandleFunc("/api/login", handlers.Login).Methods(http.MethodPost, http.MethodOptions)
r.HandleFunc("/api/create-user", handlers.CreateUser).Methods(http.MethodPost, http.MethodOptions)
r.HandleFunc("/api/create-panel", handlers.CreatePanel).Methods(http.MethodPost, http.MethodOptions)
r.HandleFunc("/api/get-panel-by-name/{name}", handlers.GetPanelByName).Methods(http.MethodPost, http.MethodOptions)
*/
