package pdfhelper

import (
	"fmt"
	"testing"
)

func TestDocTableHeader(t *testing.T) {
	doc := NewDocA4()

	header := []string{"", "Amount in EUR excl. VAT", "VAT", "VAT Amount", "Amount in EUR incl. VAT"}
	cols := []string{"d", "f", "f", "f", "f"}
	colGaps := []float64{0, 0, 5, 5, 5}
	colAligns := []string{"L", "R", "R", "R", "R"}

	th := NewDocTableHeader(doc, header, cols, colGaps, colAligns)
	th.SetBorderStr("1")
	th.Generate()
	doc.SetFont("Arial", "B", 10)
	th.SetHeight(10)
	th.Generate()

	CreatePDFInProjectRootOutFolder(doc.Fpdf, "TestDocTableHeader.pdf")
}

func TestDocTable(t *testing.T) {
	doc := NewDocA4()

	header := []string{"", "Amount in EUR excl. VAT", "VAT", "VAT Amount", "Amount in EUR incl. VAT"}
	hCols := []string{"d", "f", "f", "f", "f"}
	colGaps := []float64{0, 0, 5, 5, 5}
	colAligns := []string{"L", "R", "R", "R", "R"}

	th := NewDocTableHeader(doc, header, hCols, colGaps, colAligns)
	doc.SetFont("Arial", "B", 10)
	th.SetHeight(10)
	th.Generate()
	DrawVerticalLinePrintWidth(doc.Fpdf)

	rows := [][]string{
		{"1 Sorglospaket (Anzug 3-teilig, 2 Maßhemden, Krawatte oder Fliege und Stecktuch)", "50.0", "0.15", "7.5", "57.5"},
		{"Item 2", "50.0", "0.15", "7.5", "57.5"},
	}

	cols := th.GetColWidths()
	fmt.Println(cols)
	colAligns = []string{"L", "TR", "TR", "TR", "TR"}

	table := NewDocTable(doc, rows, cols, colAligns)

	// table.SetBorderStr("1")
	doc.SetFont("Arial", "", 8)
	table.Generate()

	CreatePDFInProjectRootOutFolder(doc.Fpdf, "TestDocTable.pdf")
}
