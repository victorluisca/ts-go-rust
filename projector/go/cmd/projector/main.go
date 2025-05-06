package main

import (
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

	fmt.Printf("%+v", config)
}
