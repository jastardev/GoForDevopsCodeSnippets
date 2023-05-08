package main

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"os"
)

// To use the encoding/csv package for reading and writing csv files, the
// data must conform to RFC4180.

type record []string

func (r record) validate() error {
	if len(r) != 2 {
		return errors.New("data format is incorrect")
	}
	return nil
}

func (r record) first() string {
	return r[0]
}

func (r record) last() string {
	return r[1]
}

func (r record) string() string {
	return r.first() + " " + r.last()
}

func readRecs() ([]record, error) {
	file, err := os.Open("data.csv")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = 2
	reader.TrimLeadingSpace = true

	var recs []record
	for {
		data, err := reader.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}
		rec := record(data)
		recs = append(recs, rec)
	}
	return recs, nil
}

func main() {
	records, err := readRecs()
	if err != nil {
		panic(err)
	}
	for _, rec := range records {
		fmt.Println(rec.string())
	}
}
