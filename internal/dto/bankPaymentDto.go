package dto

type BankPaymentDto struct {
	AccountHolder *string `json:"accountHolder" validate:"required"`
	BankName      *string `json:"bankName" validate:"required"`
	IBAN          *string `json:"iban" validate:"required"`
	BIC           *string `json:"bic,omitempty"`
}
