package main

import (
	"fmt"
	"io"
	"net/http"

	"github.com/jung-kurt/gofpdf"
)

func convertTxtToPdfHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		file, _, err := r.FormFile("file")
		if err != nil {
			http.Error(w, "Failed to read the uploaded file", http.StatusBadRequest)
			return
		}
		defer file.Close()

		text, err := io.ReadAll(file)
		if err != nil {
			http.Error(w, "Failed to read the content of the file", http.StatusInternalServerError)
			return
		}

		pdf := gofpdf.New("P", "mm", "A4", "")
		pdf.AddPage()
		pdf.SetFont("Arial", "B", 16)
		pdf.MoveTo(0, 10)
		pdf.Cell(1, 1, "File")
		pdf.MoveTo(0, 20)
		pdf.SetFont("Arial", "", 14)
		width, _ := pdf.GetPageSize()
		pdf.MultiCell(width, 10, string(text), "", "", false)

		w.Header().Set("Content-Disposition", "attachment; filename=output.pdf")
		w.Header().Set("Content-Type", "application/pdf")
		err = pdf.Output(w)
		if err != nil {
			http.Error(w, "Failed to generate the PDF", http.StatusInternalServerError)
		}
		return
	}

	http.Error(w, "Only POST requests are supported", http.StatusBadRequest)
}

func main() {
	http.HandleFunc("/convert", convertTxtToPdfHandler)
	port := ":8080"
	fmt.Printf("Server started on %s\n", port)
	http.ListenAndServe(port, nil)
}
