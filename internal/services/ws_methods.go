package services

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"main/internal/db"
	"main/internal/models"
	"main/internal/mqtt"
	"main/sqlc"
	"time"
)

func LoadPanelsInMemo() {
	start := time.Now()
	ctx := context.Background()
	dt := sqlc.New(db.DB)
	allPanels, err := dt.GetAllPanels(ctx)
	if err != nil {
		log.Println(err)
	}

	//var panelModels []models.PanelModel
	for _, panel := range allPanels {
		panelModel := models.PanelModel{
			IDPanel:    panel.Idpanel,
			Identifier: panel.Identifier,
			DscPanel:   panel.DscPanel,
			NumSerie:   panel.NumSerie,
			Active:     int(panel.Active.Int32),
			CtrlBright: int(panel.CtrlBright.Int32),
		}
		//panelModels = append(panelModels, panelModel)

		cache.Set(panel.Idpanel, panelModel)

	}
	fmt.Println("paineis carregados do banco em: ", time.Since(start))

}

func (ws *WsService) readPanels(params interface{}, id int) (interface{}, error) {

	panels := cache.GetAll()

	var panelModels []models.PanelModel
	for _, panel := range panels {
		fmt.Println(panel.(models.PanelModel))
		panelMap, ok := panel.(models.PanelModel)
		if !ok {
			fmt.Println("Formato de painel inválido")
			continue
		}

		panelModel := models.PanelModel{
			IDPanel:    panelMap.IDPanel,
			Identifier: panelMap.Identifier,
			DscPanel:   panelMap.DscPanel,
			NumSerie:   panelMap.NumSerie,
			Active:     panelMap.Active,
			CtrlBright: panelMap.CtrlBright,
		}

		panelModels = append(panelModels, panelModel)
	}

	return panelModels, nil
}

func (ws *WsService) editPanel(params interface{}, id int) (interface{}, error) {

	/*{
	"method":"editPanel",
	"params":[{
		"idpanel": "122345678",
		"identifier": "4654654",
		"dscpanel":"PMV 01",
		"num_serie":"123456",
		"active":"1",
		"ctrl_bright":"2"
	}],
	"id":12345
	}*/

	// type EditPanelParamsModel struct {
	// 	IDPanel    string `json:"idpanel"` // Único campo obrigatório
	// 	Identifier string `json:"identifier,omitempty"`
	// 	DscPanel   string `json:"dsc_panel,omitempty"`
	// 	NumSerie   string `json:"num_serie,omitempty"`
	// 	Active     int    `json:"active,omitempty"`
	// 	CtrlBright int    `json:"ctrl_bright,omitempty"`
	// }
	var panelParams []models.PanelModel

	jsonData, err := json.Marshal(params)
	if err != nil {
		return nil, errors.New("failed to marshal params to JSON")
	}

	err = json.Unmarshal(jsonData, &panelParams)
	if err != nil {
		return nil, errors.New("invalid parameters: could not decode into models.EditPanelParamsModel")
	}

	panel := panelParams[0]
	if panel.IDPanel == "" {
		return nil, errors.New("idpanel is required for editing")
	}

	//TODO buscar os campos que não vieram na requisição na memória ou banco ou usar um db.Exec puro em UPDATES

	dt := sqlc.New(db.DB)
	err = dt.UpdatePanel(context.Background(), sqlc.UpdatePanelParams{
		Idpanel:    panel.IDPanel,
		Identifier: panel.Identifier,
		DscPanel:   panel.DscPanel,
		NumSerie:   panel.NumSerie,
		Active:     sql.NullInt32{Int32: int32(panel.Active), Valid: true},
		CtrlBright: DBInt32(int32(panel.CtrlBright)),
		DthrAlt:    DBTime(time.Now()),
	})
	if err != nil {
		return nil, errors.New("failed to update panel")
	}

	cache.Set(panel.IDPanel, panel)
	panelUpdated, _ := cache.Get(string(panel.IDPanel))
	if panelUpdated == nil {
		return nil, errors.New("failed to update panel in memory")
	}

	return 1, nil
}

