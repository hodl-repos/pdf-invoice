package document

import (
	"fmt"

	"github.com/jung-kurt/gofpdf"
)

// HeadType
type HeadType int

const (
	HeadUnset HeadType = iota
	HeadNone
	HeadFirstRow
)

// CollumnType determines how a column width will be calculated.
type ColumnType int

const (
	ColCalc ColumnType = iota
	ColFixed
	ColDyn
)

// RowType determines how a row height will be calculated.
type RowType int

const (
	RowCalc RowType = iota
	RowFixed
)

// CellType determines how a cell will be rendered.
type CellType int

const (
	CellSingle CellType = iota
	CellMulti
)

// DocTable
//
// NOTE: All used indices follow the convention that indices i are for rows and
// indices j are for columns.
//
// Render process: 1-3 are user actions
//
// 1. NewDocTable(doc *Doc, cells [][]string)
//
//   - setting row and column count
//
//   - setting default values
//
// 2. (optional) change table parameters
//
// 3. table.Generate()
//
//   - validate columns
//
//   - calculate columns
//
//   - validate rows
//
//   - calculate rows
//
//   - render cells
//
// Table parameters & defaults:
//   - colTypes: ColDyn
//   - colFixedWidths: 0.0
//   - colGaps: 0.0
//   - colWidths: calculated
//   - rowTypes: RowCalc
//   - rowFixedHeights: 0.0
//   - rowGaps: 0.0
//   - rowHeights: calculated
//   - cells: constructor parameter
//   - cellTypes: CellSingle
//   - cellLineHeightFactors: 1.2
//   - cellAligns: AlignLeft
//   - cellPaddings: Padding{0,0,0,0}
//   - cellBorder: true
type DocTable struct {
	doc *Doc

	// headType determines if the first row is used as head or not.
	//
	// HeadNone - All rows contain data.
	//
	// HeadFirstRow - First row is interpreted as head. When table is rendered
	// over a page break, the head gets rerendered before continuing with
	// rendering data rows.
	headType HeadType

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
	// colFixedWidths determine the column width for columns with column type
	// ColFixed. Otherwise the values will be ignored.
	colFixedWidths []float64
	// colGaps are spaces between columns. Therefore there are only len(cols) - 1
	// gaps per table. Default is 0.
	colGaps []float64
	// colWidths will be calculated regarding all parameters and represent the
	// current column widths.
	colWidths []float64

	// rowTypes determine how a row height will be calculated.
	//
	// A row can be of type RowCalc, or RowFixed.
	//
	// RowCalc (default) - Calculate row height by going through all cells of the
	// row and takes the heightest cell as measurement. The individual cell height
	// is calculated when cellType is CellSingle by Height = fontSize(in unit) *
	// cellLineHeightFactor + paddingTop + paddingBottom and when cellType is
	// CellMulti by Height = lines * fontSize(in unit) * cellLineHeightFactor +
	// paddingTop + paddingBottom. rowFixedHeight will be ignored.
	//
	// RowFixed - rowWidth is set by rowFixedHeight
	rowTypes []RowType
	// rowFixedHeights determine the row height for rows with row type RowFixed.
	// Otherwise the values will be ignored.
	rowFixedHeights []float64
	// rowGaps are spaces between rows. Therefore there are only len(rows) - 1
	// gaps per table. Default is 0.
	rowGaps []float64
	// rowHeights will be calculated regarding all parameters and represent the
	// current row heights.
	rowHeights []float64

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
	// cellLineHeightFactors determine the line height of every cell calculated by
	// lineHeight = fontHeight * lineHeightFactor, where fontHeight is the height
	// of the current font in unit size (see Fpdf.GetFontSize()).
	cellLineHeightFactors [][]float64
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
	cellAligns [][]AlignmentType
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
	// cellBorders determine for every cell if a border should be displayed.
	cellBorders [][]bool
	// cellStyleFuncs determine the style for every cell.
	cellStyleFuncs [][]*func(gofpdf.Fpdf)

	// tableRows will be calculated by the number of rows of cells.
	tableRows int
	// tableCols will be calculated by the number of cells in a row.
	tableCols int
	// tableWidth will be calculated by calcColWidths()
	tableWidth float64
}

