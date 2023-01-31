package dto

type InvoiceDto struct {
	ShowNetColumn    *bool `json:"showNetColumn"`
	ShowGrossColumn  *bool `json:"showGrossColumn"`
	ShowTaxColumn    *bool `json:"showTaxColumn"`
	ShowAmountColumn *bool `json:"showAmountColumn"`

	ShowNetSum   *bool `json:"showNetSum"`
	ShowTaxSum   *bool `json:"showTaxSum"`
	ShowGrossSum *bool `json:"showGrossSum"`

	SumDiscountPercentage *float64 `json:"sumDiscountPercentage"`
	SumDiscountFixed      *float64 `json:"sumDiscountFixed"`

	Rows *[]InvoiceRowDto `json:"rows"`
}
