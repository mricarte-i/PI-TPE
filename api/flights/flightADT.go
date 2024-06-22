package flights

import (
    "strconv"
    "strings"
    "unicode"
)

/* PRIVATE DEFINITIONS */

type tFlightNode *flightNode

type dateType struct {
    day   int
    month int
    year  int
}

type flightNode struct {
    data FlightDataType
    tail tFlightNode
}

type flightCDT struct {
    yearSelected int
    first        tFlightNode
    iter         tFlightNode
}

/* PUBLIC STUFF */
type FlightDataType struct {
    Date          dateType
    Clasification string
    MovType       string
    IcaoOrig      string
    IcaoDest      string
    WeekDay       int
}

type FlightADT *flightCDT

func NewFlight(year int) FlightADT {
    return &flightCDT{yearSelected: year}
}

func InsertFlight(f FlightADT, data string, added *bool) bool {
    newData := toFlightDataType(data)

    if !*added {
        return false
    }

    if newData.Date.year == f.yearSelected {
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

/*
TODO:
  - refactor ToBeginFlight, HasNextFlight, NextFlight to be methods of (f FlightADT), not take it as param
*/
func ToBeginFlight(f FlightADT) {
    f.iter = f.first
}

func HasNextFlight(f FlightADT) bool {
    return f.iter != nil
}

func NextFlight(f FlightADT) FlightDataType {
    data := f.iter.data
    f.iter = f.iter.tail
    return data
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

func toFlightDataType(formattedText string) FlightDataType {
    fields := strings.Split(formattedText, SEPARATOR)

    fields[icaoOrig] = validateIcao(fields[icaoOrig])
    fields[icaoDest] = validateIcao(fields[icaoDest])

    newData := putFlightData(fields)
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

func putFlightData(fields []string) FlightDataType {
    data := FlightDataType{}

    data.Date = defineDate(fields[date])
    data.WeekDay = determineDay(data.Date)

    data.Clasification = fields[clasification]
    data.MovType = fields[mov_type]
    data.IcaoOrig = fields[icaoOrig]
    data.IcaoDest = fields[icaoDest]

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
