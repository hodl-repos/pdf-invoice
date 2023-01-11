package pdfhelper

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"image"
	"io"
	"os"

	"github.com/jung-kurt/gofpdf"
)

type Image struct {
	ImageName   string `json:"image_name"`             // ImageName of the image
	ImagePath   string `json:"image_path,omitempty"`   // ImagePath is the location of the image on the filesystem
	ImageString string `json:"image_string,omitempty"` // ImageString is the encoded string of the image
}

// FilePNGSize returns the width & height of a given png file from path.
func FilePNGSize(path string) (float64, float64, error) {
	file, err := os.Open(path)
	if err != nil {
		return -1, -1, err
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		return -1, -1, err
	}

	bounds := img.Bounds()
	width := bounds.Max.X - bounds.Min.X
	height := bounds.Max.Y - bounds.Min.Y
	fmt.Printf("Width: %d, Height: %d\n", width, height)
	return float64(width), float64(height), nil
}

func AddPNG(pdf *gofpdf.Fpdf, img *Image) (*gofpdf.ImageInfoType, error) {
	if img.ImagePath == "" && img.ImageString == "" {
		return nil, fmt.Errorf("error adding png: no image path or string")
	}

	var file io.Reader

	if img.ImagePath != "" {
		f, err := os.Open(img.ImagePath)
		if err != nil {
			return nil, fmt.Errorf("error opening image_path: %v", err)
		}
		defer f.Close()

		file = f
	}

	if img.ImageString != "" {
		decodedBytes, err := base64.StdEncoding.DecodeString(img.ImageString)
		if err != nil {
			return nil, fmt.Errorf("error decoding image_string: %v", err)
		}
		file = bytes.NewReader(decodedBytes)
	}

	return pdf.RegisterImageOptionsReader(img.ImageName, gofpdf.ImageOptions{ImageType: "png", ReadDpi: true}, file), nil
}
