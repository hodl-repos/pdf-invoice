package pdfhelper

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDocTableCellsError(t *testing.T) {
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
	_, err := NewDocTable(doc, cells)
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
	_, err = NewDocTable(doc, cells)
	assert.EqualError(t, err, "row 3 has mismatching columns: got: 3 should: 4")
}

func TestDocTableDynamicCols(t *testing.T) {
	doc := NewDocA4()

	cells := [][]string{
		{"A", "B", "C", "D"},
		{"E", "F", "G", "H"},
		{"I", "J", "K", "L"},
		{"M", "N", "O", "P"},
		{"Q", "R", "S", "T"},
		{"X", "Y", "Z", ""},
	}

	table, err := NewDocTable(doc, cells)
	assert.NoError(t, err)
	err = table.Generate()
	assert.NoError(t, err)

	CreatePDFInProjectRootOutFolder(doc.Fpdf, "DocTableDynamicCols.pdf")
}

func TestDocTableFixedCols(t *testing.T) {
	doc := NewDocA4()

	cells := [][]string{
		{"A", "B", "C", "D"},
		{"E", "F", "G", "H"},
		{"I", "J", "K", "L"},
		{"M", "N", "O", "P"},
		{"Q", "R", "S", "T"},
		{"X", "Y", "Z", ""},
	}

	table, err := NewDocTable(doc, cells)
	assert.NoError(t, err)
	table.SetAllColTypes(ColFixed)
	table.SetAllColFixedWidths(10)
	err = table.Generate()
	assert.NoError(t, err)

	CreatePDFInProjectRootOutFolder(doc.Fpdf, "DocTableFixedCols.pdf")
}

func TestDocTableCalculatedCols(t *testing.T) {
	doc := NewDocA4()

	cells := [][]string{
		{"A", "B", "C", "D"},
		{"E", "F", "G", "H"},
		{"I", "J", "K", "L"},
		{"M", "N", "O", "P"},
		{"Q", "R", "S", "T"},
		{"X", "Y", "Z", ""},
	}

	table, err := NewDocTable(doc, cells)
	assert.NoError(t, err)
	table.SetAllColTypes(ColCalc)
	err = table.Generate()
	assert.NoError(t, err)

	CreatePDFInProjectRootOutFolder(doc.Fpdf, "DocTableCalculatedCols.pdf")
}

func TestDocTableCalculatedAndDynamicCols(t *testing.T) {
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

	table, err := NewDocTable(doc, cells)
	assert.NoError(t, err)

	err = table.SetColTypes(colTypes)
	assert.NoError(t, err)

	table.SetAllCellPaddings(Padding{1, 1, 1, 1})

	err = table.Generate()
	assert.NoError(t, err)

	CreatePDFInProjectRootOutFolder(doc.Fpdf, "DocTableCalculatedAndDynamicCols.pdf")
}

func TestDocTableEllipsis(t *testing.T) {
	doc := NewDocA4()

	cells := [][]string{
		{"Christian Faustmann", "Christian Faustmann", "Christian Faustmann", "Christian Faustmann", "Christian Faustmann"},
	}

	table, err := NewDocTable(doc, cells)
	assert.NoError(t, err)
	table.SetAllColTypes(ColFixed)
	err = table.SetColFixedWidths([]float64{10, 15, 20, 25, 30})
	assert.NoError(t, err)

	err = table.Generate()
	assert.NoError(t, err)

	CreatePDFInProjectRootOutFolder(doc.Fpdf, "DocTableEllipsis.pdf")
}

func TestDocTablePadding(t *testing.T) {
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

	table, err := NewDocTable(doc, cells)
	assert.NoError(t, err)
	table.SetAllColTypes(ColCalc)
	// table.SetAllCellPaddings(Padding{2, 0, 0, 0})
	table.SetAllCellLineHeightFactors(1.)
	err = table.SetCellPaddings(paddings)
	assert.NoError(t, err)

	err = table.Generate()
	assert.NoError(t, err)

	CreatePDFInProjectRootOutFolder(doc.Fpdf, "DocTablePadding.pdf")
}

