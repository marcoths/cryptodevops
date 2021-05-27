package main

import (
	"encoding/csv"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

type Location struct {
	Date               string `json:"date"`
	Location           string `json:"location"`
	TemperatureCelsius int    `json:"temperature_celsius"`
}
type Locations []Location

func csvImport(data io.Reader) (Locations, error) {
	reader := csv.NewReader(data)
	lines, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	locations := Locations{}

	for i, line := range lines {
		//skip header line
		if i == 0 {
			continue
		}
		l := Location{}

		// assuming data is correct
		l.Date = line[0]
		l.Location = line[1]
		temp, _ := strconv.Atoi(line[2])
		l.TemperatureCelsius = temp

		locations = append(locations, l)
	}
	return locations, nil
}

func jsonExport(data Locations) (int, error) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return 0, err
	}

	out, err := os.Create("out.json")
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()

	n, err := out.Write(jsonData)
	if err != nil {
		return 0, err
	}

	return n, nil
}

func main() {
	var input = flag.String("file", "", "input file in CSV format")
	flag.Parse()
	filename := *input
	if filename == "" {
		fmt.Println("Usage: main.go -file")
		flag.PrintDefaults()
		os.Exit(1)
	}

	f, err := os.Open(filename)
	if err != nil {
		log.Fatalf("cannot open %s: %v\n", filename, err)
	}
	defer f.Close()

	locations, err := csvImport(f)
	if err != nil {
		log.Fatal(err)
	}

	bytes, err := jsonExport(locations)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d bytes parsed from csv to json\n", bytes)

}
