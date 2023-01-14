package pdfhelper

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDocTableV2CellsError(t *testing.T) {
	doc := NewDocA4()

	// first row is smaller
	cells := [][]string{
		{"A", "B", "C"},
		{"E", "F", "G", "H"},
		{"I", "J", "K", "L"},
		{"M", "N", "O", "P"},
		{"Q", "R", "S", "T"},
		{"X", "Y", "Z", ""},
	}
	_, err := NewDocTableV2(doc, cells)
	assert.EqualError(t, err, "row 2 has mismatching columns: got: 4 should: 3")

	// third row is smaller
	cells = [][]string{
		{"A", "B", "C", "D"},
		{"E", "F", "G", "H"},
		{"I", "J", "K"},
		{"M", "N", "O", "P"},
		{"Q", "R", "S", "T"},
		{"X", "Y", "Z", ""},
	}
	_, err = NewDocTableV2(doc, cells)
	assert.EqualError(t, err, "row 3 has mismatching columns: got: 3 should: 4")
}

func TestDocTableV2DynamicCols(t *testing.T) {
	doc := NewDocA4()

	cells := [][]string{
		{"A", "B", "C", "D"},
		{"E", "F", "G", "H"},
		{"I", "J", "K", "L"},
		{"M", "N", "O", "P"},
		{"Q", "R", "S", "T"},
		{"X", "Y", "Z", ""},
	}

	table, err := NewDocTableV2(doc, cells)
	assert.NoError(t, err)
	err = table.Generate()
	assert.NoError(t, err)

	CreatePDFInProjectRootOutFolder(doc.Fpdf, "TestDocTableV2DynamicCols.pdf")
}

func TestDocTableV2FixedCols(t *testing.T) {
	doc := NewDocA4()

	cells := [][]string{
		{"A", "B", "C", "D"},
		{"E", "F", "G", "H"},
		{"I", "J", "K", "L"},
		{"M", "N", "O", "P"},
		{"Q", "R", "S", "T"},
		{"X", "Y", "Z", ""},
	}

	table, err := NewDocTableV2(doc, cells)
	assert.NoError(t, err)
	table.SetAllColTypes(ColFixed)
	table.SetAllColFixedWidths(10)
	err = table.Generate()
	assert.NoError(t, err)

	CreatePDFInProjectRootOutFolder(doc.Fpdf, "TestDocTableV2FixedCols.pdf")
}

func TestDocTableV2CalculatedCols(t *testing.T) {
	doc := NewDocA4()

	cells := [][]string{
		{"A", "B", "C", "D"},
		{"E", "F", "G", "H"},
		{"I", "J", "K", "L"},
		{"M", "N", "O", "P"},
		{"Q", "R", "S", "T"},
		{"X", "Y", "Z", ""},
	}

	table, err := NewDocTableV2(doc, cells)
	assert.NoError(t, err)
	table.SetAllColTypes(ColCalc)
	err = table.Generate()
	assert.NoError(t, err)

	CreatePDFInProjectRootOutFolder(doc.Fpdf, "TestDocTableV2CalculatedCols.pdf")
}

func TestDocTableV2CalculatedAndDynamicCols(t *testing.T) {
	doc := NewDocA4()

	cells := [][]string{
		{"First column", "Second column", "Thirld column", "Forth column"},
		{"A", "B", "C", "D"},
		{"E", "F", "G", "H"},
		{"I", "J", "K", "L"},
		{"M", "N", "O", "P"},
		{"Q", "R", "S", "T"},
		{"X", "Y", "Z", ""},
	}

	colTypes := []ColumnType{ColDyn, ColCalc, ColCalc, ColCalc}

	table, err := NewDocTableV2(doc, cells)
	assert.NoError(t, err)

	err = table.SetColTypes(colTypes)
	assert.NoError(t, err)

	table.SetAllCellPaddings(Padding{1, 1, 1, 1})

	err = table.Generate()
	assert.NoError(t, err)

	CreatePDFInProjectRootOutFolder(doc.Fpdf, "TestDocTableV2CalculatedAndDynamicCols.pdf")
}

