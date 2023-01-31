package v1

import (
	"fmt"

	"github.com/hodl-repos/pdf-invoice/internal/dto"
)

func prepareAddressString(data *dto.AddressDto) string {
	//name is required
	tmp := *data.Name

	if data.Street1 != nil {
		tmp = tmp + fmt.Sprintf("\n%s", *data.Street1)
	}

	if data.Street2 != nil {
		tmp = tmp + fmt.Sprintf("\n%s", *data.Street2)
	}

	if data.Zip != nil || data.City != nil {
		if data.Zip == nil {
			tmp = tmp + fmt.Sprintf("\n%s", *data.City)
		} else if data.City == nil {
			tmp = tmp + fmt.Sprintf("\n%s", *data.Zip)
		} else {
			tmp = tmp + fmt.Sprintf("\n%s %s", *data.Zip, *data.City)
		}
	}

	if data.Country != nil {
		tmp = tmp + fmt.Sprintf("\n%s", *data.Country)
	}

	return tmp
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
	tmp := *data.AccountHolder + "\n" + *data.BankName

	if data.IBAN != nil {
		tmp = tmp + fmt.Sprintf("\nIBAN: %s", *data.IBAN)
	}

	if data.BIC != nil {
		tmp = tmp + fmt.Sprintf("\nBIC: %s", *data.BIC)
	}

	if data.PaymentReference != nil {
		tmp = tmp + fmt.Sprintf("\nPayment-Reference: %s", *data.PaymentReference)
	}

	if data.RemittanceInformation != nil {
		tmp = tmp + fmt.Sprintf("\nTransaction-Text: %s", *data.RemittanceInformation)
	}

	return tmp
}
