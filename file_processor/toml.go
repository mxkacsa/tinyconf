package file_processor

import (
	"github.com/pelletier/go-toml"
	"os"
)

type Toml struct{}

func (t Toml) DefaultFileName() string {
	return "config.toml"
}

func (t Toml) Write(file *os.File, v interface{}) error {
	return toml.NewEncoder(file).Encode(v)
}

func (t Toml) Read(file *os.File, v interface{}) error {
	return toml.NewDecoder(file).Decode(v)
}
