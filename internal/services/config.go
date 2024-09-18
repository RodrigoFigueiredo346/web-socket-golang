package services

import (
	"log"

	"github.com/joho/godotenv"
)

type cfg struct {
	SecretKey string
	Port      string
}

var appCfg *cfg

func Config() {

	Env, err := godotenv.Read()
	if err != nil {
		log.Fatalf("Erro ao carregar arquivo .env: %s", err)
	}

	appCfg = &cfg{
		SecretKey: Env["SECRET_KEY"],
		Port:      Env["PORT"],
	}

}

func GetConfig() *cfg {
	if appCfg == nil {
		log.Fatalf("Configuração não inicializada")
	}
	return appCfg
}