// NewDocTable creats a DocTable given a *Doc and cells. An error will be
// returned when given cell dimensions are not valid (e.g. one row is shorter
// than the others).
func NewDocTable(doc *Doc, cells [][]string, args ...interface{}) (*DocTable, error) {
	// checking if every row has the same length.
	rowLen := len(cells[0])
	for i, r := range cells {
		if len(r) != rowLen {
			return nil, fmt.Errorf("row %v has mismatching columns: got: %v should: %v", i+1, len(r), rowLen)
		}
	}

	for _, r := range cells {
		for i := range r {
			r[i] = doc.trUTF8(r[i])
		}
	}

	t := &DocTable{
		doc:   doc,
		cells: cells,
	}
	t.tableRows = len(cells)
	t.tableCols = len(cells[0])

	// set table parameters when given via args
	for _, a := range args {
		switch param := a.(type) {
		case HeadType:
			t.headType = param
		}
	}

	t.SetDefaults()
	return t, nil
}

func (t *DocTable) SetDefaults() {
	rows := t.tableRows
	cols := t.tableCols
	// table parameters
	if t.headType == HeadUnset {
		t.headType = HeadNone
	}
	// column parameters
	if len(t.colTypes) == 0 {
		t.colTypes = array(cols, ColDyn)
	}
	if len(t.colFixedWidths) == 0 {
		t.colFixedWidths = array(cols, 0.)
	}
	if len(t.colGaps) == 0 {
		t.colGaps = array(cols-1, 0.)
	}
	// row parameters
	if len(t.rowTypes) == 0 {
		t.rowTypes = array(rows, RowCalc)
	}
	if len(t.rowFixedHeights) == 0 {
		t.rowFixedHeights = array(rows, 0.)
	}
	if len(t.rowGaps) == 0 {
		t.rowGaps = array(rows-1, 0.)
	}
	// cell parameters
	if len(t.cellTypes) == 0 {
		t.cellTypes = matrix(rows, cols, CellSingle)
	}
	if len(t.cellLineHeightFactors) == 0 {
		t.cellLineHeightFactors = matrix(rows, cols, 1.2)
	}
	if len(t.cellAligns) == 0 {
		t.cellAligns = matrix(rows, cols, AlignLeft)
	}
	if len(t.cellPaddings) == 0 {
		t.cellPaddings = matrix(rows, cols, Padding{0, 0, 0, 0})
	}
	if len(t.cellBorders) == 0 {
		t.cellBorders = matrix(rows, cols, true)
	}
	if len(t.cellStyleFuncs) == 0 {
		t.cellStyleFuncs = matrix[*func(gofpdf.Fpdf)](rows, cols, nil)
	}
}

func (t *DocTable) Generate() error {
	if err := t.validateColumns(); err != nil {
		return err
	}
	t.calcColWidths()

	if err := t.validateRows(); err != nil {
		return err
	}
	t.calcRowHeights()

	printWidth := t.doc.GetPrintWidth()
	if t.tableWidth > printWidth {
		return fmt.Errorf("error generating table: table wider than print width: %v > %v", t.tableWidth, printWidth)
	}
	// TODO: add save current style function to doc and run it here.
	for i := 0; i < t.tableRows; i++ {
		t.addRowGap(i) //adds only gap between rows - page break is not affected by gap as there is already a gap

		//check for nedt page
		nextRowHeight := t.rowHeights[i]
		if nextRowHeight > t.doc.GetRemainingPrintHeight() {
			t.doc.AddPage()

			if t.headType == HeadFirstRow {
				for j := 0; j < t.tableCols; j++ {
					t.addColGap(0, j)
					t.renderCell(0, j)
				}
			}
		}

		for j := 0; j < t.tableCols; j++ {
			t.addColGap(i, j)
			t.renderCell(i, j)
		}
	}
	// TODO: add restore saved style function to doc and run it here.

	return nil
}

func (t *DocTable) validateColumns() error {
	// check if calculated column has cells of type CellSingle
	for j := 0; j < t.tableCols; j++ {
		if t.colTypes[j] == ColCalc {
			hasCellSingle := false
			for i := 0; i < t.tableRows; i++ {
				if t.cellTypes[i][j] == CellSingle {
					hasCellSingle = true
					break
				}
			}
			if !hasCellSingle {
				return fmt.Errorf("column %v of type ColCalc has only CellMulti cells and cannot be calculated", j+1)
			}
		}
	}
	return nil
}

