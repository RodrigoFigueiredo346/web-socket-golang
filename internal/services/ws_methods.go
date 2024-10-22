package services

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"main/internal/db"
	"main/internal/mqtt"
	"main/sqlc"
	"strconv"
	"time"

	"golang.org/x/exp/rand"
)

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
	_ = strconv.Itoa(int(idpanel))

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
	// Verifica se params é um slice de interface
	paramSlice, ok := params.([]interface{})
	if !ok || len(paramSlice) == 0 {
		return 0, errors.New("invalid parameters")
	}

	// Verifica se o primeiro item do slice é um mapa
	paramMap, ok := paramSlice[0].(map[string]interface{})
	if !ok {
		return 0, errors.New("first parameter is not a map")
	}

	paramMap["id"] = rand.Intn(89999999) + 10000000

	identifier, ok := paramMap["identifier"].(string)
	if !ok {
		return 0, errors.New("identifier is not a string or missing")
	}

	cache.Set(identifier, paramMap)

	panelCreated, _ := cache.Get(string(identifier))
	fmt.Println("Panel created:", panelCreated.(map[string]interface{})["identifier"])

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

func DBString(s string) sql.NullString {
	if s == "" {
		return sql.NullString{}
	}
	return sql.NullString{String: s, Valid: true}
}

func DBInt32(i int32) sql.NullInt32 {
	if i == 0 {
		return sql.NullInt32{}
	}
	return sql.NullInt32{Int32: i, Valid: true}
}

func DBTime(t time.Time) sql.NullTime {
	if t.IsZero() {
		return sql.NullTime{}
	}
	return sql.NullTime{Time: t, Valid: true}
}
