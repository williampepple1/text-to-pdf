package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/jung-kurt/gofpdf"
)

func convertTxtToPdfHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		if err := r.ParseMultipartForm(10 << 20); err != nil {
			http.Error(w, "File too large or invalid form", http.StatusBadRequest)
			return
		}

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
		pdf.SetY(10)
		pdf.Cell(40, 10, "File")
		pdf.SetY(25)
		pdf.SetFont("Arial", "", 14)
		width, _ := pdf.GetPageSize()
		usableWidth := width - 20
		pdf.MultiCell(usableWidth, 10, string(text), "", "", false)

		var buf bytes.Buffer
		if err := pdf.Output(&buf); err != nil {
			http.Error(w, "Failed to generate the PDF", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Disposition", "attachment; filename=output.pdf")
		w.Header().Set("Content-Type", "application/pdf")
		w.Write(buf.Bytes())
		return
	}

	http.Error(w, "Only POST requests are supported", http.StatusBadRequest)
}

func main() {
	http.HandleFunc("/convert", convertTxtToPdfHandler)
	port := ":8080"
	fmt.Printf("Server started on %s\n", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal(err)
	}
}
