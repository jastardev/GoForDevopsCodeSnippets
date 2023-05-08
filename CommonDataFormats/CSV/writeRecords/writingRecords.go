package main

import (
	"bytes"
	"errors"
	"os"
	"sort"
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
func (r record) csv() []byte {
	b := bytes.Buffer{}
	for _, field := range r {
		b.WriteString(field + ",")
	}
	b.WriteString("\n")
	return b.Bytes()
}

func writeRecs(recs []record) error {
	file, err := os.OpenFile("data-sorted.csv", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	//Sort by last name
	sort.Slice(
		recs,
		func(i, j int) bool {
			return recs[i].last() < recs[j].last()
		})

	for _, rec := range recs {
		_, err := file.Write(rec.csv())
		if err != nil {
			return err
		}
	}
	return nil
}

func main() {
	var records []record
	records = append(records, record{"Yvette", "Astar"})
	records = append(records, record{"Jacob", "Astar"})
	writeRecs(records)
}
