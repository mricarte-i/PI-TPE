package queries

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	a "tpe/api/airport"
	f "tpe/api/flights"
)

func MovsPerAirport(filename string, ap a.AirportADT, fl f.FlightADT) {
	var flights int

	// var dataAux a.AirportDataType
	file, ferr := os.Create(filename)
	if ferr != nil {
		log.Fatalf("Error while trying to create or open %s file: %v", filename, ferr)
		os.Exit(1)
	}
	defer file.Close()

	a.ToBeginAirport(ap)
	for a.HasNextAirport(ap) {
		dataAux := a.NextAirport(ap)

		flights = runThroughFlights(dataAux.Icao, fl)
		printMovsPerAirport(file, dataAux, flights)
	}
}

const (
	MOVTYPE_TAKEOFF = "Despegue"
	MOVTYPE_LANDING = "Aterrizaje"
	CLASIFICATION   = "Internacional"
)

func runThroughFlights(icao string, fl f.FlightADT) int {
	flights := 0
	f.ToBeginFlight(fl)
	var dataAux f.FlightDataType

	for f.HasNextFlight(fl) {
		dataAux = f.NextFlight(fl)
		if (strings.Compare(dataAux.IcaoOrig, icao) == 0 &&
			strings.Compare(dataAux.MovType, MOVTYPE_TAKEOFF) == 0) ||
			(strings.Compare(dataAux.IcaoDest, icao) == 0 &&
				strings.Compare(dataAux.MovType, MOVTYPE_LANDING) == 0) {
			flights++
		}
	}
	return flights
}

func printMovsPerAirport(file *os.File, data a.AirportDataType, flights int) {
	if flights > 0 {
		w := bufio.NewWriter(file)
		defer w.Flush()
		fmt.Fprintf(w, "%s;%s;%s;%d\n", data.Icao, data.Local, data.Description, flights)
	}
}

func IntMovsPerAirport(filename string, ap a.AirportADT, fl f.FlightADT) {
	var arrivals, takeOffs int

	file, ferr := os.Create(filename)
	if ferr != nil {
		log.Fatalf("Error while trying to create or open %s file: %v", filename, ferr)
		os.Exit(1)
	}
	defer file.Close()

	a.ToBeginAirport(ap)
	for a.HasNextAirport(ap) {
		data := a.NextAirport(ap)
		arrivals = 0
		takeOffs = 0

		runThroughFlightsInt(data.Icao, fl, &arrivals, &takeOffs)
		printIntMovsPerAirport(file, data, arrivals, takeOffs)
	}
}

func runThroughFlightsInt(icao string, fl f.FlightADT, arrivals *int, takeOffs *int) {
	f.ToBeginFlight(fl)
	for f.HasNextFlight(fl) {
		data := f.NextFlight(fl)

		if ((strings.Compare(data.IcaoOrig, icao) == 0 && strings.Compare(data.MovType, MOVTYPE_TAKEOFF) == 0) || (strings.Compare(data.IcaoDest, icao) == 0 && strings.Compare(data.MovType, MOVTYPE_LANDING) == 0)) && strings.Compare(data.Clasification, CLASIFICATION) == 0 {
			// TODO: refactor this mess✨
			if strings.Compare(data.MovType, MOVTYPE_TAKEOFF) == 0 {
				*takeOffs++
			} else {
				*arrivals++
			}
		}
	}
}

func printIntMovsPerAirport(file *os.File, data a.AirportDataType, arrivals int, takeOffs int) {
	if arrivals > 0 || takeOffs > 0 {
		w := bufio.NewWriter(file)
		defer w.Flush()
		fmt.Fprintf(w, "%s;%s;%d;%d;%d\n", data.Icao, data.Iata, takeOffs, arrivals, takeOffs+arrivals)
	}
}
