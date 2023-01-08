package pdf

import "github.com/jung-kurt/gofpdf"

func NewA4() *gofpdf.Fpdf {
	pdf := gofpdf.NewCustom(&gofpdf.InitType{
		OrientationStr: "P",
		UnitStr:        "mm",
		SizeStr:        gofpdf.PageSizeA4,
		FontDirStr:     "",
	})

	pdf.SetFont("Arial", "", 8)
	pdf.SetMargins(10, 10, 10)
	pdf.SetAutoPageBreak(true, 10)

	return pdf
}

func NewA6Landscape() *gofpdf.Fpdf {
	pdf := gofpdf.NewCustom(&gofpdf.InitType{
		OrientationStr: "P",
		UnitStr:        "mm",
		SizeStr:        "A6", //should fail because not present, dont forget sizetype
		Size: gofpdf.SizeType{
			Wd: 148.5,
			Ht: 105,
		},
		FontDirStr: "",
	})

	pdf.SetFont("Arial", "", 8)

	return pdf
}

func New(width, height float64) *gofpdf.Fpdf {
	pdf := gofpdf.NewCustom(&gofpdf.InitType{
		OrientationStr: "P",
		UnitStr:        "mm",
		SizeStr:        "custom", //should fail because not present, dont forget sizetype
		Size: gofpdf.SizeType{
			Wd: width,
			Ht: height,
		},
		FontDirStr: "",
	})

	pdf.SetFont("Arial", "", 8)
	pdf.SetMargins(0.0, 0.0, 0.0)

	return pdf
}
