package v1

import (
	"github.com/hodl-repos/pdf-invoice/internal/dto"
	"github.com/hodl-repos/pdf-invoice/pkg/document"
	"github.com/jung-kurt/gofpdf"
)

func Generate(data *dto.DocumentDto) (*document.Doc, error) {
	defaultsFunction := func(pdf *gofpdf.Fpdf) {
		pdf.SetFont("Arial", "", 10)
		pdf.SetLineWidth(0.2)
		pdf.SetCellMargin(0)

		pdf.SetMargins(25, 25, 20)
		pdf.SetAutoPageBreak(true, 15)
	}

	//create pdf with custom defaults (DIN)
	pdf := document.NewA4WithDefaults(&defaultsFunction)

	//generate invoice header block
	if err := generateHeaderBlock(data, pdf); err != nil {
		return nil, err
	}

	//append customer-address if provided
	if data.CustomerAddress != nil {
		pdf.SetFont("Arial", "B", 12)
		pdf.MCell(0, pdf.GetFontLineHeight(), "Vertragspartner", "", "", false)

		pdf.SetFont("Arial", "", 10)
		pdf.MCell(0, pdf.GetFontLineHeight(), prepareAddressString(data.CustomerAddress), "", "", false)

		pdf.Ln(pdf.GetFontLineHeight())
	}

	//generate invoice-block
	err := generateInvoiceBlock(data.InvoiceData, pdf)
	if err != nil {
		return nil, err
	}

	pdf.Ln(pdf.GetFontLineHeight())

	//append data suffix if provided
	if data.InvoiceDataSuffix != nil {
		pdf.SetFont("Arial", "", 10)
		pdf.MCell(0, pdf.GetFontLineHeight(), *data.InvoiceDataSuffix, "", "", false)

		pdf.Ln(pdf.GetFontLineHeight())
	}

	//generate bank-payment-block
	if data.BankPaymentData != nil {
		if err := generateBankBlock(data, pdf); err != nil {
			return nil, err
		}
	}

	return pdf, nil
}
