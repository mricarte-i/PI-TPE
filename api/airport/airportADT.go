package airport

import (
	"strings"
)

/* PRIVATE DEFINITIONS */
type tAirportNode *airportNode

type airportDataType struct {
	local       string
	icao        string
	iata        string
	kind        string // type is a GO keyword...
	description string
	condition   string
	traffic     string
}

type airportNode struct {
	data airportDataType
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
type AirportADT *airportCDT

func NewAirport() AirportADT {
	return &airportCDT{}
}

func InsertAirtport(ap AirportADT, data string, added *bool) bool {
	newData := toAirportDataType(data)

	if len(newData.icao) > 0 {
		// In GO, `->` and `.` are both represented by `.`
		ap.first = insertAirportRecc(ap.first, newData, added)
		return true
	}
	return false
}

/* INTERNAL FUNCTIONS */

func insertAirportRecc(n tAirportNode, data airportDataType, added *bool) tAirportNode {
	if n == nil || strings.Compare(data.icao, n.data.icao) < 0 {
		newNode := &airportNode{data: data, tail: n}
		// if using malloc and fails, added set to false, return
		*added = true
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

func toAirportDataType(formattedText string) airportDataType {
	// full refactor
	// formattedText is a row from the csv containing data for each category, in order, separated by ";"
	fields := strings.Split(formattedText, SEPARATOR)
	return airportDataType{
		local:       fields[local],
		icao:        fields[icao],
		iata:        fields[iata],
		kind:        fields[kind], // type
		description: fields[name],
		condition:   fields[condition],
		traffic:     fields[traffic],
	}
}

func replaceAtIndex(in string, r rune, idx int) string {
	out := []rune(in)
	out[idx] = r
	return string(out)
}
