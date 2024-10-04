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
	ID         int      `json:"id,omitempty"`
	Name       string   `json:"name"`
	Status     string   `json:"status,omitempty"`
	Messages   []string `json:"messages,omitempty"`
	BrightMode int      `json:"bright_mode,omitempty"`
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
