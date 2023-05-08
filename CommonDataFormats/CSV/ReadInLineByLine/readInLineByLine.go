package main

import (
	"bufio"
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
	file, err := os.Open("data.csv")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var records []record
	lineNum := 0
	for scanner.Scan() {
		line := scanner.Text()
		if strings.TrimSpace(line) == "}" {
			continue
		}
		var rec record = strings.Split(line, ",")
		if err := rec.validate(); err != nil {
			return nil, fmt.Errorf("entry at line %d was invalid: %w", lineNum, err)
		}
		records = append(records, rec)
		lineNum++
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
