package pdfhelper

import "fmt"

// CollumnType determines how a column width will be calculated.
type ColumnType int

const (
	ColCalc ColumnType = iota
	ColFixed
	ColDyn
)

// CellAlignment determines how the content is aligned inside a cell.
//
// cell:
// #----------------------------------------------#
// | AlignTopLeft      AlignTop     AlignTopRight |
// |                                              |
// | AlignLeft	      AlignCenter	     AlignRight |
// |                                              |
// | AlignBottomLeft AlignBottom AlignBottomRight |
// #----------------------------------------------#
type CellAlignment int

const (
	AlignCenter CellAlignment = iota
	AlignTop
	AlignRight
	AlignBottom
	AlignLeft
	AlignTopLeft
	AlignTopRight
	AlignBottomRight
	AlignBottomLeft
)

// Padding is the padding inside of a cell. The indices follow the convention
// Top(0), Right(1), Bottom(2), Left(3).
//
// #-- border -------------------------#
// |            paddingTop             |
// |             #------#              |
// | paddingLeft | cell | paddingRight |
// |             #------#              |
// |          paddingBottom            |
// #-----------------------------------#
type Padding [4]float64

const (
	paddingTop = iota
	paddingRight
	paddingBottom
	paddingLeft
)

// CellType determines how a cell will be rendered.
type CellType int

const (
	CellSingle CellType = iota
	CellMulti
)

type DocTableV2 struct {
	doc *Doc

	// colTypes determine how a column width will be calculated.
	//
	// A column can be of type ColCalc, ColFixed, or ColDyn.
	//
	// ColCalc - Calculate column width by going through all cells of the column
	// with type CellSingle and takes the widest cell as measurement. The
	// individual cell width is calculated by Width = cell content (single line) +
	// paddingLeft + paddingRight. colFixedWidth will be ignored.
	//
	// ColFixed - colWidth is set by colFixedWidth.
	//
	// ColDyn (default) - Column with dymamic width. The width is determined after
	// all calculated & fixed column widths, and column gaps are subtracted from
	// the printWidth. The remaining width will be devided by the count of dynamic
	// columns.
	colTypes []ColumnType
	// colWidths determine the column with for columns with column type ColFixed.
	// Otherwise the value will be ignored.
	colFixedWidths []float64
	// colGaps are spaces between columns. Therefore there are only len(cols) - 1
	// gaps per table. Every other gap will be ignored.
	colGaps []float64
	// colWidths will be calculated regarding all parameters and represent the
	// current column widths.
	colWidths []float64

	// rowGap is the space between two rows. Default is 0.
	rowGap float64
	// rowHeight defines the height of the
	rowHeight float64

	// cells hold the content of the table row by row. Each row holds the content
	// of each cell per column.
	//
	// The dimensions of cells are [rowCount][colCount]string.
	cells [][]string
	// cellTypes determine how a cell will be rendered.
	//
	// CellSingle will render the cell content on a single line.
	//
	// CellMulti will render the cell content over multiple lines contraint by the
	// column width.
	//
	// Default is CellSingle.
	cellTypes [][]CellType
	// cellAlings determine the alignment of every table cell. Default is Left
	//
	// cell:
	// #----------------------------------------------#
	// | AlignTopLeft      AlignTop     AlignTopRight |
	// |                                              |
	// | AlignLeft	      AlignCenter	     AlignRight |
	// |                                              |
	// | AlignBottomLeft AlignBottom AlignBottomRight |
	// #----------------------------------------------#
	cellAligns [][]CellAlignment
	// cellPaddings determine the padding of every individual table cell.
	//
	// Padding is an additional space inside the cell around the content.
	// #-- border -------------------------#
	// |            paddingTop             |
	// |             #------#              |
	// | paddingLeft | cell | paddingRight |
	// |             #------#              |
	// |          paddingBottom            |
	// #-----------------------------------#
	cellPaddings [][]Padding
	// cellLineHeightFactors determine the line height of every cell calculated by
	// lineHeight = fontHeight * lineHeightFactor, where fontHeight is the height
	// of the current font in unit size (see Fpdf.GetFontSize()).
	cellLineHeightFactors [][]float64
	//cellBorders determine for every cell if a border should be displayed.
	cellBorders [][]bool

	// tableRows will be calculated by the number of rows of cells.
	tableRows int
	// tableCols will be calculated by the number of cells in a row.
	tableCols int
	// tableWidth will be calculated by calcColWidths()
	tableWidth float64
}

func NewDocTableV2(doc *Doc, cells [][]string) (*DocTableV2, error) {
	// checking if every row has the same length.
	rowLen := len(cells[0])
	for i, r := range cells {
		if len(r) != rowLen {
			return nil, fmt.Errorf("row %v has mismatching columns: got: %v should: %v", i, len(r), rowLen)
		}
	}

	t := &DocTableV2{
		doc:   doc,
		cells: cells,
	}
	t.tableRows = len(cells)
	t.tableCols = len(cells[0])
	t.SetDefaults()
	return t, nil
}

