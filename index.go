package main

import (
	"fmt"
	"os"
	"strconv"

	process "tpe/utils"
)

func yearIsValid(yearStr string) bool {
	year, err := strconv.Atoi(yearStr)
	if err != nil && year >= 2014 && year <= 2018 {
		return true
	}
	return false
}

func main() {
	if len(os.Args) != 2 || yearIsValid(os.Args[1]) {
		fmt.Printf("\nExpected 1 parameter for the queried year.\nThe year should be in the range of [2014, 2018].\nAbort\n")
		os.Exit(1)
	} else {
		fmt.Printf("\nQueries from year: %s\nStarting files processing...\n", os.Args[1])
		// TODO: processors & structs
		ap := process.ProcessAirports("assets/aeropuertos_detalle.csv")
		fl := processFlights("assets/eana1401-1802.csv")
		fmt.Printf("Processing: READY\n")

		os.Exit(0)
	}
}
