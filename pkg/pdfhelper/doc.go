package pdfhelper

import "github.com/jung-kurt/gofpdf"

type Doc struct {
	*gofpdf.Fpdf
	Lang   string
	trUTF8 func(string) string
}

const (
	DOC_LANG_DEFAULT = "en"
)

func NewDocA4() *Doc {
	doc := &Doc{}
	doc.Fpdf = NewA4()
	doc.Lang = DOC_LANG_DEFAULT
	doc.trUTF8 = doc.UnicodeTranslatorFromDescriptor("")
	return doc
}

func (d *Doc) SetLanguage(lang string) {
	d.Lang = lang
}

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
