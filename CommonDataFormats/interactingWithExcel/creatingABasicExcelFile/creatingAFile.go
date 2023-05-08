package main

import (
	"github.com/xuri/excelize/v2"
)

func main() {
	const sheet = "Sheet1"
	xlsx := excelize.NewFile()
	xlsx.SetCellValue(sheet, "A1", "Server Name")
	xlsx.SetCellValue(sheet, "B1", "Generation")
	xlsx.SetCellValue(sheet, "C1", "Acquisition Date")
	xlsx.SetCellValue(sheet, "D1", "CPU Vendor")

	xlsx.SetCellValue(sheet, "A2", "svlaa01")
	xlsx.SetCellValue(sheet, "B2", "12")
	xlsx.SetCellValue(sheet, "C2", "10/12/2020")
	xlsx.SetCellValue(sheet, "D2", "Intel")

	xlsx.SetCellValue(sheet, "A3", "svlaa05")
	xlsx.SetCellValue(sheet, "B3", "13")
	xlsx.SetCellValue(sheet, "C3", "07/24/2022")
	xlsx.SetCellValue(sheet, "D3", "Intel")

	if err := xlsx.SaveAs("./book1.xlsx"); err != nil {
		panic(err)
	}
}
