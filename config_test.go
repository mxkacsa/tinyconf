package tinyconf

import (
	"github.com/mxkacsa/tinyconf/file_processor"
	"log"
	"reflect"
	"testing"
)

type Config struct {
	Secret string `json:"secret"`
	Name   string `json:"name"`
}

func TestLoadOrExit(t *testing.T) {
	conf := Config{
		Secret: "def",
		Name:   "def name",
	}

	tc := NewTinyConf("test.json")

	tc.SetExitFunc(func(code int) {
		if code != tc.exitCode {
			t.Fatal("exit code should be ", tc.exitCode)
		}
	})

	defer deleteFile(t, tc)
	err := tc.LoadOrExit(&conf)
	if err != nil {
		t.Fatal(err)
	}
}

func TestLoadOrExitInvalidCode(t *testing.T) {
	conf := Config{
		Secret: "def",
		Name:   "def name",
	}

	tc := NewTinyConf("test.json")

	tc.SetExitCode(1)
	tc.SetExitFunc(func(code int) {
		if code != tc.exitCode {
			t.Fatal("exit code should be ", tc.exitCode)
		}
	})

	defer deleteFile(t, tc)
	err := tc.LoadOrExit(&conf)
	if err != nil {
		t.Fatal(err)
	}
}

func TestCreateFileWithSaveMethod(t *testing.T) {
	conf := Config{
		Secret: "def",
		Name:   "def name",
	}

	tc := NewTinyConf("test.json")
	defer deleteFile(t, tc)

	err := tc.Save(&conf)
	if err != nil {
		t.Fatal(err)
	}

	if !tc.Exists() {
		t.Fatal("file should exist")
	}
}

func TestTinyConfExistsAndLoad(t *testing.T) {
	processors := []FileProcessor{
		file_processor.Json{}.WithIndent(),
		file_processor.Toml{},
		file_processor.Yaml{},
		file_processor.Xml{}.WithIndent(),
		file_processor.MsgPack{},
	}

	for _, processor := range processors {
		testProcessor(t, processor)
	}

	log.Println("TestTinyConfExistsAndLoad passed")
}

func testProcessor(t *testing.T, fp FileProcessor) {
	processorName := reflect.TypeOf(fp).Name()

	log.Println("Testing ", processorName, " processor")
	tc := NewTinyConfWithFp(fp)

	if tc.Exists() {
		t.Fatal(processorName + " TinyConf.Exists() should return false")
	}

	config := Config{
		Secret: "",
		Name:   "Default name",
	}

	err := tc.Load(&config)
	defer deleteFile(t, tc)

	if err != nil {
		t.Fatal(err)
	}

	if !tc.Exists() {
		t.Fatal(processorName + " TinyConf.Exists() should return true")
	}

	if config.Secret != "" {
		t.Fatal(processorName + " config.Secret should be empty")
	}

	if config.Name != "Default name" {
		t.Fatal(processorName + " config.Name should be 'Default name'")
	}

	config.Secret = "updated_secret"
	config.Name = "updated_name"
	err = tc.Save(&config)
	if err != nil {
		t.Fatal(err)
	}

	config.Secret = ""
	config.Name = ""

	err = tc.Load(&config)
	if err != nil {
		t.Fatal(err)
	}

	if config.Secret != "updated_secret" {
		t.Fatal(processorName+" config.Secret should be 'updated_secret', but it is ", config.Secret)
	}

	if config.Name != "updated_name" {
		t.Fatal(processorName+" config.Name should be 'updated_name', but it is ", config.Name)
	}
}

func deleteFile(t *testing.T, tc *TinyConf) {
	log.Println("Deleting ", tc.filePath)
	err := tc.Delete()
	if err != nil {
		t.Fatal(err)
	}
}
