package document

import "github.com/jung-kurt/gofpdf"

// DrawVerticalLineCustom draws a vertical line on the current y position from
// x1 to x2 with a given thinkness.
func DrawVerticalLineCustom(pdf *gofpdf.Fpdf, x1, x2, thickness float64) {
	pdf.SetLineWidth(thickness)

	y := pdf.GetY()
	pdf.Line(x1, y, x2, y)
}

// DrawVerticalLinePrintWidth draws a vertical line inside the current print
// width with a thickness of Fpdf.GetLineWidth().
func DrawVerticalLinePrintWidth(pdf *gofpdf.Fpdf) {
	ml, _, _, _ := pdf.GetMargins()
	w := GetPrintWidth(pdf)
	DrawVerticalLineCustom(pdf, ml, ml+w, pdf.GetLineWidth())
}
