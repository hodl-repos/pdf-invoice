package v1

import (
	"strings"

	"github.com/hodl-repos/pdf-invoice/internal/dto"
	"github.com/hodl-repos/pdf-invoice/pkg/delimitor"
	"github.com/hodl-repos/pdf-invoice/pkg/localize"
)

func prepareFooterString(data *dto.DocumentDto) string {
	if data.Style.FooterOverride != nil {
		return *data.Style.FooterOverride
	}

	return data.SellerInformation.Format(delimitor.Tab)
}

func ap(table *[][]string, header, content string) {
	columns := make([]string, 2)
	columns[0] = header
	columns[1] = content
	*table = append(*table, columns)
}

// prepared invoice-rows with 2 colums
func prepareInformationCells(data *dto.InvoiceInformationDto, localizeClient *localize.LocalizeClient) [][]string {
	//name is required
	tmp := make([][]string, 0)

	if data.InvoiceNumber != nil {
		ap(&tmp, localizeClient.TranslateInvoiceNumber(), *data.InvoiceNumber)
		ap(&tmp, localizeClient.TranslateDate(), data.InvoiceDate.Format("2006-01-02"))
	} else if data.OfferNumber != nil {
		ap(&tmp, localizeClient.TranslateOfferNumber(), *data.OfferNumber)
		ap(&tmp, localizeClient.TranslateDate(), data.OfferDate.Format("2006-01-02"))
	}

	ap(&tmp, localizeClient.TranslateDueDate(), data.DueDate.Format("2006-01-02"))

	if data.AdditionalInformation != nil {
		for _, additional := range *data.AdditionalInformation {
			ap(&tmp, *additional.Title, *additional.Value)
		}
	}

	return tmp
}

func prepareBankText(data *dto.BankPaymentDto, localizeClient *localize.LocalizeClient) string {
	// name is required
	var sb strings.Builder
	sb.WriteString(*data.AccountHolder)
	sb.WriteString("\n")
	sb.WriteString(*data.BankName)

	if data.IBAN != nil {
		sb.WriteString("\n")
		sb.WriteString(*data.IBAN)
	}

	if data.BIC != nil {
		sb.WriteString("\n")
		sb.WriteString(*data.BIC)
	}

	if data.PaymentReference != nil {
		sb.WriteString("\n")
		sb.WriteString(localizeClient.TranslatePaymentReference())
		sb.WriteString(": ")
		sb.WriteString(*data.PaymentReference)
	}

	if data.RemittanceInformation != nil {
		sb.WriteString("\n")
		sb.WriteString(localizeClient.TranslateRemittanceInformation())
		sb.WriteString(": ")
		sb.WriteString(*data.RemittanceInformation)
	}

	return sb.String()
}
