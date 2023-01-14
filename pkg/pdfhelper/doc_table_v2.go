package pdfhelper

import "fmt"

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

// DocTableV2
//
// NOTE: All used indices follow the convention that indices i are for rows and
// indices j are for columns.
//
// Render process: 1-3 user actions
// 1) NewDocTableV2
//   - setting row and column count
//   - setting default values
//
// 2) (optional) change table parameters
// 3) Generate
//   - validate columns
//   - calculate columns
//   - validate rows
//   - calculate rows
//   - render cells
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
	//cellBorders determine for every cell if a border should be displayed.
	cellBorders [][]bool

	// tableRows will be calculated by the number of rows of cells.
	tableRows int
	// tableCols will be calculated by the number of cells in a row.
	tableCols int
	// tableWidth will be calculated by calcColWidths()
	tableWidth float64
}

// NewDocTableV2 creats a DocTableV2 given a *Doc and cells. An error will be
// returned when given cell dimensions are not valid (e.g. one row is shorter
// than the others).
func NewDocTableV2(doc *Doc, cells [][]string) (*DocTableV2, error) {
	// checking if every row has the same length.
	rowLen := len(cells[0])
	for i, r := range cells {
		if len(r) != rowLen {
			return nil, fmt.Errorf("row %v has mismatching columns: got: %v should: %v", i+1, len(r), rowLen)
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
}

func (t *DocTableV2) Generate() error {
	if err := t.validateColumns(); err != nil {
		return err
	}
	t.calcColWidths()

	if err := t.validateRows(); err != nil {
		return err
	}
	t.calcRowHeights()

	printWidth := GetPrintWidth(t.doc.Fpdf)
	if t.tableWidth > printWidth {
		return fmt.Errorf("error generating table: table wider than print width: %v > %v", t.tableWidth, printWidth)
	}
	for i := 0; i < t.tableRows; i++ {
		t.addRowGap(i)
		for j := 0; j < t.tableCols; j++ {
			t.addColGap(i, j)
			t.renderCell(i, j)
		}
	}

	return nil
}

func (t *DocTableV2) validateColumns() error {
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

func (t *DocTableV2) calcColWidths() {
	t.colWidths = array(t.tableCols, .0)
	printWidth := GetPrintWidth(t.doc.Fpdf)
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

func (t *DocTableV2) validateRows() error {
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

func (t *DocTableV2) calcRowHeights() {
	rowHeights := array(t.tableRows, 0.)
	for i := 0; i < t.tableRows; i++ {
		rowHeights[i] = t.getRowHeight(i)
	}
	t.rowHeights = rowHeights
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
		lines := t.doc.SplitText(t.cells[i][j], t.colWidths[j])
		return float64(len(lines))*lineHt + cellPadding
	default:
		panic("unsupported CellType: " + fmt.Sprint(t.cellTypes[i][j]))
	}
}

func (t *DocTableV2) getLineHeight(i, j int) float64 {
	_, fontHt := t.doc.GetFontSize()
	return fontHt * t.cellLineHeightFactors[i][j]
}

func (t *DocTableV2) addRowGap(i int) {
	if i > 0 && t.rowGaps[i-1] > 0 {
		t.doc.SetXY(t.doc.GetX(), t.doc.GetY()+t.rowGaps[i-1])
	}
}

func (t *DocTableV2) addColGap(i, j int) {
	if j > 0 && t.colGaps[j-1] != 0. {
		t.doc.SetX(t.doc.GetX() + t.colGaps[j-1])
	}
}

func (t *DocTableV2) renderCell(i, j int) {
	doc := t.doc
	x, y := doc.GetXY()
	p := t.cellPaddings[i][j]
	w := t.colWidths[j] - p[paddingLeft] - p[paddingRight]

	if t.cellBorders[i][j] {
		// TODO: Add tableCell with border parameter
		doc.Rect(x, y, t.colWidths[j], t.rowHeights[i], "D")
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
		doc.CFormat(w, cellHt, cellStr, "", ln, alignStr, false, 0, "")
		doc.SetXY(doc.GetX()+p[paddingRight], y)
	case CellMulti:
		doc.MCell(w, t.getLineHeight(i, j), t.cells[i][j], "", alignStr, false)
		doc.SetXY(x+w+p[paddingLeft]+p[paddingRight], y)
	default:
		panic("unsupported CellType: " + fmt.Sprint(t.cellTypes[i][j]))
	}
	if j == t.tableCols-1 {
		doc.Ln(t.rowHeights[i])
	}
}

//  COL SETTERS ----------------------------------------------------------------

func (t *DocTableV2) SetAllColTypes(ct ColumnType) {
	t.colTypes = array(t.tableCols, ct)
}
func (t *DocTableV2) SetColTypes(colTypes []ColumnType) error {
	if len(colTypes) != t.tableCols {
		return fmt.Errorf("column count mismatch: got: %v should: %v", len(colTypes), t.tableCols)
	}
	t.colTypes = colTypes
	return nil
}

func (t *DocTableV2) SetAllColFixedWidths(w float64) {
	t.colFixedWidths = array(t.tableCols, w)
}
func (t *DocTableV2) SetColFixedWidths(cFixedWidth []float64) error {
	if len(cFixedWidth) != t.tableCols {
		return fmt.Errorf("column count mismatch: got: %v should: %v", len(cFixedWidth), t.tableCols)
	}
	t.colFixedWidths = cFixedWidth
	return nil
}

func (t *DocTableV2) SetAllColGaps(g float64) {
	t.colGaps = array(t.tableCols-1, g)
}
func (t *DocTableV2) SetColGaps(g []float64) error {
	if len(g) != t.tableCols-1 {
		return fmt.Errorf("column count mismatch: got: %v should: %v", len(g), t.tableCols-1)
	}
	t.colGaps = g
	return nil
}

//  ROW SETTERS ----------------------------------------------------------------

func (t *DocTableV2) SetAllRowTypes(rt RowType) {
	t.rowTypes = array(t.tableRows, rt)
}
func (t *DocTableV2) SetRowTypes(rt []RowType) error {
	if len(rt) != t.tableRows {
		return fmt.Errorf("row count mismatch: got: %v should: %v", len(rt), t.tableRows-1)
	}
	t.rowTypes = rt
	return nil
}

func (t *DocTableV2) SetAllRowFixedHeights(f float64) {
	t.rowFixedHeights = array(t.tableRows, f)
}
func (t *DocTableV2) SetRowFixedHeights(f []float64) error {
	if len(f) != t.tableRows {
		return fmt.Errorf("row count mismatch: got: %v should: %v", len(f), t.tableRows-1)
	}
	t.rowFixedHeights = f
	return nil
}

func (t *DocTableV2) SetAllRowGaps(g float64) {
	t.rowGaps = array(t.tableRows-1, g)
}
func (t *DocTableV2) SetRowGaps(g []float64) error {
	if len(g) != t.tableRows-1 {
		return fmt.Errorf("row count mismatch: got: %v should: %v", len(g), t.tableRows-1)
	}
	t.rowGaps = g
	return nil
}

//  CELL SETTERS ----------------------------------------------------------------

func (t *DocTableV2) SetCell(i, j int, str string) error {
	if err := t.checkIndices(i, j); err != nil {
		return err
	}

	t.cells[i][j] = str
	return nil
}

func (t *DocTableV2) SetCellType(i, j int, ct CellType) error {
	if err := t.checkIndices(i, j); err != nil {
		return err
	}
	t.cellTypes[i][j] = ct
	return nil
}
func (t *DocTableV2) SetAllCellTypes(ct CellType) {
	t.cellTypes = matrix(t.tableRows, t.tableCols, ct)
}
func (t *DocTableV2) SetCellTypesPerColumn(ct []CellType) error {
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

func (t *DocTableV2) SetAllCellAligns(a CellAlignment) {
	t.cellAligns = matrix(t.tableRows, t.tableCols, a)
}
func (t *DocTableV2) SetCellAlingsPerColumn(a []CellAlignment) error {
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
func (t *DocTableV2) SetCellPaddingsPerColumn(p []Padding) error {
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

func (t *DocTableV2) SetAllCellLineHeightFactors(f float64) {
	t.cellLineHeightFactors = matrix(t.tableRows, t.tableCols, f)
}

func (t *DocTableV2) SetAllCellBorders(b bool) {
	t.cellBorders = matrix(t.tableRows, t.tableCols, b)
}

// HELPER ----------------------------------------------------------------------

func (t *DocTableV2) checkIndices(i, j int) error {
	if i < 0 || i >= t.tableRows {
		return fmt.Errorf("invalid row index: got: %v should: 0-%v", i, t.tableRows-1)
	}
	if j < 0 || j >= t.tableCols {
		return fmt.Errorf("invalid column index: got: %v should: 0-%v", j, t.tableCols-1)
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
