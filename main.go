package main

import (
	"fmt"
	"github.com/jung-kurt/gofpdf"
)

func main() {
	build()
}

func build() {
	pdf := gofpdf.New(gofpdf.OrientationPortrait, gofpdf.UnitPoint, gofpdf.PageSizeA4, "")
	width, height := pdf.GetPageSize()

	fmt.Printf("width=%v, height=%v\n", width, height)
	pdf.AddPage()
	pdf.SetFont("arial", "B", 14)

	_, lineHt := pdf.GetFontSize()

	pdf.Cell(0, lineHt, "Hello")
	pdf.Cell(0, lineHt, "My name is Andre")
	err := pdf.OutputFileAndClose("resume.pdf")

	if err != nil {
		panic(err)
	}
}

func header() {

}