func TestDocTableV2Ellipsis(t *testing.T) {
	doc := NewDocA4()

	cells := [][]string{
		{"Christian Faustmann", "Christian Faustmann", "Christian Faustmann", "Christian Faustmann", "Christian Faustmann"},
	}

	table, err := NewDocTableV2(doc, cells)
	assert.NoError(t, err)
	table.SetAllColTypes(ColFixed)
	err = table.SetColFixedWidths([]float64{10, 15, 20, 25, 30})
	assert.NoError(t, err)

	err = table.Generate()
	assert.NoError(t, err)

	CreatePDFInProjectRootOutFolder(doc.Fpdf, "TestDocTableV2Ellipsis.pdf")
}

func TestDocTableV2Padding(t *testing.T) {
	doc := NewDocA4()

	cells := [][]string{
		{"######", "######", "######"},
		{"######", "######", "######"},
		{"######", "######", "######"},
		{"######", "######", "######"},
		{"######", "######", "######"},
		{"######", "######", "######"},
	}
	paddings := [][]Padding{
		{Padding{1, 1, 1, 1}, Padding{1, 1, 1, 1}, Padding{1, 1, 1, 1}},
		{Padding{0, 0, 0, 0}, Padding{2, 0, 0, 0}, Padding{2, 0, 0, 0}},
		{Padding{0, 0, 0, 0}, Padding{0, 2, 0, 0}, Padding{0, 2, 0, 0}},
		{Padding{0, 0, 0, 0}, Padding{0, 0, 2, 0}, Padding{0, 0, 2, 0}},
		{Padding{0, 0, 0, 0}, Padding{0, 0, 0, 2}, Padding{0, 0, 0, 2}},
		{Padding{0, 0, 0, 0}, Padding{2, 0, 2, 0}, Padding{0, 2, 0, 2}},
	}

	table, err := NewDocTableV2(doc, cells)
	assert.NoError(t, err)
	table.SetAllColTypes(ColCalc)
	// table.SetAllCellPaddings(Padding{2, 0, 0, 0})
	table.SetAllCellLineHeightFactors(1.)
	err = table.SetCellPaddings(paddings)
	assert.NoError(t, err)

	err = table.Generate()
	assert.NoError(t, err)

	CreatePDFInProjectRootOutFolder(doc.Fpdf, "TestDocTableV2Padding.pdf")
}

func TestDocTableV2CellAlignsPerColumn(t *testing.T) {
	doc := NewDocA4()

	cells := [][]string{
		{"First column", "Second column", "Thirld column", "Forth column"},
		{"A", "B", "C", "D"},
		{"E", "F", "G", "H"},
		{"I", "J", "K", "L"},
		{"M", "N", "O", "P"},
		{"Q", "R", "S", "T"},
		{"X", "Y", "Z", ""},
	}

	colTypes := []ColumnType{ColDyn, ColCalc, ColCalc, ColCalc}
	colAligns := []CellAlignment{AlignLeft, AlignRight, AlignRight, AlignRight}

	table, err := NewDocTableV2(doc, cells)
	assert.NoError(t, err)

	err = table.SetColTypes(colTypes)
	assert.NoError(t, err)

	table.SetAllCellPaddings(Padding{1, 1, 1, 1})

	err = table.SetCellAlingsPerColumn(colAligns[1:])
	assert.EqualError(t, err, "column count mismatch: got: 3 should: 4")
	err = table.SetCellAlingsPerColumn(colAligns)
	assert.NoError(t, err)

	err = table.Generate()
	assert.NoError(t, err)

	CreatePDFInProjectRootOutFolder(doc.Fpdf, "TestDocTableV2CellAlignsPerColumn.pdf")
}

func TestDocTableV2SetColGaps(t *testing.T) {
	doc := NewDocA4()

	cells := [][]string{
		{"A", "B", "C", "D"},
		{"E", "F", "G", "H"},
		{"I", "J", "K", "L"},
		{"M", "N", "O", "P"},
		{"Q", "R", "S", "T"},
		{"X", "Y", "Z", ""},
	}

	colGaps := []float64{5, 10, 15}

	table, err := NewDocTableV2(doc, cells)
	assert.NoError(t, err)

	err = table.SetColGaps(colGaps[1:])
	assert.EqualError(t, err, "column count mismatch: got: 2 should: 3")
	err = table.SetColGaps(colGaps)
	assert.NoError(t, err)

	err = table.Generate()
	assert.NoError(t, err)

	CreatePDFInProjectRootOutFolder(doc.Fpdf, "TestDocTableV2SetColGaps.pdf")
}

