package main

import (
    "fmt"
    "os"
    "strconv"

    process "tpe/utils"
    query "tpe/utils/queries"
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
        year, _ := strconv.Atoi(os.Args[1])
        fmt.Printf("\nQueries from year: %d\nStarting files processing...\n", year)

        ap := process.ProcessAirports("assets/aeropuertos_detalle.csv")
        fl := process.ProcessFlights("assets/eana1401-1802.csv", year)
        fmt.Printf("Processing: READY\nStarting queries...\n")

        query.MovsPerAirport("movs_aeropuertos.csv", ap, fl)
        fmt.Printf("Query #1: DONE.\n")

        query.IntMovsPerAirport("movs_internacional.csv", ap, fl)
        fmt.Printf("Query #2: DONE.\n")

        query.FlightsPerWeekDay("semanal.csv", fl)
        fmt.Printf("Query #3: DONE.\n")

        query.MovesAtoA("aerop_detalle.csv", fl, ap)
        fmt.Printf("Query #4: DONE.\n")

        fmt.Printf("Finished program.\n")
        os.Exit(0)
    }
}
