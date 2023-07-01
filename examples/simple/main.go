package main

import (
	"github.com/mxkacsa/tinyconf"
	"log"
)

type Config struct {
	Secret string `json:"secret"`
	Name   string `json:"name"`
}

// if file not exists, it will be created with default values
// if file exists, it will be loaded into config struct
func main() {
	config := Config{
		Secret: "",
		Name:   "Default name",
	}

	err := tinyconf.NewTinyConf().Load(&config)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("secret: ", config.Secret, "name: ", config.Name)
}
