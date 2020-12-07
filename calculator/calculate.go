package calculator

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

func Run(csvFile string) {
	fmt.Printf("Running Capital Gains Calculator on %v\n\n", csvFile)

	csvfile, err := os.Open(csvFile)
	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}

	r := csv.NewReader(csvfile)
	i := 0
	columns := []string{}

	state := State{}
	assets := make(map[string]*Asset)
	state.Assets = assets
	for {
		row, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		if i == 0 {
			columns = row
		} else {
			parsedTaxEvent := recordRowToTaxEvent(row, columns)
			// parsedTaxEvent.PrintToLine()
			state.applyNewTaxEvent(parsedTaxEvent)
		}
		i++
	}
	// Print state
	state.PrintGainsAndLosses()
}
