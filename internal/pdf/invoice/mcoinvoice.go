package mcoinvoice

import (
	"bytes"
	"encoding/json"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/hodl-repos/pdf-invoice/internal/pdf/mco"
	"github.com/jung-kurt/gofpdf"
)

var base *gofpdf.Fpdf

const (
	borderStr = "" // "1"...full border, ""... no border

	// margins
	mLeft  = 20
	mTop   = 16
	mRight = 14

	// Column widths
	colW1 = 90
	colW2 = 15
	colW3 = 25
	colW4 = 46
	// Column start x-position
	colX1 = 20
	colX2 = 110
	colX3 = 125
	colX4 = 150

	fullWidth = colW1 + colW2 + colW3 + colW4
)

type mcoInvoice struct {
	Invoice Invoice

	pdf    *gofpdf.Fpdf
	trUTF8 func(string) string
}

func New() *mcoInvoice {
	p := *base
	mi := &mcoInvoice{pdf: &p}
	mi.trUTF8 = base.UnicodeTranslatorFromDescriptor("")
	mi.pdf.SetFooterFunc(mco.FooterFunc(mi.pdf, fullWidth, borderStr))
	return mi
}

func (mi *mcoInvoice) SetParams(input []byte) error {
	return json.Unmarshal(input, &mi.Invoice)
}

func (mi *mcoInvoice) Generate() ([]byte, error) {
	issueDate, err := time.Parse("2006-01-02", mi.Invoice.Created)
	if err != nil {
		return nil, err
	}
	dueDate, err := time.Parse("2006-01-02", mi.Invoice.DueDate)
	if err != nil {
		return nil, err
	}
	mi.Invoice.DueDays = int(dueDate.Sub(issueDate) / 24 / time.Hour)
	mi.pdf.AddPage()

	mco.WriteLogo(mi.pdf)
	mi.pdf.SetCellMargin(5)
	mi.writeLetterInfo()
	mi.writeContractInfo()

	mi.pdf.SetLeftMargin(colX2)
	mi.pdf.SetCellMargin(0)
	mi.pdf.SetY(mTop)
	mi.writeInvoiceInfo()

	mi.pdf.SetLeftMargin(mLeft)
	mi.pdf.Ln(0)
	mi.writeSummaryTable()

	mi.writePaymentInfo()
	// mi.writeFooter()

	var output bytes.Buffer
	err = mi.pdf.Output(&output)
	return output.Bytes(), err
}

func (mi *mcoInvoice) Validate() validator.ValidationErrors { return nil }

func init() {
	// Test asset dependencies
	base = gofpdf.New("Portrait", "mm", "A4", "")
	base.SetMargins(mLeft, mTop, mRight)
	mco.RegisterLogo(base)
}
