package queries

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "strings"

    a "tpe/models/airport"
    f "tpe/models/flights"
    l "tpe/models/list"
)

func processIcaos(fl f.FlightADT) l.ListADT {
    var icaos = l.NewList()
    var aux f.FlightDataType
    f.ToBeginFlight(fl)
    for f.HasNextFlight(fl) {
        aux = f.NextFlight(fl)
        l.Insert(icaos, aux.IcaoOrig)
        l.Insert(icaos, aux.IcaoDest)
    }

    return icaos
}

func lookAtoA(file *os.File, icaoAp string, otherIcaos l.ListADT, fl f.FlightADT) {
    var arrivals, takeOffs uint
    l.ToBegin(otherIcaos)
    for l.HasNext(otherIcaos) {
        aux := l.Next(otherIcaos)
        arrivals, takeOffs = 0, 0
        runThroughAtoAFlights(icaoAp, aux, &arrivals, &takeOffs, fl)

        printMoveAtoA(icaoAp, aux, arrivals, takeOffs, file)
    }
}

func runThroughAtoAFlights(icaoAp string, aux string, arrivals *uint, takeOffs *uint, fl f.FlightADT) {
    var data f.FlightDataType
    f.ToBeginFlight(fl)

    for f.HasNextFlight(fl) {
        data = f.NextFlight(fl)
        if strings.Compare(icaoAp, data.IcaoOrig) == 0 && strings.Compare(aux, data.IcaoDest) == 0 {
            *takeOffs++
        }
        if strings.Compare(aux, data.IcaoOrig) == 0 && strings.Compare(icaoAp, data.IcaoDest) == 0 {
            *arrivals++
        }
    }
}

func printMoveAtoA(icaoAp string, icaoOtherAp string, arrivals uint, takeOffs uint, file *os.File) {
    w := bufio.NewWriter(file)
    defer w.Flush()
    if arrivals > 0 || takeOffs > 0 {
        fmt.Fprintf(w, "%s;%s;%d:%d\n", icaoAp, icaoOtherAp, takeOffs, arrivals)
    }
}

func MovesAtoA(filename string, fl f.FlightADT, ap a.AirportADT) {
    file, ferr := os.Create(filename)
    if ferr != nil {
        log.Fatalf("Error while trying to create or open %s file: %v", filename, ferr)
        os.Exit(1)
    }
    defer file.Close()

    var icaos = processIcaos(fl)
    var data a.AirportDataType
    a.ToBeginAirport(ap)
    for a.HasNextAirport(ap) {
        data = a.NextAirport(ap)
        lookAtoA(file, data.Icao, icaos, fl)
    }
}
