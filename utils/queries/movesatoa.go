package queries

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "strings"

    a "tpe/models/airport"
    f "tpe/models/flights"
)

type apChild struct {
    nextChild *apChild
    childIcao string
    arrivals  uint
    takeOffs  uint
}
type childNode *apChild

func insertChild(n childNode, childIcao string, movType string) childNode {
    cmp := 0
    if n != nil {
        cmp = strings.Compare(childIcao, n.childIcao)
    }

    if n == nil || cmp < 0 {
        newChild := &apChild{}
        newChild.childIcao = childIcao

        if movType == ARRIVAL {
            newChild.arrivals++
        } else {
            newChild.takeOffs++
        }

        newChild.nextChild = n
        return newChild

    } else if cmp == 0 {
        if movType == ARRIVAL {
            n.arrivals++
        } else {
            n.takeOffs++
        }
    } else {
        n.nextChild = insertChild(n.nextChild, childIcao, movType)
    }

    return n
}

func lookAtoA(file *os.File, icaoAp string, fl f.FlightADT) {
    f.ToBeginFlight(fl)
    var data f.FlightDataType
    var first childNode = nil
    for f.HasNextFlight(fl) {
        data = f.NextFlight(fl)
        if strings.Compare(icaoAp, data.IcaoOrig) == 0 && strings.Compare(data.MovType, MOVTYPE_TAKEOFF) == 0 {
            first = insertChild(first, data.IcaoDest, TAKEOFF)
        } else if strings.Compare(icaoAp, data.IcaoDest) == 0 && strings.Compare(data.MovType, MOVTYPE_LANDING) == 0 {
            first = insertChild(first, data.IcaoOrig, ARRIVAL)
        }
    }

    w := bufio.NewWriter(file)
    defer w.Flush()
    printMovesAtoA(icaoAp, first, w)
}

func printMovesAtoA(icaoAp string, n childNode, w *bufio.Writer) {
    if n != nil {
        fmt.Fprintf(w, "%s;%s;%d;%d\n", icaoAp, n.childIcao, n.takeOffs, n.arrivals)
        printMovesAtoA(icaoAp, n.nextChild, w)
    }
}

func MovesAtoA(filename string, fl f.FlightADT, ap a.AirportADT) {
    file, ferr := os.Create(filename)
    if ferr != nil {
        log.Fatalf("Error while trying to create or open %s file: %v", filename, ferr)
        os.Exit(1)
    }
    defer file.Close()

    var data a.AirportDataType
    a.ToBeginAirport(ap)
    for a.HasNextAirport(ap) {
        data = a.NextAirport(ap)
        lookAtoA(file, data.Icao, fl)
    }
}
