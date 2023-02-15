package v1

import (
	"strings"

	"github.com/hodl-repos/pdf-invoice/internal/dto"
)

func prepareFooterString(data *dto.AddressDto) string {
	//name is required
	var sb strings.Builder
	sb.WriteString(*data.Name)

	if data.Street1 != nil {
		sb.WriteString("\n")
		sb.WriteString(*data.Street1)
	}

	if data.Street2 != nil {
		sb.WriteString("\n")
		sb.WriteString(*data.Street2)
	}

	if data.Zip != nil || data.City != nil {
		sb.WriteString("\n")
		if data.Zip == nil {
			sb.WriteString(*data.City)
		} else if data.City == nil {
			sb.WriteString(*data.Zip)
		} else {
			sb.WriteString(*data.Zip)
			sb.WriteString(" ")
			sb.WriteString(*data.City)
		}
	}

	if data.Country != nil {
		sb.WriteString("\n")
		sb.WriteString(*data.Country)
	}

	return sb.String()
}

func prepareAddressString(data *dto.AddressDto) string {
	//name is required
	var sb strings.Builder
	sb.WriteString(*data.Name)

	if data.Street1 != nil {
		sb.WriteString("\n")
		sb.WriteString(*data.Street1)
	}

	if data.Street2 != nil {
		sb.WriteString("\n")
		sb.WriteString(*data.Street2)
	}

	if data.Zip != nil || data.City != nil {
		sb.WriteString("\n")
		if data.Zip == nil {
			sb.WriteString(*data.City)
		} else if data.City == nil {
			sb.WriteString(*data.Zip)
		} else {
			sb.WriteString(*data.Zip)
			sb.WriteString(" ")
			sb.WriteString(*data.City)
		}
	}

	if data.Country != nil {
		sb.WriteString("\n")
		sb.WriteString(*data.Country)
	}

	return sb.String()
}

func ap(table *[][]string, header, content string) {
	columns := make([]string, 2)
	columns[0] = header
	columns[1] = content
	*table = append(*table, columns)
}

// prepared invoice-rows with 2 colums
func prepareInformationCells(data *dto.InvoiceInformationDto) [][]string {
	//name is required
	tmp := make([][]string, 0)

	if data.InvoiceNumber != nil {
		ap(&tmp, "Invoice-Number", *data.InvoiceNumber)
		ap(&tmp, "Invoice-Date", data.InvoiceDate.Format("2006-01-02"))
	} else if data.OfferNumber != nil {
		ap(&tmp, "Offer-Number", *data.OfferNumber)
		ap(&tmp, "Offer-Date", data.OfferDate.Format("2006-01-02"))
	}

	ap(&tmp, "Due-Date", data.DueDate.Format("2006-01-02"))

	if data.AdditionalInformation != nil {
		for _, additional := range *data.AdditionalInformation {
			ap(&tmp, *additional.Title, *additional.Value)
		}
	}

	return tmp
}

func prepareBankText(data *dto.BankPaymentDto) string {
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
		sb.WriteString("Payment-Reference: ")
		sb.WriteString(*data.PaymentReference)
	}

	if data.RemittanceInformation != nil {
		sb.WriteString("\n")
		sb.WriteString("nTransaction-Text: ")
		sb.WriteString(*data.RemittanceInformation)
	}

	return sb.String()
}
