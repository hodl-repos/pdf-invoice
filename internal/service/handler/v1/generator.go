package v1

import (
	"github.com/hodl-repos/pdf-invoice/internal/dto"
	"github.com/hodl-repos/pdf-invoice/pkg/document"
)

func Generate(data *dto.DocumentDto) (*document.Doc, error) {
	pdf := document.NewA4()

	pdf.SetMargins(25, 25, 20)
	pdf.SetAutoPageBreak(true, 15)

	if err := generateHeaderBlock(data, pdf); err != nil {
		return nil, err
	}

	if data.CustomerAddress != nil {
		pdf.SetFont("arial", "B", 10)
		pdf.MCell(0, pdf.GetFontLineHeight(), "Vertragspartner", "", "", false)

		pdf.SetFont("arial", "", 8)
		pdf.MCell(0, pdf.GetFontLineHeight(), prepareAddressString(data.CustomerAddress), "", "", false)
	}

	pdf.Ln(pdf.GetFontLineHeight())

	if err := generateInvoiceBlock(data.InvoiceData, pdf); err != nil {
		return nil, err
	}

	pdf.Ln(pdf.GetFontLineHeight())

	if data.InvoiceDataSuffix != nil {
		pdf.SetFont("arial", "", 8)
		pdf.MCell(0, pdf.GetFontLineHeight(), *data.InvoiceDataSuffix, "", "", false)
		pdf.Ln(pdf.GetFontLineHeight())
	}

	if data.BankPaymentData != nil {
		if err := generateBankBlock(data, pdf); err != nil {
			return nil, err
		}
	}

	return pdf, nil
}
