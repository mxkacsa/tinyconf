package file_processor

import (
	"encoding/xml"
	"os"
)

type Xml struct {
	Indent       string
	IndentPrefix string
}

func (x Xml) DefaultFileName() string {
	return "config.xml"
}

func (x Xml) WithIndent() Xml {
	x.IndentPrefix = "  "
	x.Indent = "    "

	return x
}

func (x Xml) Write(file *os.File, v interface{}) error {
	encoder := xml.NewEncoder(file)
	if x.Indent != "" || x.IndentPrefix != "" {
		encoder.Indent(x.IndentPrefix, x.Indent)
	}
	return encoder.Encode(v)
}

func (x Xml) Read(file *os.File, v interface{}) error {
	return xml.NewDecoder(file).Decode(v)
}
