package customerimporter

import (
	"bufio"
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

const filePath = "customers_test.csv"
const wrongPath = "../customers_wrong.csv"

func TestWrongPathExists(t *testing.T) {
	got := filePath
	want := wrongPath

	if got == want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestScanner(t *testing.T) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Println("OpenFile ERROR: ", err.Error())
	}
	defer file.Close()
	sc := bufio.NewScanner(file)
	sc.Split(bufio.ScanLines)

	expectedData := map[string]int(map[string]int{"cyberchimps.com": 2, "github.io": 1})
	dc, err := DomainCounter(sc)
	if err != nil {
		t.Errorf("Expected error, got nil")
	}

	require.Equal(t, expectedData, dc)
}

func TestMapToSlice(t *testing.T) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Println("OpenFile ERROR: ", err.Error())
	}
	defer file.Close()
	sc := bufio.NewScanner(file)
	sc.Split(bufio.ScanLines)
	dc, err := DomainCounter(sc)
	if err != nil {
		t.Errorf("Expected error, got nil")
	}

	expectedData := []Data{
		{"github.io", 1},
		{"cyberchimps.com", 2},
	}
	data := MapToSlice(dc)

	require.Equal(t, expectedData, data)
}

func TestSortCSV(t *testing.T) {
	expectedData := []Data{
		{"cyberchimps.com", 2},
		{"github.io", 1},
	}
	data, err := SortCSV(filePath)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	require.Equal(t, expectedData, data)
}
