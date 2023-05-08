package main

import (
	"encoding/csv"
	"errors"
	"os"
)

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

func writeCSV(records []record) error {

	file, err := os.OpenFile("data.csv", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	w := csv.NewWriter(file)
	defer w.Flush()
	for _, rec := range records {
		if err := w.Write(rec); err != nil {
			return err
		}
	}
	return nil
}

func main() {
	var records = []record{
		record{"John", "Adams"},
		record{"Jacob", "Astar"},
	}
	if err := writeCSV(records); err != nil {
		panic(err)
	}
}
