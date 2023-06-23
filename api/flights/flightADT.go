package flights

import (
	"strconv"
	"strings"
	"unicode"
)

/* PRIVATE DEFINITIONS */

type tFlightNode *flightNode

type flightDataType struct {
	date          dateType
	clasification string
	movType       string
	icaoOrig      string
	icaoDest      string
	weekDay       int
}

type dateType struct {
	day   int
	month int
	year  int
}

type flightNode struct {
	data flightDataType
	tail tFlightNode
}

type flightCDT struct {
	yearSelected int
	first        tFlightNode
	iter         tFlightNode
}

/* PUBLIC STUFF */

type FlightADT *flightCDT

func NewFlight(year int) FlightADT {
	return &flightCDT{yearSelected: year}
}

func InsertFlight(f FlightADT, data string, added *bool) bool {
	newData := toFlightDataType(data, added)

	if !*added {
		return false
	}

	if newData.date.year == f.yearSelected {
		newNode := &flightNode{}
		// if newNode failed, freeFlightData(newData) + added = false + return true

		*added = true
		newNode.data = newData
		newNode.tail = f.first
		f.first = newNode
		return true
	}

	*added = false
	// freeFlightData(newData)
	return false
}

/* INTERNAL FUNCTIONS  */
const (
	SEPARATOR      = ";"
	MAX_ICAO       = 4
	NOT_APPLICABLE = "N/A"
	SEPARATOR_DATE = "/"
	CENTURY        = 100
)

type field int

const (
	date field = iota
	time
	class
	clasification
	mov_type
	icaoOrig
	icaoDest
	airline
	airship
	apc_airship
)

func toFlightDataType(formattedText string, err *bool) flightDataType {
	fields := strings.Split(formattedText, SEPARATOR)

	fields[icaoOrig] = validateIcao(fields[icaoOrig])
	fields[icaoDest] = validateIcao(fields[icaoDest])

	newData := putFlightData(fields, err)
	return newData
}

func validateIcao(s string) string {
	for i, c := range s {
		if unicode.IsDigit(c) || i >= MAX_ICAO {
			return NOT_APPLICABLE
		}
	}
	return s
}

func putFlightData(fields []string, err *bool) flightDataType {
	data := flightDataType{}
	// errorAns := false defineFlightField was kinda useless so this also gets removed...
	*err = false

	data.date = defineDate(fields[date])
	data.weekDay = determineDay(data.date)

	data.clasification = fields[clasification]
	data.movType = fields[mov_type]
	data.icaoOrig = fields[icaoOrig]
	data.icaoDest = fields[icaoDest]

	return data
}

type dateField int

const (
	day dateField = iota
	month
	year
)

func defineDate(date string) dateType {
	dateParts := strings.Split(date, SEPARATOR_DATE)

	day, _ := strconv.Atoi(dateParts[day])
	month, _ := strconv.Atoi(dateParts[month])
	year, _ := strconv.Atoi(dateParts[year])

	return dateType{day, month, year}
}

func determineDay(date dateType) int {
	var c, g int
	d := date.day
	m := date.month
	y := date.year
	eLookUpTable := []int{0, 3, 2, 5, 0, 3, 5, 1, 4, 6, 2, 4}
	fLookUpTable := []int{0, 5, 3, 1}

	if m < 3 {
		y--
	}
	c = y / CENTURY
	g = y - (CENTURY * c)

	return (d + eLookUpTable[m-1] + fLookUpTable[c%4] + g + (g / 4)) % 7
}
