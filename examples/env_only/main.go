package main

import (
	"github.com/mxkacsa/tinyconf"
	"log"
)

type Config struct {
	Secret   string `env:"APP_SECRET"`
	Name     string `env:"APP_NAME"`
	Port     int    `env:"APP_PORT"`
	Debug    bool   `env:"APP_DEBUG"`
	Database struct {
		Host string `env:"DB_HOST"`
		Port int    `env:"DB_PORT"`
	}
}

// This example demonstrates how to load configuration only from environment variables,
// without using any config file.
//
// Usage:
//
//	APP_SECRET=mysecret APP_NAME=myapp APP_PORT=8080 APP_DEBUG=true DB_HOST=localhost DB_PORT=5432 go run main.go
func main() {
	config := Config{
		Secret: "",
		Name:   "Default name",
		Port:   3000,
		Debug:  false,
	}
	config.Database.Host = "localhost"
	config.Database.Port = 5432

	// LoadFromEnv loads only from environment variables, no config file needed
	err := tinyconf.LoadFromEnv(&config)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("secret:", config.Secret, "name:", config.Name, "port:", config.Port, "debug:", config.Debug)
	log.Println("db host:", config.Database.Host, "db port:", config.Database.Port)
}
