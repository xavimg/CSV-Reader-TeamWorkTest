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

type DataImproved struct {
	Domain string
	Count  int
}

func SortCSVImproved(fp string) ([]Data, error) {
	file, err := os.Open(fp)
	if err != nil {
		log.Println("OpenFile ERROR: ", err.Error())
		return nil, err
	}
	defer file.Close()

	bufferSize := 4096
	reader := bufio.NewReaderSize(file, bufferSize)

	dc, err := domainCounterImproved(reader)
	if err != nil {
		log.Println("DomainCounter ERROR: ", err.Error())
		return nil, err
	}

	s := mapToSlice(dc)
	sort.Slice(s, func(i, j int) bool {
		return s[i].Count > s[j].Count
	})

	return s, nil
}

// DomainCounter scans the buffer file and executed the logic to count the repeated domains.
// var domainCounter; is a map that will help us to count all repeated domains.
func domainCounterImproved(reader *bufio.Reader) (map[string]int, error) {
	dc := make(map[string]int)
	for {
		record, err := reader.ReadString('\n')
		if err != nil && err.Error() != "EOF" {
			return nil, err
		}
		if err != nil && err.Error() == "EOF" && record == "" {
			break
		}

		_, after, _ := strings.Cut(record, "@")
		atComma := strings.Index(after, ",")
		if atComma == -1 {
			continue
		}

		dc[after[:atComma]]++
	}

	return dc, nil
}

// MapToSlice transforms our map to a slice of Data{}.
func mapToSliceImproved(dc map[string]int) []Data {
	var data []Data
	for k, v := range dc {
		data = append(data, Data{k, v})
	}

	return data
}
