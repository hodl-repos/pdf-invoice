package utils

import (
	"github.com/jung-kurt/gofpdf"
)

func WriteLetterInfo(pdf *gofpdf.Fpdf) {
	pdf.Ln(17.8)

	// pdf.SetFont("Arial", "", 8)
	// addressStr := "medicloud one OG, Thaliastraße 53/15, 1160 Wien"
	// pdf.CellFormat(colW1, 2.5, mi.trUTF8(addressStr), borderStr, 1, "", false, 0, "")
	// pdf.Ln(9.6)
	// pdf.SetFont("Arial", "", 11)
	// pdf.CellFormat(colW1, 4.2, mi.trUTF8(mi.Invoice.Customer.Name), borderStr, 1, "", false, 0, "")
	// pdf.CellFormat(colW1, 4.2, mi.trUTF8(mi.Invoice.Customer.Street1), borderStr, 1, "", false, 0, "")
	// pdf.CellFormat(colW1, 4.2, fmt.Sprintf("%d %s", mi.Invoice.Customer.Zip, mi.trUTF8(mi.Invoice.Customer.City)), borderStr, 1, "", false, 0, "")
	// pdf.CellFormat(colW1, 4.2, mi.trUTF8(mi.Invoice.Customer.Country), borderStr, 1, "", false, 0, "")

}
