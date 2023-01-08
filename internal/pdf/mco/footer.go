package mco

import "github.com/jung-kurt/gofpdf"

func FooterFunc(pdf *gofpdf.Fpdf, fullWidth float64, borderStr string) func() {
	return func() {
		pdf.SetXY(25, 279)
		pdf.SetTextColor(80, 80, 80)
		pdf.SetFont("Arial", "", 7)
		pdf.CellFormat(fullWidth-5, 2.6, pdf.UnicodeTranslatorFromDescriptor("")("Rotknopf OG, Lindengasse 32, 1070 Wien"), borderStr, 2, "", false, 0, "")
		pdf.CellFormat(fullWidth-5, 2.6, "UniCredit Bank Austria AG, IBAN: AT69 1200 0507 8601 1511, BIC: BKAUATWWXXX", borderStr, 2, "", false, 0, "")
		pdf.CellFormat(fullWidth-5, 2.6, "FN 291617z, Handelsgericht Wien, Firmensitz: Wien, UID-Nr. ATU63549235", borderStr, 2, "", false, 0, "")
	}
}
