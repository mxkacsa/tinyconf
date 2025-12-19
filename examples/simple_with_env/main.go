package main

import (
	"github.com/mxkacsa/tinyconf"
	"log"
)

type Config struct {
	Secret string `json:"secret" env:"APP_SECRET"`
	Name   string `json:"name" env:"APP_NAME"`
	Port   int    `json:"port" env:"APP_PORT"`
	Debug  bool   `json:"debug" env:"APP_DEBUG"`
}

// This example demonstrates how to use environment variables to override config file values.
// The config file is loaded first, then environment variables override any matching fields.
//
// Usage:
//
//	APP_SECRET=mysecret APP_NAME=myapp APP_PORT=8080 APP_DEBUG=true go run main.go
func main() {
	config := Config{
		Secret: "",
		Name:   "Default name",
		Port:   3000,
		Debug:  false,
	}

	// LoadWithEnv loads the config file first, then overrides with env variables
	err := tinyconf.NewTinyConf().LoadWithEnv(&config)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("secret:", config.Secret, "name:", config.Name, "port:", config.Port, "debug:", config.Debug)
}
