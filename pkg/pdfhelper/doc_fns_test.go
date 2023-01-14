package pdfhelper

import "testing"

func TestMCell(t *testing.T) {
	doc := NewDocA4()

	doc.SetX(100)
	x, y := doc.GetXY()
	doc.MCell(20, 5, "extra-long string which is too long", "1", "L", false)
	cellHt := doc.GetY() - y
	doc.SetXY(x+20, y)
	// doc.Rect(x+20, y, 20, cellHt, "D")
	doc.MCell(20, cellHt, "extra", "1", "L", false)

	CreatePDFInProjectRootOutFolder(doc.Fpdf, "TestMCell.pdf")
}
