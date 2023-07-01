package file_processor

import (
	"gopkg.in/yaml.v3"
	"os"
)

type Yaml struct{}

func (y Yaml) DefaultFileName() string {
	return "config.yaml"
}

func (y Yaml) Write(file *os.File, v interface{}) error {
	return yaml.NewEncoder(file).Encode(v)
}

func (y Yaml) Read(file *os.File, v interface{}) error {
	return yaml.NewDecoder(file).Decode(v)
}
