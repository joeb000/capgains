package calculator

import (
	"log"
	"strconv"
)

func recordRowToTaxEvent(records []string, columns []string) TaxEvent {
	taxEvent := TaxEvent{}
	for i := 0; i < 7; i++ {
		switch columns[i] {
		case "BoughtAsset":
			taxEvent.BoughtAsset = records[i]
		case "BoughtQty":
			taxEvent.BoughtQty = stringToFloat(records[i])
		case "BoughtPriceUSD":
			taxEvent.BoughtPriceUSD = stringToFloat(records[i])
		case "SoldAsset":
			taxEvent.SoldAsset = records[i]
		case "SoldQty":
			taxEvent.SoldQty = stringToFloat(records[i])
		case "SoldPriceUSD":
			taxEvent.SoldPriceUSD = stringToFloat(records[i])
		case "Timestamp":
			taxEvent.Timestamp = stringToUint(records[i])
		}
	}
	return taxEvent
}

func stringToFloat(s string) float32 {
	f, err := strconv.ParseFloat(s, 32)
	if err != nil {
		log.Fatalln("Couldn't parse", err)
	}
	return float32(f)
}

func stringToUint(s string) uint32 {
	f, err := strconv.ParseUint(s, 10, 32)
	if err != nil {
		log.Fatalln("Couldn't parse", err)
	}
	return uint32(f)
}
