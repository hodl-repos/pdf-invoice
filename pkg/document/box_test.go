package document

import "testing"

func TestBoxDebug(t *testing.T) {
	doc := NewA4()
	doc.SetDebug(true)

	x, y := doc.GetXY()
	b := doc.NewBox(x, y, 40, 30)

	b.SetFocus()

	CreatePDFInProjectRootOutFolder(doc.Fpdf, "BoxDebug.pdf")
}