func (t *DocTableV2) SetDefaults() {
	rows := t.tableRows
	cols := t.tableCols
	if len(t.colTypes) == 0 {
		t.colTypes = array(cols, ColDyn)
	}
	if len(t.colFixedWidths) == 0 {
		t.colFixedWidths = array(cols, 0.)
	}
	if len(t.colGaps) == 0 {
		t.colGaps = array(cols, 0.)
	}

	t.rowGap = 0
	t.rowHeight = 0

	if len(t.cellBorders) == 0 {
		t.cellBorders = matrix(rows, cols, true)
	}
	if len(t.cellLineHeightFactors) == 0 {
		t.cellLineHeightFactors = matrix(rows, cols, 1.2)
	}
	if len(t.cellPaddings) == 0 {
		t.cellPaddings = matrix(rows, cols, Padding{0, 0, 0, 0})
	}
	if len(t.cellAligns) == 0 {
		t.cellAligns = matrix(rows, cols, AlignLeft)
	}
	if len(t.cellTypes) == 0 {
		t.cellTypes = matrix(rows, cols, CellSingle)
	}
}

func (t *DocTableV2) calcColWidths() {
	t.colWidths = array(t.tableCols, .0)
	printWidth := GetPrintWidth(t.doc.Fpdf)
	tableWidth := 0.

	dynCount := 0
	for i, cType := range t.colTypes {
		if cType == ColFixed {
			t.colWidths[i] = t.colFixedWidths[i]
			tableWidth += t.colFixedWidths[i]
			continue
		}

		if cType == ColCalc {
			maxCellWidth := t.maxCellWidthOfColumn(i)
			t.colWidths[i] = maxCellWidth
			tableWidth += maxCellWidth
		}

		if cType == ColDyn {
			dynCount++
		}
	}

	// remove all colGaps from the remaining width
	for _, g := range t.colGaps {
		tableWidth += g
	}

	dynWidth := 0.
	if remWidth := printWidth - tableWidth; dynCount > 0 && remWidth > 0 {
		dynWidth = remWidth / float64(dynCount)
	}

	for i, cType := range t.colTypes {
		if cType == ColDyn {
			t.colWidths[i] = dynWidth
			tableWidth += dynWidth
		}
	}

	// set table width
	t.tableWidth = tableWidth
}

// FIXME: There is no check for columns with type ColCalc which only have cells
// of type CellMulti.
func (t *DocTableV2) maxCellWidthOfColumn(j int) float64 {
	maxCellWidth := 0.
	for i := 0; i < t.tableRows; i++ {
		p := t.cellPaddings[i][j]
		// filter all CellMulti cells
		if t.cellTypes[i][j] == CellMulti {
			continue
		}
		w := t.doc.GetStringWidth(t.cells[i][j]) + p[paddingLeft] + p[paddingRight]
		if w > maxCellWidth {
			maxCellWidth = w
		}
	}
	return maxCellWidth
}

func (t *DocTableV2) Generate() error {
	t.calcColWidths()
	printWidth := GetPrintWidth(t.doc.Fpdf)
	if t.tableWidth > printWidth {
		return fmt.Errorf("error generating table: table wider than print width: %v > %v", t.tableWidth, printWidth)
	}
	for i := 0; i < t.tableRows; i++ {
		rowHt := t.getRowHeight(i)
		for j := 0; j < t.tableCols; j++ {
			t.renderCell(i, j, rowHt)
		}
	}

	return nil
}

func (t *DocTableV2) renderCell(i, j int, rowHt float64) {
	doc := t.doc
	x, y := doc.GetXY()
	p := t.cellPaddings[i][j]
	w := t.colWidths[j] - p[paddingLeft] - p[paddingRight]

	if t.cellBorders[i][j] {
		doc.Rect(x, y, t.colWidths[j], rowHt, "D")
	}

	alignStr := alignToFpdf(t.cellAligns[i][j])

	doc.SetXY(x+p[paddingLeft], y+p[paddingTop])
	switch t.cellTypes[i][j] {
	case CellSingle:
		cellStr := t.cells[i][j]
		// ln indicates where the current position should go after the call. Possible
		// values are 0 (to the right), 1 (to the beginning of the next line), and 2
		// (below). Putting 1 is equivalent to putting 0 and calling Ln() just after.
		ln := 0
		// border is ""... no border because the cell border is created separately
		strWidth := doc.GetStringWidth(cellStr)
		// check if strWidth is within cell width. This can happen on the ColFixed
		// column type.
		if strWidth > w {
			for k := 1; k < len(cellStr); k++ {
				if doc.GetStringWidth(cellStr[:len(cellStr)-k]+doc.Ellipsis()) < w {
					cellStr = cellStr[:len(cellStr)-k] + doc.Ellipsis()
					break
				}
			}
		}
		cellHt := rowHt - p[paddingTop] - p[paddingBottom]
		doc.CFormat(w, cellHt, cellStr, "", ln, alignStr, false, 0, "")
		doc.SetXY(doc.GetX()+p[paddingRight], y)
	case CellMulti:
		doc.MCell(w, t.getLineHeight(i, j), t.cells[i][j], "", alignStr, false)
		doc.SetXY(x+w, y)
	default:
		panic("unsupported CellType: " + fmt.Sprint(t.cellTypes[i][j]))
	}
	if j == t.tableCols-1 {
		doc.Ln(rowHt)
	}
}

