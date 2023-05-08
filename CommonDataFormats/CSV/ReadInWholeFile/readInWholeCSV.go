package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
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

func readRecs() ([]record, error) {
	b, err := os.ReadFile("data.csv")
	if err != nil {
		return nil, err
	}
	content := string(b)
	lines := strings.Split(content, "\n") //split on newline
	var records []record
	for i, line := range lines {
		//skip empty lines
		if strings.TrimSpace(line) == "" {
			continue
		}
		var rec record = strings.Split(line, ",")
		if err := rec.validate(); err != nil {
			return nil, fmt.Errorf("entry at line %d was invalid: %w", i, err)
		}
		records = append(records, rec)
	}
	return records, nil
}

func main() {
	records, err := readRecs()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(records)
}