func (ws *WsService) readPanelStatus(params interface{}, id int) (interface{}, error) {

	ctx := context.Background()

	paramSlice, ok := params.([]interface{})
	if !ok {
		return nil, errors.New("invalid parameters")
	}

	if len(paramSlice) == 0 {
		panels := cache.GetAll()
		return panels, nil

	}

	identifier := paramSlice[0].(string)
	fmt.Println("identifier", identifier)

	mqtt.Publish(identifier, `{
				"method": "getStatus",
				"params": {
					"token": "XXYYZZ"
				},
				"id": 12563
			}`)

	// fmt.Println("Params: ", params)
	// fmt.Println("paramSlice: ", paramSlice)

	// var identifierParam sql.NullString
	// if identifier == "" {
	// 	// Se a string estiver vazia, o valor será tratado como nulo
	// 	identifierParam = sql.NullString{
	// 		String: "",
	// 		Valid:  false,
	// 	}
	// } else {
	// 	// Se houver uma string, definimos como válida
	// 	identifierParam = sql.NullString{
	// 		String: identifier,
	// 		Valid:  true,
	// 	}
	// }

	dt := sqlc.New(db.DB)
	idpanel, err := dt.GetPanelByIdentifier(ctx, identifier)
	if err != nil {
		return nil, err
	}

	fmt.Println("idpanel: ", idpanel)

	/*
		salvo em sinc3
		- idsinc            1
		- idpanel           2
		- tag               readPanelStatus
		- data
		- dthr_ins          2024-10-17 12:30:00.000
		- sinc              0
		- dthr_sinc
	*/

	idsinc, err := dt.CreateSinc(ctx, sqlc.CreateSincParams{
		Idpanel: identifier,
		Tag:     "readPanelStatus",
		Data:    DBString(""),
		Sinc:    0})

	if err != nil {
		return nil, err
	}

	fmt.Println("idsinc: ", idsinc)

	/*
		envio para mqtt no topico <identifier>
		***com ID de idsinc = 1***
		- {"method": "getStatus", "params": {"token": "XXYYZZ"}, "ID": <idsinc = 1>}
	*/

	/*
		quando chegar pelo mqtt a resposta com esse ID
	*/

	/*
		update em sinc3 com ID = 1
		- sinc              1
		- dthr_sinc         2024-10-17 12:30:00.000
		- data              {"result": {"pv":[1263.0,0.0,3600.0],"load"... }, "error": null, "id": 1}}
	*/

	/*
		envio para frontend
	*/

	// _ = dt.CreateSinc(ctx, sqlc.CreateSincParams{Idpanel: identifierParam})

	return map[string]string{"status": "ok"}, nil
}

func (ws *WsService) createPanel(params interface{}, id int) (interface{}, error) {
	start := time.Now()
	var panelParams []models.PanelModel

	jsonData, err := json.Marshal(params)
	if err != nil {
		return nil, errors.New("failed to marshal params to JSON")
	}

	err = json.Unmarshal(jsonData, &panelParams)
	if err != nil {
		return nil, errors.New("invalid parameters: could not decode into models.CreatePanelParamsModel")
	}

	panel := panelParams[0]

	//TODO validação dos campos - fazer um método pra isso??
	if panel.Identifier == "" {
		return nil, errors.New("identifier is required")
	}
	if panel.DscPanel == "" {
		return nil, errors.New("dsc_panel is required")
	}
	if panel.Active < 0 || panel.Active > 1 {
		return nil, errors.New("active must be 0 or 1")
	}
	if panel.CtrlBright < 1 || panel.CtrlBright > 2 {
		return nil, errors.New("ctrl_bright must be 1 or 2")
	}

	dt := sqlc.New(db.DB)
	idpanelCreated, err := dt.CreatePanel(context.Background(), sqlc.CreatePanelParams{
		Identifier: panel.Identifier,
		DscPanel:   panel.DscPanel,
		NumSerie:   panel.NumSerie,
		Active:     sql.NullInt32{Int32: int32(panel.Active), Valid: true},
		CtrlBright: DBInt32(int32(panel.CtrlBright)),
	})
	if err != nil {
		return nil, err
	}

	panel.IDPanel = idpanelCreated

	//TODO se falhar no cache mas já ter sido salvo no banco???
	cache.Set(panel.IDPanel, panel)
	fmt.Println("Created panel:", panel)
	panelCreated, _ := cache.Get(string(panel.IDPanel))
	if panelCreated == nil {
		return nil, errors.New("failed to create panel in memory")
	}

	fmt.Println("Created panel", idpanelCreated)
	fmt.Println("Created panel in time:", time.Since(start))

	return 1, nil
}

func (ws *WsService) cfgDateTime(params interface{}, id int) (interface{}, error) {

	// {
	// 	"method": "cfgDateTime",
	// 	"params": {
	// 		"token": "XXYYZZ",
	// 		"dateTime": "2024-04-26 10:56:54"
	// 	},
	// 	"id": 12563
	// }
	return 1, nil
}

func (ws *WsService) activateMsg(params interface{}, id int) (interface{}, error) {

	paramsMap, ok := params.(map[string]interface{})
	if !ok {
		return nil, errors.New("invalid parameters")
	}

	identifier, ok := paramsMap["identifier"].(string)
	if !ok {
		return nil, errors.New("identifier is not a string or missing")
	}
	msg, ok := paramsMap["msg"].(string)
	if !ok {
		return nil, errors.New("msg is not a string or missing")
	}

	mqtt.Publish(identifier, `{
		"method": "activateMsg",
		"params": {
			"token": "XXYYZZ",
			"msg": "`+msg+`"
		},
		"id": 12563
	}`)

	return 1, nil
}
