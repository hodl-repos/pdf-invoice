package document

import "testing"

func TestLayout(t *testing.T) {
	doc := NewA4()
	doc.SetDebug(true)

	b := doc.NewBox(25, 98.46, 210-25-20, 297-98.46-25-12.5, BoxOpen)

	l := NewLayout(doc, b)

	l.NewPage()

	CreatePDFInProjectRootOutFolder(doc.Fpdf, "Layout.pdf")
}