func TestDocTableV2SetAllRowGaps(t *testing.T) {
	doc := NewDocA4()

	cells := [][]string{
		{"A", "B", "C", "D"},
		{"E", "F", "G", "H"},
		{"I", "J", "K", "L"},
		{"M", "N", "O", "P"},
		{"Q", "R", "S", "T"},
		{"X", "Y", "Z", ""},
	}

	rowGap := 2.5

	table, err := NewDocTableV2(doc, cells)
	assert.NoError(t, err)

	table.SetAllRowGaps(rowGap)

	err = table.Generate()
	assert.NoError(t, err)

	CreatePDFInProjectRootOutFolder(doc.Fpdf, "TestDocTableV2SetAllRowGaps.pdf")
}

func TestDocTableV2SetRowGaps(t *testing.T) {
	doc := NewDocA4()

	cells := [][]string{
		{"A", "B", "C", "D"},
		{"E", "F", "G", "H"},
		{"I", "J", "K", "L"},
		{"M", "N", "O", "P"},
		{"Q", "R", "S", "T"},
		{"X", "Y", "Z", ""},
	}

	rowGaps := []float64{1, 2, 3, 4, 5}

	table, err := NewDocTableV2(doc, cells)
	assert.NoError(t, err)

	err = table.SetRowGaps(rowGaps[1:])
	assert.EqualError(t, err, "row count mismatch: got: 4 should: 5")
	err = table.SetRowGaps(rowGaps)
	assert.NoError(t, err)

	err = table.Generate()
	assert.NoError(t, err)

	CreatePDFInProjectRootOutFolder(doc.Fpdf, "TestDocTableV2SetRowGaps.pdf")
}

func TestDocTableV2RowAndColGap(t *testing.T) {
	doc := NewDocA4()

	cells := [][]string{
		{"A", "B", "C", "D"},
		{"E", "F", "G", "H"},
		{"I", "J", "K", "L"},
		{"M", "N", "O", "P"},
		{"Q", "R", "S", "T"},
		{"X", "Y", "Z", ""},
	}

	rowGaps := []float64{1, 2, 3, 4, 5}
	colGaps := []float64{1, 2, 3}

	table, err := NewDocTableV2(doc, cells)
	assert.NoError(t, err)

	err = table.SetRowGaps(rowGaps)
	assert.NoError(t, err)

	err = table.SetColGaps(colGaps)
	assert.NoError(t, err)

	err = table.Generate()
	assert.NoError(t, err)

	CreatePDFInProjectRootOutFolder(doc.Fpdf, "TestDocTableV2ColAndRowGap.pdf")
}

func TestAlignToFpdf(t *testing.T) {
	doc := NewDocA4()

	w := 30.
	h := 20.

	doc.CFormat(w, h, "center", "1", 1, alignToFpdf(AlignCenter), false, 0, "")
	doc.CFormat(w, h, "top", "1", 1, alignToFpdf(AlignTop), false, 0, "")
	doc.CFormat(w, h, "right", "1", 1, alignToFpdf(AlignRight), false, 0, "")
	doc.CFormat(w, h, "bottom", "1", 1, alignToFpdf(AlignBottom), false, 0, "")
	doc.CFormat(w, h, "left", "1", 1, alignToFpdf(AlignLeft), false, 0, "")
	doc.CFormat(w, h, "topLeft", "1", 1, alignToFpdf(AlignTopLeft), false, 0, "")
	doc.CFormat(w, h, "topRight", "1", 1, alignToFpdf(AlignTopRight), false, 0, "")
	doc.CFormat(w, h, "bottomRight", "1", 1, alignToFpdf(AlignBottomRight), false, 0, "")
	doc.CFormat(w, h, "bottomLeft", "1", 1, alignToFpdf(AlignBottomLeft), false, 0, "")

	CreatePDFInProjectRootOutFolder(doc.Fpdf, "TestAlignToFpdf.pdf")
}
