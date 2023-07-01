package main

import (
	"github.com/mxkacsa/tinyconf"
	"log"
)

type Config struct {
	Secret string `json:"secret"`
	Name   string `json:"name"`
}

func main() {
	config := Config{
		Secret: "",
		Name:   "Default name",
	}

	tc := tinyconf.NewTinyConf()

	err := tc.Load(&config)
	if err != nil {
		log.Fatal(err)
	}

	config.Name = "Updated name"

	err = tc.Save(&config)
	if err != nil {
		log.Fatal(err)
	}
}
