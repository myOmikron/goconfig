# goconfig


## Usage

```go
package main

import (
	"fmt"
	"github.com/myOmikron/goconfig"
)

type Database struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
}

type MyConfig struct {
	Uri     string        `json:"uri"`     // String
	Count   int           `json:"count"`   // Int
	Servers []interface{} `json:"servers"` // Generic list
	DB      Database      `json:"db"`      // Nested Object
}

func main() {
	// Use empty config without defaults
	conf := MyConfig{}

	// or set some defaults
	conf = MyConfig{
		Servers: []interface{}{
			"127.0.0.1",
			"localhost",
		},
		DB: Database{Port: 3306},
	}

	// Create / Parse config file
	goconfig.ParseConfig("config.json", &conf)

	fmt.Printf("Current configuration: %#v\n", conf)
}

```

