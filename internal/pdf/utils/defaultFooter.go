package utils

import "github.com/jung-kurt/gofpdf"

func FooterFunc(pdf *gofpdf.Fpdf) func() {
	return func() {
		oldX, oldY := pdf.GetXY()
		_, _, _, bottom := pdf.GetMargins()
		_, height := pdf.GetPageSize()
		_, lineHt := pdf.GetFontSize()

		totalLineHt := 3 * lineHt

		//add one line more as margin -> space between footer and content
		pdf.SetAutoPageBreak(true, bottom+totalLineHt+lineHt)

		pdf.SetXY(oldX, height-bottom-totalLineHt)

		pdf.Write(lineHt, "Rotknopf OG, Lindengasse 32, 1070 Wien")
		pdf.Ln(lineHt)

		pdf.Write(lineHt, "UniCredit Bank Austria AG, IBAN: AT69 1200 0507 8601 1511, BIC: BKAUATWWXXX")
		pdf.Ln(lineHt)

		pdf.Write(lineHt, "FN 291617z, Handelsgericht Wien, Firmensitz: Wien, UID-Nr. ATU63549235")
		pdf.Ln(lineHt)

		//set back to old positions as the footer gets called before everything else
		pdf.SetXY(oldX, oldY)
	}
}
