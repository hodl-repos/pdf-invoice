package pdfhelper

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/jung-kurt/gofpdf"
)

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

func CreatePDFInRootOutFolder(pdf *gofpdf.Fpdf, fileName string) error {
	srcPath, err := getSrcPathInToRootOutFolder(fileName)
	if err != nil {
		return fmt.Errorf("could not find out folder: %v", err)
	}

	file, err := os.OpenFile(srcPath, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		if os.IsNotExist(err) {
			if err := os.MkdirAll(filepath.Dir(srcPath), os.ModePerm); err != nil {
				return err
			}
			file, err = os.OpenFile(srcPath, os.O_CREATE|os.O_WRONLY, 0644)
			if err != nil {
				return err
			}
		} else {
			fmt.Println(err)
			return err
		}
	}
	defer file.Close()

	pdf.Output(file)
	file.Close()

	return nil
}

func getSrcPathInToRootOutFolder(fileName string) (string, error) {
	// Get the current working directory
	currentDir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	// Evaluate symbolic links
	rootDir, err := filepath.EvalSymlinks(currentDir)
	if err != nil {
		return "", err
	}

	// Search up the directory tree until finding the root
	for {
		if _, err := os.Stat(filepath.Join(rootDir, "go.mod")); err == nil {
			break
		}
		prev := rootDir
		rootDir = filepath.Dir(rootDir)
		if rootDir == prev {
			return "", fmt.Errorf("no go.mod file found")
		}
	}

	// add out folder
	rootDir = filepath.Join(rootDir, "out")

	// Get the relative path from the current working directory to the root folder
	relPath, err := filepath.Rel(currentDir, rootDir)
	if err != nil {
		return "", err
	}

	var srcPath string

	if !strings.HasPrefix(fileName, "/") {
		srcPath = filepath.Join("/", fileName)
	} else {
		srcPath = fileName
	}

	srcPath = filepath.Join(relPath, srcPath)

	fmt.Println(srcPath)

	//  You can use the srcPath here to do something with the PDF.
	return srcPath, nil
}
