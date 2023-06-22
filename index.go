package main

import (
	"fmt"
	"os"
	"strconv"
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
		fmt.Printf("\nHELLO WORLD, param was: %s\n", os.Args[1])
		os.Exit(0)
	}
}
