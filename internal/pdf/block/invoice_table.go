package block

import (
	"encoding/json"
	"fmt"

	"github.com/hodl-repos/pdf-invoice/pkg/pdfhelper"
	"github.com/jung-kurt/gofpdf"
)

func init() {
	var translations = map[string]map[string]string{
		"de": {
			"AMOUNT_IN_EUR_EXCL_VAT": "Betrag in EUR zzgl. USt.",
			"VAT":                    "USt.",
			"AMOUNT_VAT":             "USt. Betrag",
			"AMOUNT_IN_EUR_INCL_VAT": "Betrag in EUR inkl. USt.",
		},
		"en": {
			"AMOUNT_IN_EUR_EXCL_VAT": "Amount in EUR excl. VAT",
			"VAT":                    "VAT",
			"AMOUNT_VAT":             "VAT Amount",
			"AMOUNT_IN_EUR_INCL_VAT": "Amount in EUR incl. VAT",
		},
	}

	err := pdfhelper.LoadTranslation(translations)
	if err != nil {
		panic(err)
	}
}

// InvoiceTableItem represents a single item in an invoice table
type InvoiceTableItem struct {
	Name     string  `json:"name"`               // Required field representing the name of the item
	Net      float32 `json:"net,omitempty"`      // Net value of the item, omitted if zero
	Tax_Rate float32 `json:"tax_rate,omitempty"` // Tax rate applied to the item, omitted if zero
	Tax      float32 `json:"tax,omitempty"`      // Tax value calculated from the net and tax rate, omitted if zero
	Gross    float32 `json:"gross,omitempty"`    // Gross value calculated from net and tax, omitted if zero
}

// InvoiceTable represents a table of invoice items
type InvoiceTable struct {
	SumNet   float32            `json:"sum_net,omitempty"`   // Sum of net values of all items, omitted if zero
	SumGross float32            `json:"sum_gross,omitempty"` // Sum of gross values of all items, omitted if zero
	Items    []InvoiceTableItem `json:"items"`               // Slice of invoice table items
}

// NewInvoiceTableFromJSON creates a new InvoiceTable from json byte data.
func NewInvoiceTableFromJSON(data []byte) (*InvoiceTable, error) {
	invoiceTable := &InvoiceTable{}
	err := json.Unmarshal(data, invoiceTable)
	if err != nil {
		return nil, err
	}
	return invoiceTable, nil
}

// AddInvoiceTableBlock  create a table in pdf using information from an invoice
// table struct.
func AddInvoiceTableBlock(pdf *gofpdf.Fpdf, invoiceTable *InvoiceTable) {
	pdfhelper.DrawVerticalLinePrintWidth(pdf)

	// w := pdfhelper.GetPrintWidth(pdf)

	// Add the table headers
	pdf.CellFormat(40, 10, "Name", "1", 0, "C", false, 0, "")
	pdf.CellFormat(20, 10, pdfhelper.T("AMOUNT_IN_EUR_EXCL_VAT", "de"), "1", 0, "C", false, 0, "")
	pdf.CellFormat(20, 10, "Tax Rate", "1", 0, "C", false, 0, "")
	pdf.CellFormat(20, 10, "Tax", "1", 0, "C", false, 0, "")
	pdf.CellFormat(20, 10, "Gross", "1", 1, "C", false, 0, "")
	// Loop through each item in the invoice table
	for _, item := range invoiceTable.Items {
		// Add the item information to the table
		pdf.CellFormat(40, 10, item.Name, "1", 0, "C", false, 0, "")
		pdf.CellFormat(20, 10, fmt.Sprintf("%.2f", item.Net), "1", 0, "C", false, 0, "")
		pdf.CellFormat(20, 10, fmt.Sprintf("%.2f", item.Tax_Rate), "1", 0, "C", false, 0, "")
		pdf.CellFormat(20, 10, fmt.Sprintf("%.2f", item.Tax), "1", 0, "C", false, 0, "")
		pdf.CellFormat(20, 10, fmt.Sprintf("%.2f", item.Gross), "1", 1, "C", false, 0, "")
	}
	// Add the sum net and gross to the table
	pdf.CellFormat(40, 10, "Sum Net", "1", 0, "C", false, 0, "")
	pdf.CellFormat(60, 10, fmt.Sprintf("%.2f", invoiceTable.SumNet), "1", 0, "C", false, 0, "")
	pdf.CellFormat(20, 10, "Sum Gross", "1", 0, "C", false, 0, "")
	pdf.CellFormat(20, 10, fmt.Sprintf("%.2f", invoiceTable.SumGross), "1", 1, "C", false, 0, "")

}

// func getLongestItemNameString(pdf *gofpdf.Fpdf, invoiceTable *InvoiceTable) float64 {
// 	var longestWidth float64
// 	for _, item := range invoiceTable.Items {
// 		width := pdf.GetStringWidth(item.Name)
// 		if width > longestWidth {
// 			longestWidth = width
// 		}
// 	}
// 	return longestWidth
// }
