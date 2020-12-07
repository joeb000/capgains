package main

import (
	"flag"
	"log"

	"github.com/joeb000/capgains/calculator"
)

var fileFlag = flag.String("file", "", "path to a csv file")

func main() {
	flag.Parse()

	if *fileFlag == "" {
		log.Fatal("no file provided")
	}

	calculator.Run(*fileFlag)
}
