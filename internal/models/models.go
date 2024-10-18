package models

import "net/http"

type JsonRpcRequest struct {
	Method string      `json:"method"`
	Params interface{} `json:"params"`
	ID     int         `json:"id"`
}

type JsonRpcResponse struct {
	Result interface{} `json:"result,omitempty"`
	Error  interface{} `json:"error,omitempty"`
	ID     int         `json:"id"`
}

type PanelModel struct {
	ID          int      `json:"id,omitempty"`
	Identifier  string   `json:"identifier,omitempty"`
	Dscpanel    string   `json:"dscpanel"`
	Num_serie   string   `json:"num_serie,omitempty"`
	Active      int      `json:"active,omitempty"`
	Ctrl_bright int      `json:"ctrl_bright,omitempty"`
	Messages    []string `json:"messages,omitempty"`
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
