package main

import (
	"fmt"
	"log"
	"time"

	"teamwork/customerimporter"
)

func main() {
	timeStart := time.Now()
	data, err := customerimporter.SortCSV("customers.csv")
	if err != nil {
		log.Println("SortCSV Error: ", err.Error())
	}
	timeNeeded := time.Since(timeStart)
	fmt.Println("time consumed: ", timeNeeded)

	fmt.Println(data)
}
