package pdfhelper

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
