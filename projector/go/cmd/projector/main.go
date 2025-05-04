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

	fmt.Printf("%+v", options)
}
