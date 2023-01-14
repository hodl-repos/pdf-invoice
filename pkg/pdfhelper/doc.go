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
