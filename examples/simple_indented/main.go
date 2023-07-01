package main

import (
	"github.com/mxkacsa/tinyconf"
	"github.com/mxkacsa/tinyconf/file_processor"
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

	err := tinyconf.NewTinyConfWithFp(file_processor.Json{}.WithIndent()).Load(&config)
	if err != nil {
		log.Fatal(err)
	}

	// or if you want to customize indent:

	//err = tinyconf.NewTinyConfWithFp(file_processor.Json{
	//	Indent:       "   ",
	//	IndentPrefix: " ",
	//}).Load(&config)
	//if err != nil {
	//	log.Fatal(err)
	//}

	log.Println("secret: ", config.Secret, "name: ", config.Name)
}
