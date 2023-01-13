package pdfhelper

type DocTableHeader struct {
	doc *Doc

	header []string
	// Columns can be either fixed "f" or dynamic "d"
	cols []string
	// colGap is an additional width added to a fixed column. On dynamic columns
	// no gap will be added.
	colGaps   []float64
	colAligns []string

	height    float64
	borderStr string

	// colWidths are the calculated column widths
	colWidths []float64
}

func NewDocTableHeader(doc *Doc, header, cols []string, colGaps []float64, colAligns []string) *DocTableHeader {
	th := &DocTableHeader{
		doc:       doc,
		header:    header,
		cols:      cols,
		colGaps:   colGaps,
		colAligns: colAligns,
	}
	th.setDefaults()

	return th
}

func (th *DocTableHeader) setDefaults() {
	_, unitSize := th.doc.GetFontSize()
	th.height = unitSize * 1.2
	th.borderStr = ""
}

func (th *DocTableHeader) SetHeight(h float64) {
	th.height = h
}

func (th *DocTableHeader) SetBorderStr(str string) {
	th.borderStr = str
}

func (th *DocTableHeader) calcColWidths() {
	th.colWidths = make([]float64, len(th.header))
	printWidth := GetPrintWidth(th.doc.Fpdf)
	remWidth := printWidth

	dynCount := 0
	for i, c := range th.cols {
		switch c {
		case "f":
			w := th.doc.GetStringWidth(th.header[i]) + th.colGaps[i]
			remWidth -= w
			th.colWidths[i] = w
		case "d":
			dynCount++
		default:
			panic("invalid column type: " + c)
		}
	}
	dynWidth := remWidth / float64(dynCount)

	for i, c := range th.cols {
		if c == "d" {
			th.colWidths[i] = dynWidth
		}
	}
}

func (th *DocTableHeader) GetColWidths() []float64 {
	th.calcColWidths()
	return th.colWidths
}

func (th *DocTableHeader) Generate() {
	doc := th.doc
	th.calcColWidths()

	for i, h := range th.header {
		doc.CFormat(th.colWidths[i], th.height, h, th.borderStr, 0, th.colAligns[i], false, 0, "")
	}
	doc.Ln(th.height)
}

type DocTable struct {
	doc *Doc

	rows      [][]string
	cols      []float64
	colAligns []string

	cellMarginTop    float64
	cellMarginBottom float64
	lineHeightFactor float64 // 1 = 100% font size height in pdf units
	borderStr        string
}

func NewDocTable(doc *Doc, rows [][]string, cols []float64, colAligns []string) *DocTable {
	dt := &DocTable{
		doc:       doc,
		rows:      rows,
		cols:      cols,
		colAligns: colAligns,
	}
	dt.setDefaults()

	return dt
}

func (dt *DocTable) setDefaults() {
	dt.cellMarginTop = 1
	dt.cellMarginBottom = 1
	dt.borderStr = ""
	dt.lineHeightFactor = 1.2 // setting only the factor, so that we always use the current font size
}

func (dt *DocTable) SetCellMarginTop(m float64) {
	dt.cellMarginTop = m
}

func (dt *DocTable) SetCellMarginBottom(m float64) {
	dt.cellMarginBottom = m
}

func (dt *DocTable) SetlineHeightFactor(f float64) {
	dt.lineHeightFactor = f
}

func (dt *DocTable) SetBorderStr(bStr string) {
	dt.borderStr = bStr
}

func (dt *DocTable) Generate() {
	doc := dt.doc
	_, fontHt := doc.GetFontSize()
	lineHt := fontHt * dt.lineHeightFactor

	for _, row := range dt.rows {
		var rowHt = lineHt
		// get row height
		for i, c := range row {
			lines := doc.SplitText(c, dt.cols[i])
			cellHt := float64(len(lines))*lineHt + dt.cellMarginTop + dt.cellMarginBottom
			if cellHt > rowHt {
				rowHt = cellHt
			}
		}

		// rendering row
		for i, cStr := range row {
			x, y := doc.GetXY()

			// rendering cell border when borderStr = "1"
			if dt.borderStr == "1" {
				doc.Rect(x, y, dt.cols[i], rowHt, "D")
			}

			doc.SetXY(x, y+dt.cellMarginTop)
			doc.MCell(dt.cols[i], lineHt, cStr, "", dt.colAligns[i], false)
			if i < len(dt.cols)-1 {
				doc.SetXY(x+dt.cols[i], y)
			} else {
				doc.SetY(y + rowHt)
			}
		}
	}
}
