package main

import (
	"log"
	"os"

	"github.com/ksahli/lyagushka/cmd/generate"
)

func main() {
	log.Println("lyagushka - static site generator")
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	if err := generate.Run(wd); err != nil {
		log.Fatal(err)
	}
}
