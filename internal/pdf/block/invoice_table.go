package block

import (
	"encoding/json"
	"fmt"

	"github.com/hodl-repos/pdf-invoice/pkg/pdfhelper"
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

const (
	b                = "1" // border string
	columnGap        = 5.0
	cellMargin       = 1 // top and bottom margin of the table cell
	headerHeight     = 8.
	headerFontSize   = 10.
	itemFontSize     = 10
	lineHeightFactor = 1.1
)

// AddInvoiceTableBlock  create a table in pdf using information from an invoice
// table struct.
func AddInvoiceTableBlock(doc *pdfhelper.Doc, invoiceTable *InvoiceTable) {
	lang := doc.Lang
	pdfhelper.DrawVerticalLinePrintWidth(doc.Fpdf)
	doc.Ln(2.5)

	// table header ----------------
	doc.SetFont("Arial", "B", headerFontSize)
	h2 := pdfhelper.T("AMOUNT_IN_EUR_EXCL_VAT", lang)
	h3 := pdfhelper.T("VAT", lang)
	h4 := pdfhelper.T("AMOUNT_VAT", lang)
	h5 := pdfhelper.T("AMOUNT_IN_EUR_INCL_VAT", lang)

	w2 := doc.GetStringWidth(h2)
	w3 := doc.GetStringWidth(h3) + columnGap
	w4 := doc.GetStringWidth(h4) + columnGap
	w5 := doc.GetStringWidth(h5) + columnGap
	w1 := pdfhelper.GetPrintWidth(doc.Fpdf) - w2 - w3 - w4 - w5

	doc.CFormat(w1, headerHeight, "", b, 0, "L", false, 0, "")
	doc.CFormat(w2, headerHeight, h2, b, 0, "R", false, 0, "")
	doc.CFormat(w3, headerHeight, h3, b, 0, "R", false, 0, "")
	doc.CFormat(w4, headerHeight, h4, b, 0, "R", false, 0, "")
	doc.CFormat(w5, headerHeight, h5, b, 1, "R", false, 0, "")
	pdfhelper.DrawVerticalLinePrintWidth(doc.Fpdf)
	doc.Ln(1.5)

	doc.SetFont("Arial", "", itemFontSize) // reset font

	// table items ----------------
	// Loop through each item in the invoice table
	for _, item := range invoiceTable.Items {
		// Add the item information to the table

		lines := doc.SplitText(item.Name, w1)
		_, fontHt := doc.GetFontSize()
		lineHt := fontHt * lineHeightFactor
		cellHt := float64(len(lines))*lineHt + 2*cellMargin

		x, y := doc.GetXY()
		doc.Rect(x, y, w1, cellHt, "D")

		doc.SetY(y + cellMargin)
		doc.MCell(w1, lineHt, item.Name, "", "L", false)
		// y2 := doc.GetY()   // get y after MultiCell
		// h := y2 - y        // calculate MultiCell Height
		// doc.SetXY(x+w1, y) // set cursor to continue after MultiCell

		h := cellHt
		doc.SetXY(x+w1, y)

		doc.CFormat(w2, h, fmt.Sprintf("%.2f", item.Net), b, 0, "TR", false, 0, "")
		doc.CFormat(w3, h, fmt.Sprintf("%.2f", item.Tax_Rate), b, 0, "TR", false, 0, "")
		doc.CFormat(w4, h, fmt.Sprintf("%.2f", item.Tax), b, 0, "TR", false, 0, "")
		doc.CFormat(w5, h, fmt.Sprintf("%.2f", item.Gross), b, 1, "TR", false, 0, "")

		doc.SetY(y + cellHt) // set the cursor on the correct height after MultiCell
	}
	// Add the sum net and gross to the table
	// doc.CFormat(40, 10, "Sum Net", "1", 0, "C", false, 0, "")
	// doc.CFormat(60, 10, fmt.Sprintf("%.2f", invoiceTable.SumNet), "1", 0, "C", false, 0, "")
	// doc.CFormat(20, 10, "Sum Gross", "1", 0, "C", false, 0, "")
	// doc.CFormat(20, 10, fmt.Sprintf("%.2f", invoiceTable.SumGross), "1", 1, "C", false, 0, "")
}

func tableRow(row []string, cols []float64)

func tableCell(doc *pdfhelper.Doc, w float64, txtStr, borderStr, alignStr string) float64 {
	lines := doc.SplitText(txtStr, w)
	_, fontHt := doc.GetFontSize()
	lineHt := fontHt * lineHeightFactor
	cellHt := float64(len(lines))*lineHt + 2*cellMargin

	x, y := doc.GetXY()
	if borderStr == "1" {
		doc.Rect(x, y, w, cellHt, "D")
	}

	doc.SetY(y + cellMargin)
	doc.MCell(w, lineHt, txtStr, "", "L", false)

	doc.SetXY(x+w, y)
	return cellHt
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
