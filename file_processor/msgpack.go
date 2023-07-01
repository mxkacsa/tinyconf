package file_processor

import (
	"github.com/vmihailenco/msgpack/v5"
	"os"
)

type MsgPack struct{}

func (s MsgPack) DefaultFileName() string {
	return "config.msgpack"
}

func (s MsgPack) Write(file *os.File, v interface{}) error {
	return msgpack.NewEncoder(file).Encode(v)
}

func (s MsgPack) Read(file *os.File, v interface{}) error {
	return msgpack.NewDecoder(file).Decode(v)
}