func TestDocTableCelPaddingsPerColumn(t *testing.T) {
	doc := NewDocA4()

	cells := [][]string{
		{
			"Polaroid gastropub stumptown microdosing vegan fanny pack. Ugh prism quinoa keytar organic hexagon before they sold out poutine taiyaki whatever four dollar toast photo booth small batch.",
			"Vice humblebrag edison bulb cloud bread heirloom yes plz direct trade ennui lo-fi cronut fingerstache knausgaard pickled pabst small batch.",
			"Brooklyn ascot listicle pitchfork edison bulb pok pok disrupt single-origin coffee wayfarers banh mi pabst plaid.",
			"Cronut kogi pour-over retro affogato, scenester occupy godard. Schlitz taxidermy umami bushwick occupy kitsch. Irony retro wolf hot chicken +1 thundercats microdosing pour-over truffaut butcher air plant organic crucifix.",
		},
		{
			"Brooklyn ascot listicle pitchfork edison bulb pok pok disrupt single-origin coffee wayfarers banh mi pabst plaid.",
			"Polaroid gastropub stumptown microdosing vegan fanny pack. Ugh prism quinoa keytar organic hexagon before they sold out poutine taiyaki whatever four dollar toast photo booth small batch.",
			"Vice humblebrag edison bulb cloud bread heirloom yes plz direct trade ennui lo-fi cronut fingerstache knausgaard pickled pabst small batch.",
			"Cronut kogi pour-over retro affogato, scenester occupy godard. Schlitz taxidermy umami bushwick occupy kitsch. Irony retro wolf hot chicken +1 thundercats microdosing pour-over truffaut butcher air plant organic crucifix.",
		},
		{
			"Cronut kogi pour-over retro affogato, scenester occupy godard. Schlitz taxidermy umami bushwick occupy kitsch. Irony retro wolf hot chicken +1 thundercats microdosing pour-over truffaut butcher air plant organic crucifix.",
			"Polaroid gastropub stumptown microdosing vegan fanny pack. Ugh prism quinoa keytar organic hexagon before they sold out poutine taiyaki whatever four dollar toast photo booth small batch.",
			"Brooklyn ascot listicle pitchfork edison bulb pok pok disrupt single-origin coffee wayfarers banh mi pabst plaid.",
			"Vice humblebrag edison bulb cloud bread heirloom yes plz direct trade ennui lo-fi cronut fingerstache knausgaard pickled pabst small batch.",
		},
	}

	paddings := []Padding{{1, 1, 1, 1}, {1, 1, 1, 1}, {1, 1, 1, 5}, {1, 1, 1, 1}}

	table, err := NewDocTable(doc, cells)
	assert.NoError(t, err)

	table.SetAllCellTypes(CellMulti)

	err = table.SetCellPaddingsPerColumn(paddings[1:])
	assert.EqualError(t, err, "column count mismatch: got: 3 should: 4")
	err = table.SetCellPaddingsPerColumn(paddings)
	assert.NoError(t, err)

	err = table.Generate()
	assert.NoError(t, err)

	CreatePDFInProjectRootOutFolder(doc.Fpdf, "DocTableCelPaddingsPerColumn.pdf")
}

func TestDocTableCellAlignsPerColumn(t *testing.T) {
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

	table, err := NewDocTable(doc, cells)
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

	CreatePDFInProjectRootOutFolder(doc.Fpdf, "DocTableCellAlignsPerColumn.pdf")
}

func TestDocTableSetColGaps(t *testing.T) {
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

	table, err := NewDocTable(doc, cells)
	assert.NoError(t, err)

	err = table.SetColGaps(colGaps[1:])
	assert.EqualError(t, err, "column count mismatch: got: 2 should: 3")
	err = table.SetColGaps(colGaps)
	assert.NoError(t, err)

	err = table.Generate()
	assert.NoError(t, err)

	CreatePDFInProjectRootOutFolder(doc.Fpdf, "DocTableSetColGaps.pdf")
}

func TestDocTableSetAllRowGaps(t *testing.T) {
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

	table, err := NewDocTable(doc, cells)
	assert.NoError(t, err)

	table.SetAllRowGaps(rowGap)

	err = table.Generate()
	assert.NoError(t, err)

	CreatePDFInProjectRootOutFolder(doc.Fpdf, "DocTableSetAllRowGaps.pdf")
}

func TestDocTableSetRowGaps(t *testing.T) {
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

	table, err := NewDocTable(doc, cells)
	assert.NoError(t, err)

	err = table.SetRowGaps(rowGaps[1:])
	assert.EqualError(t, err, "row count mismatch: got: 4 should: 5")
	err = table.SetRowGaps(rowGaps)
	assert.NoError(t, err)

	err = table.Generate()
	assert.NoError(t, err)

	CreatePDFInProjectRootOutFolder(doc.Fpdf, "DocTableSetRowGaps.pdf")
}

func TestDocTableRowAndColGap(t *testing.T) {
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

	table, err := NewDocTable(doc, cells)
	assert.NoError(t, err)

	err = table.SetRowGaps(rowGaps)
	assert.NoError(t, err)

	err = table.SetColGaps(colGaps)
	assert.NoError(t, err)

	err = table.Generate()
	assert.NoError(t, err)

	CreatePDFInProjectRootOutFolder(doc.Fpdf, "DocTableColAndRowGap.pdf")
}

