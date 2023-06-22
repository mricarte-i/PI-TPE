package utils

import (
	"bufio"
	"log"
	"os"
)

// TODO: airport struct ...
func processAirports(filename string) AirportADT {
	file, ferr := os.OpenFile(filename, os.O_RDONLY, os.ModePerm)
	if ferr != nil {
		log.Fatalf("Error while trying to open airport file: %v", ferr)
		os.Exit(1)
	}
	defer file.Close()

	err, added := false, false
	// TODO: airport struct...
	ap := NewAirport()

	sc := bufio.NewScanner(file)
	for sc.Scan() {
		// TODO: insert airport data into ap
		if insertAirtport(ap, sc.Text(), &added) && added != false {
			err = true
		}
	}

	if err {
		log.Print("Some of the lines in the airport file weren't copied!")
	}

	return ap
}
