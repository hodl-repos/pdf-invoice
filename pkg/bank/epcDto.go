package bank

import "fmt"

type EpcDto struct {
	ServiceTag       *string
	Version          *string
	CharacterSet     *string
	Identification   *string
	BIC              *string
	Name             *string
	IBAN             *string
	Amount           *string
	Reason           *string
	InvoiceReference *string
	Text             *string
	Information      *string
}

func (dto *EpcDto) GenerateCode() string {
	tmp := fmt.Sprintf("%s\n%s\n%s\n%s", *dto.ServiceTag, *dto.Version, *dto.CharacterSet, *dto.Identification)

	if dto.BIC != nil {
		tmp = fmt.Sprintf("%s\n%s", tmp, *dto.BIC)
	} else {
		tmp = tmp + "\n"
	}

	if dto.Name != nil {
		tmp = fmt.Sprintf("%s\n%s", tmp, *dto.Name)
	} else {
		tmp = tmp + "\n"
	}

	if dto.IBAN != nil {
		tmp = fmt.Sprintf("%s\n%s", tmp, *dto.IBAN)
	} else {
		tmp = tmp + "\n"
	}

	if dto.Amount != nil {
		tmp = fmt.Sprintf("%s\n%s", tmp, *dto.Amount)
	} else {
		tmp = tmp + "\n"
	}

	if dto.Reason != nil {
		tmp = fmt.Sprintf("%s\n%s", tmp, *dto.Reason)
	} else {
		tmp = tmp + "\n"
	}

	if dto.InvoiceReference != nil {
		tmp = fmt.Sprintf("%s\n%s", tmp, *dto.InvoiceReference)
	} else {
		tmp = tmp + "\n"
	}

	if dto.Text != nil {
		tmp = fmt.Sprintf("%s\n%s", tmp, *dto.Text)
	} else {
		tmp = tmp + "\n"
	}

	if dto.Information != nil {
		tmp = fmt.Sprintf("%s\n%s", tmp, *dto.Information)
	} else {
		tmp = tmp + "\n"
	}

	return tmp
}

func (dto *EpcDto) SetDefaults() {
	sTag := "BCD"
	version := "001"
	charSet := "1"
	identify := "SCT"

	dto.ServiceTag = &sTag
	dto.Version = &version
	dto.CharacterSet = &charSet
	dto.Identification = &identify
}

func (dto *EpcDto) SetAmount(amount float64) {
	tmp := fmt.Sprintf("EUR%.2f", amount)
	dto.Amount = &tmp
}