func (t *DocTable) calcColWidths() {
	t.colWidths = array(t.tableCols, .0)
	printWidth := t.doc.GetPrintWidth()
	tableWidth := 0.

	dynCount := 0
	for j, cType := range t.colTypes {
		if cType == ColFixed {
			t.colWidths[j] = t.colFixedWidths[j]
			tableWidth += t.colFixedWidths[j]
			continue
		}

		if cType == ColCalc {
			maxCellWidth := t.maxCellWidthOfColumn(j)
			t.colWidths[j] = maxCellWidth
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

	for j, cType := range t.colTypes {
		if cType == ColDyn {
			t.colWidths[j] = dynWidth
			tableWidth += dynWidth
		}
	}

	// set table width
	t.tableWidth = tableWidth
}

func (t *DocTable) maxCellWidthOfColumn(j int) float64 {
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

func (t *DocTable) validateRows() error {
	// check if a row of type RowFixed can render every cell completely
	for i := 0; i < t.tableRows; i++ {
		if t.rowTypes[i] == RowFixed {
			if t.getRowHeight(i) > t.rowFixedHeights[i] {
				return fmt.Errorf("row %v cannot display all cells; insufficient height", i+1)
			}
		}
	}
	return nil
}

func (t *DocTable) calcRowHeights() {
	rowHeights := array(t.tableRows, 0.)
	for i := 0; i < t.tableRows; i++ {
		if t.rowTypes[i] == RowFixed {
			rowHeights[i] = t.rowFixedHeights[i]
			continue
		}
		rowHeights[i] = t.getRowHeight(i)
	}
	t.rowHeights = rowHeights
}

func (t *DocTable) getRowHeight(i int) float64 {
	rowHt := 0.
	for j := 0; j < t.tableCols; j++ {
		h := t.getCellHeight(i, j)
		if h > rowHt {
			rowHt = h
		}
	}
	return rowHt
}

func (t *DocTable) getCellHeight(i, j int) float64 {
	lineHt := t.getLineHeight(i, j)
	p := t.cellPaddings[i][j]
	cellPadding := p[paddingTop] + p[paddingBottom]

	switch t.cellTypes[i][j] {
	case CellSingle:
		return lineHt + cellPadding
	case CellMulti:
		lines := t.doc.SplitText(t.cells[i][j], t.colWidths[j])
		return float64(len(lines))*lineHt + cellPadding
	default:
		panic("unsupported CellType: " + fmt.Sprint(t.cellTypes[i][j]))
	}
}

func (t *DocTable) getLineHeight(i, j int) float64 {
	_, fontHt := t.doc.GetFontSize()
	return fontHt * t.cellLineHeightFactors[i][j]
}

func (t *DocTable) addRowGap(i int) {
	if i > 0 && t.rowGaps[i-1] > 0 {
		t.doc.SetXY(t.doc.GetX(), t.doc.GetY()+t.rowGaps[i-1])
	}
}

func (t *DocTable) addColGap(i, j int) {
	if j > 0 && t.colGaps[j-1] != 0. {
		t.doc.SetX(t.doc.GetX() + t.colGaps[j-1])
	}
}

func (t *DocTable) renderCell(i, j int) {
	//check for page break

	t.doc.SetFillColor(255, 255, 255)
	oldColX, oldColY, oldColZ := t.doc.GetFillColor()

	//style
	if f := t.cellStyleFuncs[i][j]; f != nil {
		tmpFunc := *f
		tmpFunc(*t.doc.Fpdf)
	}

	//draw
	doc := t.doc
	x, y := doc.GetXY()
	p := t.cellPaddings[i][j]
	w := t.colWidths[j] - p[paddingLeft] - p[paddingRight]

	if t.cellBorders[i][j] {
		// TODO: Add tableCell with border parameter
		doc.Rect(x, y, t.colWidths[j], t.rowHeights[i], "FD")
	} else {
		doc.Rect(x, y, t.colWidths[j], t.rowHeights[i], "F")
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
		cellHt := t.rowHeights[i] - p[paddingTop] - p[paddingBottom]
		doc.SetFillColor(oldColX, oldColY, oldColZ)
		doc.CFormat(w, cellHt, cellStr, "", ln, alignStr, false, 0, "")
		doc.SetXY(doc.GetX()+p[paddingRight], y)
	case CellMulti:
		doc.SetFillColor(oldColX, oldColY, oldColZ)
		doc.MCell(w, t.getLineHeight(i, j), t.cells[i][j], "", alignStr, false)
		doc.SetXY(x+w+p[paddingLeft]+p[paddingRight], y)
	default:
		panic("unsupported CellType: " + fmt.Sprint(t.cellTypes[i][j]))
	}
	if j == t.tableCols-1 {
		doc.Ln(t.rowHeights[i])
	}
}

// TABLE PARAMETER SETTERS
func (t *DocTable) SetHeadType(ht HeadType) {
	if ht != HeadUnset {
		t.headType = ht
	}
}

//  COL SETTERS ----------------------------------------------------------------

func (t *DocTable) SetAllColTypes(ct ColumnType) {
	t.colTypes = array(t.tableCols, ct)
}
func (t *DocTable) SetColTypes(colTypes []ColumnType) error {
	if len(colTypes) != t.tableCols {
		return fmt.Errorf("column count mismatch: got: %v should: %v", len(colTypes), t.tableCols)
	}
	t.colTypes = colTypes
	return nil
}

func (t *DocTable) SetAllColFixedWidths(w float64) {
	t.colFixedWidths = array(t.tableCols, w)
}
func (t *DocTable) SetColFixedWidths(cFixedWidth []float64) error {
	if len(cFixedWidth) != t.tableCols {
		return fmt.Errorf("column count mismatch: got: %v should: %v", len(cFixedWidth), t.tableCols)
	}
	t.colFixedWidths = cFixedWidth
	return nil
}

func (t *DocTable) SetAllColGaps(g float64) {
	t.colGaps = array(t.tableCols-1, g)
}
func (t *DocTable) SetColGaps(g []float64) error {
	if len(g) != t.tableCols-1 {
		return fmt.Errorf("column count mismatch: got: %v should: %v", len(g), t.tableCols-1)
	}
	t.colGaps = g
	return nil
}

//  ROW SETTERS ----------------------------------------------------------------

func (t *DocTable) SetAllRowTypes(rt RowType) {
	t.rowTypes = array(t.tableRows, rt)
}
func (t *DocTable) SetRowTypes(rt []RowType) error {
	if len(rt) != t.tableRows {
		return fmt.Errorf("row count mismatch: got: %v should: %v", len(rt), t.tableRows-1)
	}
	t.rowTypes = rt
	return nil
}

func (t *DocTable) SetAllRowFixedHeights(f float64) {
	t.rowFixedHeights = array(t.tableRows, f)
}
func (t *DocTable) SetRowFixedHeights(f []float64) error {
	if len(f) != t.tableRows {
		return fmt.Errorf("row count mismatch: got: %v should: %v", len(f), t.tableRows-1)
	}
	t.rowFixedHeights = f
	return nil
}

func (t *DocTable) SetAllRowGaps(g float64) {
	t.rowGaps = array(t.tableRows-1, g)
}
func (t *DocTable) SetRowGaps(g []float64) error {
	if len(g) != t.tableRows-1 {
		return fmt.Errorf("row count mismatch: got: %v should: %v", len(g), t.tableRows-1)
	}
	t.rowGaps = g
	return nil
}

//  CELL SETTERS ----------------------------------------------------------------

func (t *DocTable) SetCell(i, j int, str string) error {
	if err := t.checkIndices(i, j); err != nil {
		return err
	}

	t.cells[i][j] = str
	return nil
}

func (t *DocTable) SetCellType(i, j int, ct CellType) error {
	if err := t.checkIndices(i, j); err != nil {
		return err
	}
	t.cellTypes[i][j] = ct
	return nil
}
func (t *DocTable) SetAllCellTypes(ct CellType) {
	t.cellTypes = matrix(t.tableRows, t.tableCols, ct)
}
func (t *DocTable) SetCellTypesPerColumn(ct []CellType) error {
	if len(ct) != t.tableCols {
		return fmt.Errorf("column count mismatch: got: %v should: %v", len(ct), t.tableCols)
	}

	for j := 0; j < t.tableCols; j++ {
		for i := 0; i < t.tableRows; i++ {
			t.cellTypes[i][j] = ct[j]
		}
	}
	return nil
}

func (t *DocTable) SetAllCellAligns(a AlignmentType) {
	t.cellAligns = matrix(t.tableRows, t.tableCols, a)
}
func (t *DocTable) SetCellAlingsPerColumn(a []AlignmentType) error {
	if len(a) != t.tableCols {
		return fmt.Errorf("column count mismatch: got: %v should: %v", len(a), t.tableCols)
	}

	for j := 0; j < t.tableCols; j++ {
		for i := 0; i < t.tableRows; i++ {
			t.cellAligns[i][j] = a[j]
		}
	}
	return nil
}

func (t *DocTable) SetAllCellPaddings(p Padding) {
	t.cellPaddings = matrix(t.tableRows, t.tableCols, p)
}
func (t *DocTable) SetCellPaddings(p [][]Padding) error {
	if len(p) != t.tableRows || len(p[0]) != t.tableCols {
		return fmt.Errorf("row or column count mismatch: got: (rows:%v cols:%v) table: (rows:%v cols%v)", len(p), len(p[0]), t.tableRows, t.tableCols)
	}
	t.cellPaddings = p
	return nil
}
func (t *DocTable) SetCellPaddingsPerColumn(p []Padding) error {
	if len(p) != t.tableCols {
		return fmt.Errorf("column count mismatch: got: %v should: %v", len(p), t.tableCols)
	}

	for j := 0; j < t.tableCols; j++ {
		for i := 0; i < t.tableRows; i++ {
			t.cellPaddings[i][j] = p[j]
		}
	}
	return nil
}

func (t *DocTable) SetAllCellLineHeightFactors(f float64) {
	t.cellLineHeightFactors = matrix(t.tableRows, t.tableCols, f)
}

func (t *DocTable) SetAllCellBorders(b bool) {
	t.cellBorders = matrix(t.tableRows, t.tableCols, b)
}

// TODO: Add SetCellStyleFunc
func (t *DocTable) SetAllCellStyleFuncs(f *func(gofpdf.Fpdf)) {
	t.cellStyleFuncs = matrix(t.tableRows, t.tableCols, f)
}
func (t *DocTable) SetCellStyleFuncsRow(i int, f *func(gofpdf.Fpdf)) error {
	if err := t.checkRowIndex(i); err != nil {
		return err
	}

	t.cellStyleFuncs[i] = array(t.tableCols, f)

	return nil
}
func (t *DocTable) SetCellStyleFuncsPerRow(fs []*func(gofpdf.Fpdf)) error {
	if len(fs) != t.tableRows {
		return fmt.Errorf("row count mismatch: got: %v should: %v", len(fs), t.tableRows)
	}

	for i := 0; i < t.tableRows; i++ {
		for j := 0; j < t.tableCols; j++ {
			t.cellStyleFuncs[i][j] = fs[i]
		}
	}

	return nil
}
func (t *DocTable) SetCellStyleFuncsPerAlternateRows(f1, f2 *func(gofpdf.Fpdf)) {
	// settings all rows with f1
	fs := array(t.tableRows, f1)

	// change every other row to f2
	for i := 1; i < t.tableRows; i += 2 {
		fs[i] = f2
	}

	// sanity check ;D
	if err := t.SetCellStyleFuncsPerRow(fs); err != nil {
		panic(err)
	}
}
func (t *DocTable) SetCellStyleFuncsPerColumn(fs []*func(gofpdf.Fpdf)) error {
	if len(fs) != t.tableCols {
		return fmt.Errorf("column count mismatch: got: %v should: %v", len(fs), t.tableCols)
	}

	for j := 0; j < t.tableCols; j++ {
		for i := 0; i < t.tableRows; i++ {
			t.cellStyleFuncs[i][j] = fs[j]
		}
	}

	return nil
}

// HELPER ----------------------------------------------------------------------

func (t *DocTable) checkRowIndex(i int) error {
	if i < 0 || i >= t.tableRows {
		return fmt.Errorf("invalid row index: got: %v should: 0-%v", i, t.tableRows-1)
	}
	return nil
}

func (t *DocTable) checkColIndex(j int) error {
	if j < 0 || j >= t.tableCols {
		return fmt.Errorf("invalid column index: got: %v should: 0-%v", j, t.tableCols-1)
	}
	return nil
}

func (t *DocTable) checkIndices(i, j int) error {
	if err := t.checkRowIndex(i); err != nil {
		return err
	}
	if err := t.checkColIndex(j); err != nil {

		return err
	}
	return nil
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

// alignToFpdf converts Alignments to the Fpdf equivalent(see below).
//
// alignStr specifies how the text is to be positioned within the cell.
// Horizontal alignment is controlled by including "L", "C" or "R" (left,
// center, right) in alignStr. Vertical alignment is controlled by including
// "T", "M", "B" or "A" (top, middle, bottom, baseline) in alignStr. The default
// alignment is left middle.
func alignToFpdf(a AlignmentType) string {
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
