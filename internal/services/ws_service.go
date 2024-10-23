package services

import (
	"errors"
)

type WsService struct {
	methods map[string]func(params interface{}, id int) (interface{}, error)
}

// NewWsService cria uma nova instância de WsService
func NewWsService() *WsService {
	ws := &WsService{
		methods: make(map[string]func(params interface{}, id int) (interface{}, error)),
	}
	ws.registerMethods()
	return ws
}

func (ws *WsService) registerMethods() {
	ws.methods["readPanelStatus"] = ws.readPanelStatus
	ws.methods["createPanel"] = ws.createPanel
	ws.methods["editPanel"] = ws.editPanel
	// Adicionar outros métodos aqui...
}

func (ws *WsService) HandleRequest(method string, params interface{}, id int) (interface{}, error) {

	if handler, ok := ws.methods[method]; ok {
		return handler(params, id)
	}
	return nil, errors.New("method not found")
}
