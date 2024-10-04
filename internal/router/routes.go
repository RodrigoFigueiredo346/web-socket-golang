package router

import (
	"main/internal/handler"
	"main/internal/models"
	"net/http"
)

//

var Routes = []models.Routes{
	{
		URI:                   "/api/sub",
		Function:              handler.ApiSub,
		Method:                []string{http.MethodPost},
		RequireAuthentication: false,
	},
	{
		URI:                   "/api/login",
		Function:              handler.Login,
		Method:                []string{http.MethodPost},
		RequireAuthentication: false,
	},
	{
		URI:                   "/ws",
		Function:              handler.WsHandler,
		Method:                []string{""},
		RequireAuthentication: true,
	},

	{
		URI:                   "/api/create-user",
		Function:              handler.CreateUser,
		Method:                []string{http.MethodPost},
		RequireAuthentication: true,
	},
	{
		URI:                   "/api/create-panel",
		Function:              handler.CreatePanel,
		Method:                []string{http.MethodPost},
		RequireAuthentication: true,
	},
}

/*
r.HandleFunc("/ws", handler.WsHandler)
r.HandleFunc("/api/sub", handler.ApiSub).Methods(http.MethodPost, http.MethodOptions)
r.HandleFunc("/api/login", handler.Login).Methods(http.MethodPost, http.MethodOptions)
r.HandleFunc("/api/create-user", handler.CreateUser).Methods(http.MethodPost, http.MethodOptions)
r.HandleFunc("/api/create-panel", handler.CreatePanel).Methods(http.MethodPost, http.MethodOptions)
r.HandleFunc("/api/get-panel-by-name/{name}", handler.GetPanelByName).Methods(http.MethodPost, http.MethodOptions)
*/
