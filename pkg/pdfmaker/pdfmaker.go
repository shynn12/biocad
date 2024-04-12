package pdfmaker

import (
	"fmt"
	"os"

	"github.com/jung-kurt/gofpdf"
	"github.com/shynn12/biocad/internal/config"
)

func MakePDF(headers []string, field []string, cfg *config.Config) error {
	pwd, _ := os.Getwd()
	pdf := gofpdf.New("P", "mm", "A4", pwd+"/font/")
	pdf.AddPage()
	pdf.AddFont("Helvetica", "", "helvetica_1251.json")
	pdf.SetCompression(false)
	tr := pdf.UnicodeTranslatorFromDescriptor("cp-1251")
	for c, i := range headers {
		pdf.SetFont("Helvetica", "B", 16)
		pdf.CellFormat(40, 7, i, "1", 0, "", false, 0, "")
		pdf.SetFont("Helvetica", "", 12)
		pdf.CellFormat(200, 7, tr(field[c]), "1", 0, "", false, 0, "")
		pdf.Ln(-1)
	}

	pdf.Ln(-1)
	//File name like item.Number_item.guid due to reapiting value of the latter
	err := pdf.OutputFileAndClose(fmt.Sprintf("%s/%s_%s.pdf", cfg.Pdfpath, field[0], field[3]))
	if err != nil {
		return fmt.Errorf("cannot create file due to error: %v", err)
	}

	return nil
}
