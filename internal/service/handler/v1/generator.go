package v1

import (
	"fmt"

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
	pdf.AliasNbPages("{nb}")

	//perpare footer
	footerData := prepareFooterString(data.InvoiceAddress)

	footerLines := pdf.SplitText(footerData, pdf.GetPrintWidth())
	totalFooterTextHeight := float64(len(footerLines)) * pdf.GetFontLineHeight()
	totalFooterBlockHeight := totalFooterTextHeight + 10 + 4.23 + 4.23 + pdf.GetFontLineHeight()
	pdf.SetAutoPageBreak(true, totalFooterBlockHeight)

	pdf.SetFooterFunc(func() {
		if data.Style.ShowMarkerFolding != nil && *data.Style.ShowMarkerFolding {
			if *data.Style.Layout == document.LayoutTypeDIN5008A {
				pdf.Line(0, 87, 10, 87)
				pdf.Line(0, 192, 10, 192)
			}

			if *data.Style.Layout == document.LayoutTypeDIN5008B {
				pdf.Line(0, 105, 10, 105)
				pdf.Line(0, 210, 10, 210)
			}
		}

		if data.Style.ShowMarkerPuncher != nil && *data.Style.ShowMarkerPuncher {
			_, h := pdf.GetPageSize()
			pdf.Line(0, h/2, 14, h/2)
		}

		pdf.SetFont("Arial", "", 8)

		//always display page numbers
		pdf.SetY(-(totalFooterTextHeight + 10 + 4.23 + pdf.GetFontLineHeight()))
		pdf.MCell(pdf.GetPrintWidth(), pdf.GetFontLineHeight(), fmt.Sprintf("Page %d from {nb}", pdf.PageNo()), "1", "R", false)

		//always display footer
		pdf.SetY(-(totalFooterTextHeight + 10))
		pdf.MCell(pdf.GetPrintWidth(), pdf.GetFontLineHeight(), footerData, "1", "M", false)
	})

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
	if data.BankPaymentData != nil && data.Style.ShowBankPaymentQrCode != nil && *data.Style.ShowBankPaymentQrCode {
		if err := generateBankBlock(data, pdf); err != nil {
			return nil, err
		}
	}

	return pdf, nil
}
