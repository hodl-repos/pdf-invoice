package mcoinvoice

import (
	"io/fs"
	"os"
	"testing"
)

func TestGenerate(t *testing.T) {
	i := New()
	b, err := os.ReadFile("./test-invoice.json")
	if err != nil {
		t.Fatal("error reading test-invoice.json:", err)
	}

	err = i.SetParams(b)
	if err != nil {
		t.Fatal("error setting params:", err)
	}

	pdf, err := i.Generate()
	if err != nil {
		t.Fatal("error generating pdf:", err)
	}

	err = os.WriteFile("test.pdf", pdf, fs.FileMode(0777))
	if err != nil {
		t.Fatal("error writing file to disk:", err)
	}
}
