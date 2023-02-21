package dto

import (
	"strings"

	"github.com/hodl-repos/pdf-invoice/pkg/delimitor"
)

type InvoiceAddressDto struct {
	*AddressDto

	VAT *string `json:"vat,omitempty"`
}

func (data *InvoiceAddressDto) Format(d delimitor.Delimitor) string {
	var sb strings.Builder
	sb.WriteString(data.AddressDto.Format(d))

	if data.VAT != nil {
		sb.WriteString(d.String())
		sb.WriteString(*data.VAT)
	}

	return sb.String()
}