func (t *DocTableV2) getRowHeight(i int) float64 {
	rowHt := 0.
	for j := 0; j < t.tableCols; j++ {
		h := t.getCellHeight(i, j)
		if h > rowHt {
			rowHt = h
		}
	}
	return rowHt
}

func (t *DocTableV2) getCellHeight(i, j int) float64 {
	lineHt := t.getLineHeight(i, j)
	p := t.cellPaddings[i][j]
	cellPadding := p[paddingTop] + p[paddingBottom]

	switch t.cellTypes[i][j] {
	case CellSingle:
		return lineHt + cellPadding
	case CellMulti:
		lines := t.doc.SplitText(t.cells[i][j], t.colWidths[i])
		return float64(len(lines))*lineHt + cellPadding
	default:
		panic("unsupported CellType: " + fmt.Sprint(t.cellTypes[i][j]))
	}
}

func (t *DocTableV2) getLineHeight(i, j int) float64 {
	_, fontHt := t.doc.GetFontSize()
	return fontHt * t.cellLineHeightFactors[i][j]
}

// SETTER ----------------------------------------------------------------------
func (t *DocTableV2) SetAllColTypes(ct ColumnType) {
	t.colTypes = array(t.tableCols, ct)
}

func (t *DocTableV2) SetColTypes(colTypes []ColumnType) error {
	if len(colTypes) != t.tableCols {
		return fmt.Errorf("column count mismatch: parameterCols: %v tableCols: %v", len(colTypes), t.tableCols)
	}
	t.colTypes = colTypes
	return nil
}

func (t *DocTableV2) SetAllColFixedWidths(w float64) {
	t.colFixedWidths = array(t.tableCols, w)
}

func (t *DocTableV2) SetColFixedWidths(cFixedWidth []float64) error {
	if len(cFixedWidth) != t.tableCols {
		return fmt.Errorf("column count mismatch: parameterCols: %v tableCols: %v", len(cFixedWidth), t.tableCols)
	}
	t.colFixedWidths = cFixedWidth
	return nil
}

func (t *DocTableV2) SetAllCellBorders(b bool) {
	t.cellBorders = matrix(t.tableRows, t.tableCols, b)
}

func (t *DocTableV2) SetAllCellPaddings(p Padding) {
	t.cellPaddings = matrix(t.tableRows, t.tableCols, p)
}

func (t *DocTableV2) SetCellPaddings(p [][]Padding) error {
	if len(p) != t.tableRows || len(p[0]) != t.tableCols {
		return fmt.Errorf("row or column count mismatch: got: (rows:%v cols:%v) table: (rows:%v cols%v)", len(p), len(p[0]), t.tableRows, t.tableCols)
	}
	t.cellPaddings = p
	return nil
}

func (t *DocTableV2) SetAllCellAligns(a CellAlignment) {
	t.cellAligns = matrix(t.tableRows, t.tableCols, a)
}

func (t *DocTableV2) SetAllCellTypes(ct CellType) {
	t.cellTypes = matrix(t.tableRows, t.tableCols, ct)
}

func (t *DocTableV2) SetAllCellLineHeightFactors(f float64) {
	t.cellLineHeightFactors = matrix(t.tableRows, t.tableCols, f)
}

func matrix[T any](rows, cols int, val T) [][]T {
	m := make([][]T, rows)

	for i := 0; i < rows; i++ {
		m[i] = array(cols, val)
	}

	return m
}

func array[T any](len int, val T) []T {
	arr := make([]T, len)
	for i := 0; i < len; i++ {
		arr[i] = val
	}
	return arr
}

// alignToFpdf converts CellAlignments to the Fpdf equivalent(see below).
//
// alignStr specifies how the text is to be positioned within the cell.
// Horizontal alignment is controlled by including "L", "C" or "R" (left,
// center, right) in alignStr. Vertical alignment is controlled by including
// "T", "M", "B" or "A" (top, middle, bottom, baseline) in alignStr. The default
// alignment is left middle.
func alignToFpdf(a CellAlignment) string {
	switch a {
	case AlignCenter:
		return "MC"
	case AlignTop:
		return "TC"
	case AlignRight:
		return "MR"
	case AlignBottom:
		return "BC"
	case AlignLeft:
		return "ML"
	case AlignTopLeft:
		return "TL"
	case AlignTopRight:
		return "TR"
	case AlignBottomRight:
		return "BR"
	case AlignBottomLeft:
		return "BL"
	default:
		return ""
	}
}
