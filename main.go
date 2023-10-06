package main

import (
	"fmt"
	"io/ioutil"

	"github.com/jung-kurt/gofpdf"
)

func main() {
	text, err := ioutil.ReadFile("lorem.txt")
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 16)
	pdf.MoveTo(0, 10)
	pdf.Cell(1, 1, "Lorem Ipsum")
	pdf.MoveTo(0, 20)
	pdf.SetFont("Arial", "", 14)
	width, _ := pdf.GetPageSize()
	pdf.MultiCell(width, 10, string(text), "", "", false)
	err = pdf.OutputFileAndClose("hello.pdf")
	if err == nil {
		fmt.Println("PDF generated successfully")
	}
}
