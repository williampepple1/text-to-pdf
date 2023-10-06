package main

import (
	"fmt"

	"github.com/jung-kurt/gofpdf"
)

func main() {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 16)
	pdf.MoveTo(0, 10)
	pdf.Cell(1, 1, "Hello world")
	err := pdf.OutputFileAndClose("hello.pdf")
	if err == nil {
		fmt.Println("PDF generated successfully")
	}
}