func TestDocTableCellMulti(t *testing.T) {
	doc := NewDocA4()

	cells := [][]string{{
		"Polaroid gastropub stumptown microdosing vegan fanny pack. Ugh prism quinoa keytar organic hexagon before they sold out poutine taiyaki whatever four dollar toast photo booth small batch.",
		"Vice humblebrag edison bulb cloud bread heirloom yes plz direct trade ennui lo-fi cronut fingerstache knausgaard pickled pabst small batch.",
		"Brooklyn ascot listicle pitchfork edison bulb pok pok disrupt single-origin coffee wayfarers banh mi pabst plaid.",
		"Cronut kogi pour-over retro affogato, scenester occupy godard. Schlitz taxidermy umami bushwick occupy kitsch. Irony retro wolf hot chicken +1 thundercats microdosing pour-over truffaut butcher air plant organic crucifix.",
	},
	}

	table, err := NewDocTable(doc, cells)
	assert.NoError(t, err)

	table.SetCellType(0, 1, CellMulti)
	table.SetCellType(0, 3, CellMulti)

	err = table.Generate()
	assert.NoError(t, err)

	CreatePDFInProjectRootOutFolder(doc.Fpdf, "DocTableCellMulti.pdf")
}

func TestDocTableValidateColumns(t *testing.T) {
	doc := NewDocA4()

	cells := [][]string{
		{"A", "B", "C", "D"},
		{"E", "F", "G", "H"},
		{"I", "J", "K", "L"},
		{"M", "N", "O", "P"},
		{"Q", "R", "S", "T"},
		{"X", "Y", "Z", ""},
	}

	colTypes := []ColumnType{ColCalc, ColDyn, ColDyn, ColDyn}
	cellTypes := []CellType{CellMulti, CellSingle, CellSingle, CellSingle}

	table, err := NewDocTable(doc, cells)
	assert.NoError(t, err)

	err = table.SetColTypes(colTypes)
	assert.NoError(t, err)

	err = table.SetCellTypesPerColumn(cellTypes)
	assert.NoError(t, err)

	err = table.Generate()
	assert.EqualError(t, err, "column 1 of type ColCalc has only CellMulti cells and cannot be calculated")
}

func TestDocTableValidateRows(t *testing.T) {
	doc := NewDocA4()

	cells := matrix(8, 8, "")

	table, err := NewDocTable(doc, cells)
	assert.NoError(t, err)

	err = table.SetCell(2, 2, "Banjo tumeric letterpress plaid echo park la croix iceland gastropub dreamcatcher.")
	assert.NoError(t, err)
	err = table.SetCellType(2, 2, CellMulti)
	assert.NoError(t, err)

	table.SetAllColTypes(ColFixed)
	table.SetAllColFixedWidths(15)
	table.SetAllRowTypes(RowFixed)
	table.SetAllRowFixedHeights(15)

	err = table.Generate()
	assert.EqualError(t, err, "row 3 cannot display all cells; insufficient height")
}

func TestDocTableRowFixed(t *testing.T) {
	doc := NewDocA4()

	cells := matrix(8, 8, "")

	table, err := NewDocTable(doc, cells)
	assert.NoError(t, err)

	table.SetAllColTypes(ColFixed)
	table.SetAllColFixedWidths(15)
	table.SetAllRowTypes(RowFixed)
	table.SetAllRowFixedHeights(15)

	err = table.Generate()
	assert.NoError(t, err)

	CreatePDFInProjectRootOutFolder(doc.Fpdf, "DocTableRowFixed.pdf")
}

func TestDocTableCheckIndices(t *testing.T) {
	doc := NewDocA4()

	cells := matrix(8, 7, "")

	table, err := NewDocTable(doc, cells)
	assert.NoError(t, err)

	err = table.SetCell(-1, 0, "")
	assert.EqualError(t, err, "invalid row index: got: -1 should: 0-7")
	err = table.SetCell(8, 0, "")
	assert.EqualError(t, err, "invalid row index: got: 8 should: 0-7")

	err = table.SetCell(0, -1, "")
	assert.EqualError(t, err, "invalid column index: got: -1 should: 0-6")
	err = table.SetCell(0, 7, "")
	assert.EqualError(t, err, "invalid column index: got: 7 should: 0-6")

	err = table.SetCell(0, 0, "")
	assert.NoError(t, err)
	err = table.SetCell(7, 6, "")
	assert.NoError(t, err)
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

	CreatePDFInProjectRootOutFolder(doc.Fpdf, "AlignToFpdf.pdf")
}
