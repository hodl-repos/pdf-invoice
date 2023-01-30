package dto

type InvoiceRowDto struct {
	Name        *string `json:"name" validate:"required"`
	Description *string `json:"description" validate:"omitempty"`

	Amount     *float64 `json:"amount"`
	AmountUnit *string  `json:"amountUnit"`

	Net           *float64 `json:"net"`
	TaxPercentage *float64 `json:"taxPercentage"`
	Tax           *float64 `json:"tax"`
	Gross         *float64 `json:"gross"`

	DiscountPercentage *float64 `json:"discountPercentage"`
	DiscountFixed      *float64 `json:"discountFixed"`
}
