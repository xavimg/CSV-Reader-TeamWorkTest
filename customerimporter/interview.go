// package customerimporter reads from the given customers.csv file and returns a
// sorted (data structure of your choice) of email domains along with the number
// of customers with e-mail addresses for each domain.  Any errors should be
// logged (or handled). Performance matters (this is only ~3k lines, but *could*
// be 1m lines or run on a small machine).
package customerimporter

import (
	"bufio"
	"log"
	"os"
	"sort"
	"strings"
)

type Data struct {
	Domain string
	Count  int
}

func SortCSV(filePath string) ([]Data, error) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Println("OpenFile ERROR: ", err.Error())
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	domainCounter, err := Scanner(scanner)
	if err != nil {
		log.Println("Scanner ERROR: ", err.Error())
		return nil, err
	}

	sliceData := MapToSlice(domainCounter)
	sort.Slice(sliceData, func(i, j int) bool {
		return sliceData[i].Count > sliceData[j].Count
	})

	return sliceData, nil
}

// Scanner scans the scanner.
// var domainCounter; is a map that will help us to count all repeated domains.
func Scanner(scanner *bufio.Scanner) (map[string]int, error) {
	domainCounter := make(map[string]int)
	for scanner.Scan() {
		record := scanner.Text()
		_, after, _ := strings.Cut(record, "@")
		atComma := strings.Index(after, ",")
		if atComma == -1 {
			continue
		}

		domainCounter[after[:atComma]]++
	}

	return domainCounter, nil
}

// MapToSlice transforms our map to a slice of Data{}.
func MapToSlice(dc map[string]int) []Data {
	var data []Data
	for k, v := range dc {
		data = append(data, Data{k, v})
	}

	return data
}
