package dto

import (
	"strings"

	"github.com/hodl-repos/pdf-invoice/pkg/delimitor"
)

type AddressDto struct {
	Name    *string `json:"name" validate:"required"`
	Street1 *string `json:"street1,omitempty"`
	Street2 *string `json:"street2,omitempty"`
	Zip     *string `json:"zip,omitempty"`
	City    *string `json:"city,omitempty"`
	Country *string `json:"country,omitempty"`
}

func (data *AddressDto) Format(d delimitor.Delimitor) string {
	var sb strings.Builder
	sb.WriteString(*data.Name)

	if data.Street1 != nil {
		sb.WriteString(d.String())
		sb.WriteString(*data.Street1)
	}

	if data.Street2 != nil {
		sb.WriteString(d.String())
		sb.WriteString(*data.Street2)
	}

	if data.Zip != nil || data.City != nil {
		sb.WriteString(d.String())

		if data.Zip != nil && data.City != nil {
			sb.WriteString(*data.Zip)
			sb.WriteString(" ")
			sb.WriteString(*data.City)
		} else if data.Zip != nil {
			sb.WriteString(*data.Zip)
		} else if data.City != nil {
			sb.WriteString(*data.City)
		}
	}

	if data.Country != nil {
		sb.WriteString(d.String())
		sb.WriteString(*data.Country)
	}

	return sb.String()
}
