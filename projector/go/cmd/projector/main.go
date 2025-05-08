package main

import (
	"encoding/json"
	"fmt"
	"log"

	projector "github.com/victorluisca/ts-go-rust/pkg/cli"
)

func main() {
	options, err := projector.GetOptions()
	if err != nil {
		log.Fatalf("unable to get the options %v", err)
	}

	config, err := projector.NewConfig(options)
	if err != nil {
		log.Fatalf("unable to get the config %v", err)
	}

	proj := projector.NewProjector(config)

	if config.Operation == projector.Print {
		if len(config.Arguments) == 0 {
			data := proj.GetValueAll()
			json, err := json.Marshal(data)
			if err != nil {
				log.Fatalf("this line should never be reached %v", err)
			}
			fmt.Printf("%v", string(json))
		} else if value, ok := proj.GetValue(config.Arguments[0]); ok {
			fmt.Printf("%v", value)
		}
	}

	if config.Operation == projector.Add {
		proj.SetValue(config.Arguments[0], config.Arguments[1])
		proj.Save()
	}

	if config.Operation == projector.Delete {
		proj.DeleteValue(config.Arguments[0])
		proj.Save()
	}
}
