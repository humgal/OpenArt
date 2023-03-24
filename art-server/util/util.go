package util

import (
	"encoding/json"
	"log"
	"os"
)

type MyConfig struct {
	Redishost   string
	Redispasswd string
	Redisdb     int
	Mysqlhost   string
}

var Config MyConfig
var Logger *log.Logger

func init() {
	file, err := os.OpenFile("openart.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("Failed to open log file", err)
	}
	Logger = log.New(file, "INFO", log.Ldate|log.Ltime|log.Lshortfile)

	configfile, err := os.Open("config.json")
	if err != nil {
		log.Fatal("Failed to open config file", err)
	}

	decoder := json.NewDecoder(configfile)
	Config = MyConfig{}
	err = decoder.Decode(&Config)
	if err != nil {
		log.Fatal("Failed to open config file", err)
	}

}
