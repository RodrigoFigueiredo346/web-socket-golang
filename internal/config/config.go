package config

import (
	"log"

	"github.com/joho/godotenv"
)

type cfg struct {
	SecretKey []byte
	Port      string
	Host      string
	DbPort    string
	User      string
	Password  string
	Dbname    string
}

var appCfg *cfg

func Config() {

	Env, err := godotenv.Read()
	if err != nil {
		log.Fatalf("Error loading .env: %s", err)
	}

	appCfg = &cfg{
		SecretKey: []byte(Env["SECRET_KEY"]),
		Port:      Env["PORT"],
		Host:      Env["DB_HOST"],
		DbPort:    Env["DB_PORT"],
		User:      Env["DB_USER"],
		Password:  Env["DB_PASSWORD"],
		Dbname:    Env["DB_NAME"],
	}

}

func GetConfig() *cfg {
	if appCfg == nil {
		log.Fatalf("Configuration not loaded")
	}
	return appCfg
}
