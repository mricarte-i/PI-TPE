package airport

import (
    "strings"
)

/* PRIVATE DEFINITIONS */
type tAirportNode *airportNode

type airportNode struct {
    data AirportDataType
    tail tAirportNode
}

type airportCDT struct {
    first tAirportNode
    iter  tAirportNode
}

const (
    MAXFIELDS_AIRPORT = 23
    DELIM             = ";"
    SEPARATOR         = ";"
)

/* PUBLIC STUFF */
type AirportDataType struct {
    Local       string
    Icao        string
    Iata        string
    Kind        string // type is a GO keyword...
    Description string
    Condition   string
    Traffic     string
}

type AirportADT *airportCDT

func NewAirport() AirportADT {
    return &airportCDT{}
}

func InsertAirtport(ap AirportADT, data string) bool {
    newData := toAirportDataType(data)
    added := true

    if len(newData.Icao) > 0 {
        // In GO, `->` and `.` are both represented by `.`
        ap.first = insertAirportRecc(ap.first, newData, &added)
    }
    return added
}

/*
TODO:
  - refactor ToBeginAirport, HasNextAirport, NextAirport to be methods of (ap AirportADT), not take it as param
*/

func ToBeginAirport(ap AirportADT) {
    ap.iter = ap.first
}

func HasNextAirport(ap AirportADT) bool {
    return ap.iter != nil
}

func NextAirport(ap AirportADT) AirportDataType {
    data := ap.iter.data
    ap.iter = ap.iter.tail
    return data
}

/* INTERNAL FUNCTIONS */

func insertAirportRecc(n tAirportNode, data AirportDataType, added *bool) tAirportNode {
    if n == nil || strings.Compare(data.Icao, n.data.Icao) < 0 {
        newNode := &airportNode{data: data, tail: n}
        // if using malloc and fails, added set to false, return
        *added = false
        return newNode
    }

    n.tail = insertAirportRecc(n.tail, data, added)
    return n
}

type field int

const (
    local field = iota
    icao
    iata
    kind // type
    name
    coordinates
    latitude
    longitude
    elev
    uom_elem
    ref
    distance_ref
    direction_ref
    condition
    control
    region
    fir
    uso
    traffic
    sna
    consesionate
    state
    inhab
)

func toAirportDataType(formattedText string) AirportDataType {
    // full refactor
    // formattedText is a row from the csv containing data for each category, in order, separated by ";"
    fields := strings.Split(formattedText, SEPARATOR)
    return AirportDataType{
        Local:       fields[local],
        Icao:        fields[icao],
        Iata:        fields[iata],
        Kind:        fields[kind], // type
        Description: fields[name],
        Condition:   fields[condition],
        Traffic:     fields[traffic],
    }
}

func replaceAtIndex(in string, r rune, idx int) string {
    out := []rune(in)
    out[idx] = r
    return string(out)
}
