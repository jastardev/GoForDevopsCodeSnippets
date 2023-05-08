package main

import (
	"errors"
	"github.com/xuri/excelize/v2"
	"strconv"
	"sync"
	"time"
)

type CPUVendor string

const (
	UnknownCPUVendor CPUVendor = "Unknown"
	Intel            CPUVendor = "Intel"
	AMD              CPUVendor = "AMD"
)

var validCPUVendors = map[CPUVendor]bool{
	Intel: true,
	AMD:   true,
}

type serverSheet struct {
	mu        sync.Mutex
	sheetName string
	xlsx      *excelize.File
	nextRow   int
}

func (s *serverSheet) add(name string, gen int, acquisitionDate time.Time, vendor CPUVendor) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if name == "" {
		return errors.New("name cannot be clank")
	}
	if gen < 1 || gen > 13 {
		return errors.New("gen was not in acceptable range")
	}
	if acquisitionDate.IsZero() {
		return errors.New("acquisition date must be set")
	}
	if !validCPUVendors[vendor] {
		return errors.New("vendor is not valid")
	}

	s.xlsx.SetCellValue(s.sheetName, "A"+strconv.Itoa(s.nextRow), name)
	s.xlsx.SetCellValue(s.sheetName, "B"+strconv.Itoa(s.nextRow), gen)
	s.xlsx.SetCellValue(s.sheetName, "C"+strconv.Itoa(s.nextRow), acquisitionDate)
	s.xlsx.SetCellValue(s.sheetName, "D"+strconv.Itoa(s.nextRow), vendor)
	s.nextRow++
	return nil
}

func newServerSheet() (*serverSheet, error) {
	s := &serverSheet{
		sheetName: "Sheet1",
		xlsx:      excelize.NewFile(),
		nextRow:   2,
	}
	s.xlsx.SetCellValue(s.sheetName, "A1", "Server Name")
	s.xlsx.SetCellValue(s.sheetName, "B1", "Generation")
	s.xlsx.SetCellValue(s.sheetName, "C1", "Acquisition Date")
	s.xlsx.SetCellValue(s.sheetName, "D1", "CPU Vendor")
	return s, nil
}

func main() {

}
