package block

import (
	"testing"

	"github.com/hodl-repos/pdf-invoice/pkg/pdfhelper"
)

func TestAddInvoiceTableBlock(t *testing.T) {
	doc := pdfhelper.NewDocA4()

	doc.SetLanguage("de")

	data := []byte(`{
		"sum_net": 100.0,
		"sum_gross": 115.0,
		"items": [
			{
					"name": "1 Sorglospaket (Anzug 3-teilig, 2 Maßhemden, Krawatte oder Fliege und Stecktuch)",
					"net": 50.0,
					"tax_rate": 0.15,
					"tax": 7.5,
					"gross": 57.5
			},
			{
					"name": "Item 2",
					"net": 50.0,
					"tax_rate": 0.15,
					"tax": 7.5,
					"gross": 57.5
			}
		]
	}`)

	invoiceTable, err := NewInvoiceTableFromJSON(data)
	if err != nil {
		t.Fatal("error creating new logo from json:", err)
	}

	AddInvoiceTableBlock(doc, invoiceTable)
	pdfhelper.CreatePDFInProjectRootOutFolder(doc.Fpdf, "TestAddInvoiceTableBlock.pdf")
}
