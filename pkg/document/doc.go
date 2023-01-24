package document

import "github.com/jung-kurt/gofpdf"

type Doc struct {
	*gofpdf.Fpdf
	lang string
	// lineHeight determines the height of one line of text given the current
	// font size. lineHeight is in percent where 1 = 100%. Percent is used so it
	// is independent of the used unit (pt, mm, in, etc.). The height of one line
	// of text is calculated by fontSize (in units) * lineHeight. Default is 1.2.
	lineHeight float64
	trUTF8     func(string) string
	debug      bool
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
// lineHeight: 1.2
//
// document margins: left: 10, top: 10, right: 10
//
// line width: 0.2
func NewA4() *Doc {
	doc := &Doc{}
	doc.Fpdf = newA4()
	doc.lang = DOC_LANG_DEFAULT
	doc.lineHeight = 1.2
	doc.trUTF8 = doc.UnicodeTranslatorFromDescriptor("")
	doc.debug = false
	return doc
}

func (d *Doc) SetLanguage(lang string) {
	d.lang = lang
}
func (d *Doc) GetLanguage() string {
	return d.lang
}

// SetLineHeight sets the line height. Values 0 and lower will be disgarded.
func (d *Doc) SetLineHeight(lh float64) {
	if lh > 0 {
		d.lineHeight = lh
	}
}
func (d *Doc) GetLineHeight() float64 {
	return d.lineHeight
}
func (d *Doc) GetFontLineHeight() float64 {
	_, fontHt := d.GetFontSize()
	return fontHt * d.lineHeight
}

// SetDebug sets the debug flag for the whole document.
func (d *Doc) SetDebug(b bool) {
	d.debug = b
}
func (d *Doc) Debug() bool {
	return d.debug
}

// GetPrintWidth returns the current print width, which is the page width
// subtracted by the left and right margin.
func (d *Doc) GetPrintWidth() float64 {
	pageWidth, _ := d.Fpdf.GetPageSize()
	marginL, _, marginR, _ := d.Fpdf.GetMargins()
	return pageWidth - marginL - marginR
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
