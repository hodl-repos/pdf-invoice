package dto

import (
	"strings"

	"github.com/hodl-repos/pdf-invoice/pkg/delimitor"
)

type SellerInformationDto struct {
	Address *AddressDto `json:"address" validate:"required"`

	Email                   *string `json:"email,omitempty" validate:"omitempty,email"`
	Phone                   *string `json:"phone,omitempty"`
	Website                 *string `json:"website,omitempty"`
	VAT                     *string `json:"vat,omitempty"`
	CorporateRegisterNumber *string `json:"corporateRegisterNumber,omitempty"`
}

func (data *SellerInformationDto) Format(d delimitor.Delimitor) string {
	var sb strings.Builder
	sb.WriteString(data.Address.Format(d))

	if data.Email != nil {
		sb.WriteString(d.String())
		sb.WriteString("E-Mail: ")
		sb.WriteString(*data.Email)
	}

	if data.Phone != nil {
		sb.WriteString(d.String())
		sb.WriteString("Phone: ")
		sb.WriteString(*data.Phone)
	}

	if data.Website != nil {
		sb.WriteString(d.String())
		sb.WriteString("Website: ")
		sb.WriteString(*data.Website)
	}

	if data.VAT != nil {
		sb.WriteString(d.String())
		sb.WriteString("UID: ")
		sb.WriteString(*data.VAT)
	}

	if data.CorporateRegisterNumber != nil {
		sb.WriteString(d.String())
		sb.WriteString("Firmenbuch: ")
		sb.WriteString(*data.CorporateRegisterNumber)
	}

	return sb.String()
}
