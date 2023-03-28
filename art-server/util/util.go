package util

import (
	"encoding/json"
	"errors"
	"log"
	"net"
	"net/http"
	"os"
	"strings"
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

// GetIP returns request real ip.
func GetIP(r *http.Request) (string, error) {
	ip := r.Header.Get("X-Real-IP")
	if net.ParseIP(ip) != nil {
		return ip, nil
	}

	ip = r.Header.Get("X-Forward-For")
	for _, i := range strings.Split(ip, ",") {
		if net.ParseIP(i) != nil {
			return i, nil
		}
	}

	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return "", err
	}

	if net.ParseIP(ip) != nil {
		return ip, nil
	}

	return "", errors.New("no valid ip found")
}
