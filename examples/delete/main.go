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

	err = tc.Delete()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("File deleted")
}
