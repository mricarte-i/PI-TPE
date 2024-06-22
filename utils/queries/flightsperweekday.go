package queries

import (
    "bufio"
    "fmt"
    "log"
    "os"

    f "tpe/api/flights"
)

const (
    NONEDAY     = "DOMINGO"
    ONEDAY      = "LUNES"
    TWOSDAY     = "MARTES"
    TREBLESDAY  = "MIERCOLES"
    FOURSDAY    = "JUEVES"
    FIVEDAY     = "VIERNES"
    SIXDAY      = "SABADO"
    CENTURY     = 100
    DAYSINAWEEK = 7
)

type tWeek struct {
    dayName    string
    dayCounter int32
}

func FlightsPerWeekDay(filename string, fl f.FlightADT) {
    file, ferr := os.Create(filename)
    if ferr != nil {
        log.Fatalf("Error while trying to create or open %s file: %v", filename, ferr)
        os.Exit(1)
    }
    defer file.Close()

    var weekDays = []tWeek{
        {
            NONEDAY, 0,
        }, {
            ONEDAY, 0,
        }, {
            TWOSDAY, 0,
        }, {
            TREBLESDAY, 0,
        }, {
            FOURSDAY, 0,
        }, {
            FIVEDAY, 0,
        }, {
            SIXDAY, 0,
        },
    }
    var flData f.FlightDataType

    f.ToBeginFlight(fl)

    for f.HasNextFlight(fl) {
        flData = f.NextFlight(fl)
        weekDays[flData.WeekDay].dayCounter++

    }
    printFlightsPerWeek(file, weekDays)

}

func printFlightsPerWeek(file *os.File, weekDays []tWeek) {
    w := bufio.NewWriter(file)
    defer w.Flush()
    for i := 1; i < DAYSINAWEEK; i++ {
        fmt.Fprintf(w, "%s;%d\n", weekDays[i%DAYSINAWEEK].dayName, weekDays[i%DAYSINAWEEK].dayCounter)
    }
}
