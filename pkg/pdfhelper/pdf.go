package pdfhelper

import "github.com/jung-kurt/gofpdf"

// NewA4 creates a new pdf in DIN A4 format with one page added.
//
// Orientation: portrait
//
// Unit: mm
//
// Size: A4
//
// Font: Arial
//
// FontSize: 8
//
// Margins: left: 10, top: 10, right: 10
func NewA4() *gofpdf.Fpdf {
	pdf := gofpdf.NewCustom(&gofpdf.InitType{
		OrientationStr: "P",
		UnitStr:        "mm",
		SizeStr:        gofpdf.PageSizeA4,
		FontDirStr:     "",
	})

	pdf.SetFont("Arial", "", 8)
	pdf.SetMargins(10, 10, 10)
	// pdf.SetAutoPageBreak(true, 10)

	pdf.AddPage()

	return pdf
}
