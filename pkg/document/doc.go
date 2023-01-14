package document

import "github.com/jung-kurt/gofpdf"

type Doc struct {
	*gofpdf.Fpdf
	lang   string
	trUTF8 func(string) string
}

const (
	DOC_LANG_DEFAULT = "en"
)

// NewA4 creates a new pdf in DIN A4 format with one page added.
//
// Orientation: portrait
//
// lang: en
//
// unit: mm
//
// size: A4
//
// font: Arial
//
// fontSize: 8
//
// document margins: left: 10, top: 10, right: 10
//
// line width: 0.2
func NewA4() *Doc {
	doc := &Doc{}
	doc.Fpdf = newA4()
	doc.lang = DOC_LANG_DEFAULT
	doc.trUTF8 = doc.UnicodeTranslatorFromDescriptor("")
	return doc
}

func (d *Doc) SetLanguage(lang string) {
	d.lang = lang
}
func (d *Doc) GetLanguage() string {
	return d.lang
}

func newA4() *gofpdf.Fpdf {
	pdf := gofpdf.NewCustom(&gofpdf.InitType{
		OrientationStr: "P",
		UnitStr:        "mm",
		SizeStr:        gofpdf.PageSizeA4,
		FontDirStr:     "",
	})

	pdf.SetFont("Arial", "", 8)
	pdf.SetMargins(10, 10, 10)
	pdf.SetCellMargin(0)
	pdf.SetLineWidth(0.2)
	// pdf.SetAutoPageBreak(true, 10)

	pdf.AddPage()

	return pdf
}
