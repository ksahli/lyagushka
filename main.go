package main

import (
	"log"
	"os"

	"github.com/ksahli/lyagushka/cmd/clean"
	"github.com/ksahli/lyagushka/cmd/generate"
	"github.com/ksahli/lyagushka/cmd/initialize"
)

func main() {
	log.Println("lyagushka - static site generator")

	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	switch os.Args[1] {
	case "initialize":
		if err := initialize.Run(wd); err != nil {
			log.Fatal(err)
		}
	case "generate":
		if err := generate.Run(wd); err != nil {
			log.Fatal(err)
		}
	case "clean":
		if err := clean.Run(wd); err != nil {
			log.Fatal(err)
		}
	}
}
