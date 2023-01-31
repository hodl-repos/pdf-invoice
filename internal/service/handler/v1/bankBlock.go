package v1

import (
	"bytes"
	"math"

	go2 "github.com/adam-hanna/arrayOperations"
	"github.com/hodl-repos/pdf-invoice/internal/dto"
	"github.com/hodl-repos/pdf-invoice/pkg/bank"
	"github.com/hodl-repos/pdf-invoice/pkg/document"
	"github.com/hodl-repos/pdf-invoice/pkg/qr"
	"github.com/jung-kurt/gofpdf"
)

// as this function is called at first - no checks for site-breaks are made
func generateBankBlock(data *dto.DocumentDto, pdf *document.Doc) error {
	bankDto := bank.EpcDto{}
	bankDto.SetDefaults()
	bankDto.InvoiceReference = data.BankPaymentData.PaymentReference
	bankDto.Text = data.BankPaymentData.RemittanceInformation
	bankDto.Name = data.BankPaymentData.AccountHolder
	bankDto.IBAN = data.BankPaymentData.IBAN
	bankDto.BIC = data.BankPaymentData.BIC

	sum := go2.Reduce(*data.InvoiceData.Rows, func(f float64, ird dto.InvoiceRowDto) float64 { return *ird.Gross }, 0.0)
	bankDto.SetAmount(sum)

	l, t, r, _ := pdf.GetMargins()
	currentPosition := pdf.GetY()

	pdf.SetMargins(l+30, t, r)
	pdf.SetXY(l+30, currentPosition+5)
	pdf.MCell(0, pdf.GetFontLineHeight(), prepareBankText(data.BankPaymentData), "", "LM", false)
	newPosition := pdf.GetY()

	spaceY := newPosition - (currentPosition + 5)

	imageSize := math.Min(25, spaceY)

	leftMargin := (30 - imageSize) / 2
	topMargin := (spaceY - imageSize) / 2

	qr, _ := qr.GenerateQrCode(bankDto.GenerateCode())

	pdf.RegisterImageOptionsReader("banktransfer-qr-code", gofpdf.ImageOptions{ImageType: "png", ReadDpi: true}, bytes.NewReader(*qr))
	pdf.ImageOptions("banktransfer-qr-code", l+leftMargin, currentPosition+5+topMargin, imageSize, imageSize, false, gofpdf.ImageOptions{ReadDpi: true}, 0, "")

	pdf.SetMargins(l, t, r)

	pdf.Rect(l, currentPosition, pdf.GetPrintWidth(), spaceY+10, "D")

	pdf.SetXY(newPosition, newPosition+5)

	return nil
}
