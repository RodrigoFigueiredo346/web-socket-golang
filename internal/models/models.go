package models

import "net/http"

type JsonRpcRequest struct {
	Method string      `json:"method"`
	Params interface{} `json:"params"`
	ID     int         `json:"id"`
}

type Error struct {
	Code int `json:"code"`
}

type JsonRpcResponse struct {
	Result interface{} `json:"result"`
	Error  *Error      `json:"error"`
	ID     int         `json:"id"`
}

type PanelModel struct {
	IDPanel    string `json:"idpanel,omitempty"`
	Identifier string `json:"identifier,omitempty"`
	DscPanel   string `json:"dsc_panel"`
	NumSerie   string `json:"num_serie,omitempty"`
	Active     int    `json:"active"`
	CtrlBright int    `json:"ctrl_bright,omitempty"`
	//Messages    []string `json:"messages,omitempty"`
}

type User struct {
	Username string
	Password string
	ID       int
	Token    string
}

type Routes struct {
	URI                   string
	Method                []string
	Function              func(http.ResponseWriter, *http.Request)
	RequireAuthentication bool
}

// type CreatePanelParamsModel struct {
// 	Idpanel    string `json:"idpanel,omitempty"`
// 	Identifier string `json:"identifier,omitempty"`
// 	DscPanel   string `json:"dsc_panel,omitempty"`
// 	NumSerie   string `json:"num_serie,omitempty"`
// 	Active     int    `json:"active,omitempty"`
// 	CtrlBright int    `json:"ctrl_bright,omitempty"`
// }
