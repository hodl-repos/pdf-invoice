package block

import (
	"os"
	"testing"

	"github.com/hodl-repos/pdf-invoice/pkg/pdfhelper"
)

func TestAddLogoBlockDefault(t *testing.T) {
	pdf := pdfhelper.NewA4()

	data := []byte(`{ "image": null, "width": 0, "height": 0 }`)

	logo, err := NewLogoFromJSON(data)
	if err != nil {
		t.Fatal("error creating new logo from json:", err)
	}

	AddLogoBlock(pdf, logo)
	pdfhelper.CreatePDFInProjectRootOutFolder(pdf, "TestAddLogoBlockDefault.pdf")
}

func TestAddLogoBlockDefaultImage(t *testing.T) {
	pdf := pdfhelper.NewA4()

	data := []byte(`{ "image": null, "width": 45, "height": 0 }`)

	logo, err := NewLogoFromJSON(data)
	if err != nil {
		t.Fatal("error creating new logo from json:", err)
	}

	AddLogoBlock(pdf, logo)
	pdfhelper.CreatePDFInProjectRootOutFolder(pdf, "TestAddLogoBlockDefaultImage.pdf")
}

func TestAddLogoBlockImagePath(t *testing.T) {
	pdf := pdfhelper.NewA4()

	data := []byte(`{ 
		"image": { 
			"image_name": "example_logo", 
			"image_path": "test/example_logo.png" 
		}, 
		"width": 0, 
		"height": 15 
	}`)

	logo, err := NewLogoFromJSON(data)
	if err != nil {
		t.Fatal("error creating new logo from json:", err)
	}

	AddLogoBlock(pdf, logo)
	pdfhelper.CreatePDFInProjectRootOutFolder(pdf, "TestAddLogoBlockImagePath.pdf")
}

func TestAddLogoBlockImageString(t *testing.T) {
	pdf := pdfhelper.NewA4()

	data, err := os.ReadFile("test/TestAddLogoBlockImageString.json")
	if err != nil {
		t.Fatal("error reading json file:", err)
	}

	logo, err := NewLogoFromJSON(data)
	if err != nil {
		t.Fatal("error creating new logo from json:", err)
	}

	AddLogoBlock(pdf, logo)
	pdfhelper.CreatePDFInProjectRootOutFolder(pdf, "TestAddLogoBlockImageString.pdf")
}
