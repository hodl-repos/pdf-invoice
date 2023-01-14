package document

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/jung-kurt/gofpdf"
)

// TODO: add functionality to generate example files from tests.
func CreatePDFInProjectRootOutFolder(pdf *gofpdf.Fpdf, fileName string) error {
	srcPath, err := getSrcPathToProjectRootOutFolder(fileName)
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

func getSrcPathToProjectRootOutFolder(fileName string) (string, error) {
	// Get the current working directory
	currentDir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	// Evaluate symbolic links
	projectRootDir, err := filepath.EvalSymlinks(currentDir)
	if err != nil {
		return "", err
	}

	// Search up the directory tree until finding the root
	for {
		if _, err := os.Stat(filepath.Join(projectRootDir, "go.mod")); err == nil {
			break
		}
		prev := projectRootDir
		projectRootDir = filepath.Dir(projectRootDir)
		if projectRootDir == prev {
			return "", fmt.Errorf("no go.mod file found")
		}
	}

	// add out folder
	projectRootDir = filepath.Join(projectRootDir, "out")

	// Get the relative path from the current working directory to the root folder
	relPath, err := filepath.Rel(currentDir, projectRootDir)
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
