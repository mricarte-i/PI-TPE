package utils

import (
	"bufio"
	"log"
	"os"

	api "tpe/api"
)

func ProcessAirports(filename string) api.AirportADT {
	file, ferr := os.OpenFile(filename, os.O_RDONLY, os.ModePerm)
	if ferr != nil {
		log.Fatalf("Error while trying to open airport file: %v", ferr)
		os.Exit(1)
	}
	defer file.Close()

	err, added := false, false
	ap := api.NewAirport()

	sc := bufio.NewScanner(file)
	for sc.Scan() {
		// REFACTOR: insert airport data into ap
		if api.InsertAirtport(ap, sc.Text(), &added) && added {
			err = true
		}
	}

	if err {
		log.Print("Some of the lines in the airport file weren't copied!")
	}

	return ap
}
