package utils

import (
	"os"

	"github.com/jung-kurt/gofpdf"
)

const logoName = "logo.png"

func RegisterLogo(pdf *gofpdf.Fpdf) {
	logoFile, err := os.Open(logoName)
	if err != nil {
		panic(err)
	}
	pdf.RegisterImageOptionsReader(logoName, gofpdf.ImageOptions{ImageType: "png", ReadDpi: true}, logoFile)
}

func WriteLogo(pdf *gofpdf.Fpdf) {
	x, y := pdf.GetXY()
	pdf.ImageOptions(logoName, x-0.8, y, 60, 0, true, gofpdf.ImageOptions{ReadDpi: true}, 0, "")
}
