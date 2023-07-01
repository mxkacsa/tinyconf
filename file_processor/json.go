package file_processor

import (
	"encoding/json"
	"os"
)

type Json struct {
	Indent       string
	IndentPrefix string
}

func (j Json) DefaultFileName() string {
	return "config.json"
}

func (j Json) WithIndent() Json {
	j.IndentPrefix = ""
	j.Indent = "   "

	return j
}

func (j Json) Write(file *os.File, v interface{}) error {
	encoder := json.NewEncoder(file)
	if j.Indent != "" || j.IndentPrefix != "" {
		encoder.SetIndent(j.IndentPrefix, j.Indent)
	}

	return encoder.Encode(v)
}

func (j Json) Read(file *os.File, config interface{}) error {
	return json.NewDecoder(file).Decode(&config)
}
