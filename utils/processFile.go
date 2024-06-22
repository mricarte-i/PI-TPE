package utils

import (
    "bufio"
    "log"
    "os"

    a "tpe/models/airport"
    f "tpe/models/flights"
)

func ProcessAirports(filename string) a.AirportADT {
    file, ferr := os.OpenFile(filename, os.O_RDONLY, os.ModePerm)
    if ferr != nil {
        log.Fatalf("Error while trying to open airport file: %v", ferr)
        os.Exit(1)
    }
    defer file.Close()

    err := false
    ap := a.NewAirport()
    // if creation could fail, exit

    sc := bufio.NewScanner(file)
    sc.Scan() // skip first row, its just names and stuff
    for sc.Scan() {
        // REFACTOR: insert airport data into ap
        if !a.InsertAirtport(ap, sc.Text()) {
            err = true
        }
    }

    if err {
        log.Print("Some of the lines in the airport file weren't copied!")
    }

    return ap
}

func ProcessFlights(filename string, year int) f.FlightADT {
    file, ferr := os.OpenFile(filename, os.O_RDONLY, os.ModePerm)
    if ferr != nil {
        log.Fatalf("Error while trying to open flights file: %v", ferr)
        os.Exit(1)
    }
    defer file.Close()

    fl := f.NewFlight(year)
    // if creation could fail, exit

    err := false
    sc := bufio.NewScanner(file)
    sc.Scan() // skip first row, its just names and stuff
    for sc.Scan() {
        if f.InsertFlight(fl, sc.Text()) {
            err = true
        }
    }

    if err {
        log.Print("Some of the lines in the flights file weren't copied!")
    }

    return fl
}
