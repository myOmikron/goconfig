package staticconfig

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
)

type Conf struct {
	config interface{}
}

func NewConfiguration(config interface{}) (c Conf) {
	c.config = config
	return
}

func (c Conf) ParseConfig(filename string) interface{} {
	b, err := ioutil.ReadFile(filename)
	if errors.Is(err, os.ErrNotExist) {
		fmt.Printf("%s does not exist, creating config\n", filename)
		if encoded, err := json.MarshalIndent(&c.config, "", "  "); err != nil {
			panic(err.Error())
		} else {
			if err := ioutil.WriteFile(filename, encoded, fs.FileMode(0700)); err != nil {
				panic(fmt.Sprintf("Could not write to %s\n", filename))
			}
			fmt.Println("Done, quitting ..")
			os.Exit(0)
		}
	}

	if err := json.Unmarshal(b, &c.config); err != nil {
		panic(err.Error())
	}
	return &c.config
}
