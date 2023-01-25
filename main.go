package main

import (
	"fmt"
	"log"
	"time"

	"teamwork/customerimporter"
)

func main() {
	t := time.Now()
	data, err := customerimporter.SortCSV("customers.csv")
	if err != nil {
		log.Println("SortCSV Error: ", err.Error())
	}
	td := time.Since(t)
	fmt.Println("time consumed: ", td)

	fmt.Println(data)
}
