package main

import (
	"github.com/Brialius/servicesceleton/cmd"
	"log"
)

func main() {
	if err := cmd.RootCommand().Execute(); err != nil {
		log.Fatal(err)
	}
}
