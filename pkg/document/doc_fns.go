package document

// CFormat wraps CellFormat and transforms the given txtStr in a UTF-8
// UnicodeTranslator to render special characters such as german Umlaute
func (d *Doc) CFormat(w, h float64, txtStr, borderStr string, ln int,
	alignStr string, fill bool, link int, linkStr string) {
	d.CellFormat(w, h, d.trUTF8(txtStr), borderStr, ln, alignStr, fill, link, linkStr)
}

// CFormat wraps MultiCell and transforms the given txtStr in a UTF-8
// UnicodeTranslator to render special characters such as german Umlaute
func (d *Doc) MCell(w, h float64, txtStr, borderStr, alignStr string, fill bool) {
	d.MultiCell(w, h, d.trUTF8(txtStr), borderStr, alignStr, fill)
}

// Ellipsis are three dots (...) representing that a string is longer than it is
// displayed
func (d *Doc) Ellipsis() string {
	return "..."
}

// VerticalLine draws a vertical line on the current y position from x1 to x2
// with a given thinkness.
func (d *Doc) VerticalLine(x1, x2, thickness float64) {
	d.SetLineWidth(thickness)

	y := d.GetY()
	d.Line(x1, y, x2, y)
}

// VerticalLinePrintWidth draws a vertical line inside the current print width
// with a thickness of Fpdf.GetLineWidth().
func (d *Doc) VerticalLinePrintWidth() {
	ml, _, _, _ := d.GetMargins()
	w := d.GetPrintWidth()
	d.VerticalLine(ml, ml+w, d.GetLineWidth())
}
