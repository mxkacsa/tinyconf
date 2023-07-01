# TinyConf

The `tinyconf` package provides a simple way to handle configuration files in Go.

## Features

- Easy to create and manage configuration files.
- Load, save and delete configuration files easily.
- Able to check if a configuration file exists.

## Installation

To install tinyconf, simply run:

```bash
go get github.com/mxkacsa/tinyconf
```

## Examples

### Here is an example on how to use it: 

```go
// any struct can be used as a configuration
type Config struct {
    Secret string `json:"secret"`
    Name   string `json:"name"`
}

func main() {
    // create your configuration variable with default values
    config := Config{
        Secret: "",
        Name:   "Default name",
    }
    
    // if the configuration file exists, load it
    // otherwise create it and save with default values ({secret: "", name: "Default name"})
    err := tinyconf.NewTinyConf().Load(&config)
    if err != nil {
        log.Fatal(err)
    }
    
    log.Println("secret: ", config.Secret, "name: ", config.Name)
}
```

### Example for load or exit

If your service needs a configuration file to run, you can use the `LoadOrExit` method to load the configuration file or exit the program if the configuration file does not exist.

```go
type Config struct {
	Secret string `json:"secret"`
	Name   string `json:"name"`
}

func main() {
    config := Config{
        Secret: "",
        Name:   "Default name",
    }
    
    // If file not exists, it will be created with default values, and program will exit with message.
    // If file exists, it will be loaded into config struct.
    err := tinyconf.NewTinyConf().LoadOrExit(&config)

    // You can set the file name, if you don't set it, the default name will be used.
    // err := tinyconf.NewTinyConf("custom_name.json").LoadOrExit(&config)
	
    if err != nil {
        log.Fatal(err)
    }
    
    log.Println("secret: ", config.Secret, "name: ", config.Name)
}
```

### Example for YAML

```go
type Config struct {
    Secret string `json:"secret"`
    Name   string `json:"name"`
}

func main() {
  config := Config{
    Secret: "",
    Name:   "Default name",
  }

  // You can set the file processor, another processors are: Json, Xml, Toml, MsgPack 
  tc := tinyconf.NewTinyConfWithFp(file_processor.Yaml{})

  // You can set the file name, if you don't set it, the default name will be used
  // tc := tinyconf.NewTinyConfWithFp(file_processor.Yaml{}, "other_file_name.yaml")

  err := tc.Load(&config)
  if err != nil {
    log.Fatal(err)
  }

  config.Secret = "updated yaml secret"

  err = tc.Save(&config)
  if err != nil {
    log.Fatal(err)
  }

  log.Println("secret: ", config.Secret, "name: ", config.Name)
}

```

### Save
You can save the configuration file with the `Save` method
If the configuration file does not exist, it will be created.

```go
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
```

### Delete

You can delete the configuration file with the `Delete` method

```go
// delete the configuration file
err = tc.Delete()
if err != nil {
  log.Fatal(err)
}

log.Println("File deleted")
```

### For more examples, check out the [examples](examples) folder.

# Documentation

### FileProcessor Interface

This interface has methods for reading from and writing to a file, and for returning the default name of a file.

Available implementations:

- JSON (default)
- XML
- YAML - Used package: [yaml.v3](gopkg.in/yaml.v3) 
- TOML - Used package: [go-toml](github.com/pelletier/go-toml)
- MsgPack - Used package: [msgpack](github.com/vmihailenco/msgpack/v5)
### Functions

- `NewTinyConf(fileName ...string) *TinyConf`
    - This function returns a new `TinyConf` object with a JSON `FileProcessor` if no filenames are provided.

- `NewTinyConfWithFp(fp FileProcessor, fileName ...string) *TinyConf`
    - This function returns a new `TinyConf` object with the provided `FileProcessor`.

- `SetExitFunc(ef ExitFunc) *TinyConf`
    - This method sets the function to be used when exiting the program.

- `SetExitCode(code int) *TinyConf`
    - This method sets the exit code.

- `Exists() bool`
    - This method checks if the file at the provided file path exists.

- `LoadOrExit(config interface{}) error`
    - This function tries to load the configuration from a file, if the file doesn't exist, a new one is created and the program exits.

- `Load(config interface{}) error`
    - This function loads the configuration from file to provided config object, if the file doesn't exist, a new one is created with default values.

- `Save(config interface{}) error`
    - This function saves the provided config to the file.

- `Delete() error`
    - This method deletes the file at the file path.
