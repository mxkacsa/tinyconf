package tinyconf

import (
	"fmt"
	"github.com/mxkacsa/tinyconf/file_processor"
	"log"
	"os"
	"path/filepath"
)

// FileProcessor is an interface which has methods for reading and writing a file,
// and for returning the default name of a file.
type FileProcessor interface {
	DefaultFileName() string
	Write(file *os.File, v interface{}) error
	Read(file *os.File, v interface{}) error
}

// ExitFunc is a function type that accepts an integer parameter representing the exit code.
type ExitFunc func(code int)

// TinyConf is a structure which contains the configurations for handling file operations.
type TinyConf struct {
	filePath      string
	fileProcessor FileProcessor
	exitFunc      ExitFunc
	exitCode      int
}

// NewTinyConf returns a new TinyConf object with a default JSON file processor if no filenames are provided.
func NewTinyConf(fileName ...string) *TinyConf {
	return NewTinyConfWithFp(file_processor.Json{}, fileName...)
}

// NewTinyConfWithFp returns a new TinyConf object with the provided FileProcessor and
// filename(s) if provided, otherwise it takes the default file name from FileProcessor.
func NewTinyConfWithFp(fp FileProcessor, fileName ...string) *TinyConf {
	name := fp.DefaultFileName()
	if len(fileName) > 0 {
		name = fileName[0]
	}

	return &TinyConf{
		filePath:      name,
		fileProcessor: fp,
		exitFunc:      os.Exit,
	}
}

// SetExitFunc sets the function to be used when exiting program.
func (c *TinyConf) SetExitFunc(ef ExitFunc) *TinyConf {
	c.exitFunc = ef

	return c
}

// SetExitCode sets the exit code.
func (c *TinyConf) SetExitCode(code int) *TinyConf {
	c.exitCode = code

	return c
}

// Exists checks whether the file in the filePath exists or not.
func (c *TinyConf) Exists() bool {
	_, err := os.Stat(c.filePath)
	return err == nil || !os.IsNotExist(err)
}

// LoadOrExit tries to load the configuration from a file, if the file doesn't exist, a new one is created and the program exits.
func (c *TinyConf) LoadOrExit(config interface{}) error {
	exists := c.Exists()

	if err := c.Load(config); err != nil {
		return err
	}

	if !exists {
		log.Println(fmt.Sprintf("A new file has been created: %s. Please edit it and restart the program.", c.filePath))
		c.exitFunc(c.exitCode)
	}

	return nil
}

// Load loads the configuration from file to provided config object, if the file doesn't exist, a new one is created with default values.
func (c *TinyConf) Load(config interface{}) error {
	file, err := os.Open(c.filePath)
	defer file.Close()
	if err != nil {
		if os.IsNotExist(err) {
			return c.createFile(config)
		}
		return err
	}

	return c.fileProcessor.Read(file, config)
}

// Save writes the given configuration to the file.
func (c *TinyConf) Save(config interface{}) error {
	if !c.Exists() {
		return c.createFile(config)
	}

	file, err := os.OpenFile(c.filePath, os.O_RDWR, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	return c.fileProcessor.Write(file, config)
}

// Delete deletes the file from the filePath.
func (c *TinyConf) Delete() error {
	return os.Remove(c.filePath)
}

// createFile creates a new configuration file.
func (c *TinyConf) createFile(config interface{}) error {
	dir := filepath.Dir(c.filePath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}

	file, err := os.Create(c.filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	return c.fileProcessor.Write(file, config)
}

// LoadWithEnv loads configuration from file and then overrides values with environment variables.
// Use the `env` struct tag to specify the environment variable name.
// Example: `env:"APP_SECRET"` will load the value from the APP_SECRET environment variable.
func (c *TinyConf) LoadWithEnv(config interface{}) error {
	if err := c.Load(config); err != nil {
		return err
	}

	return applyEnvOverrides(config)
}

// LoadOrExitWithEnv tries to load configuration from file and overrides with env variables.
// If the file doesn't exist, a new one is created and the program exits.
func (c *TinyConf) LoadOrExitWithEnv(config interface{}) error {
	exists := c.Exists()

	if err := c.LoadWithEnv(config); err != nil {
		return err
	}

	if !exists {
		log.Println(fmt.Sprintf("A new file has been created: %s. Please edit it and restart the program.", c.filePath))
		c.exitFunc(c.exitCode)
	}

	return nil
}
