package dto

type DocumentDto struct {
	Style *DocumentStyleDto `json:"style" validate:"required"`

	InvoiceAddress *AddressDto `json:"invoiceAddress" validate:"required"`

	InvoiceInformation *InvoiceInformationDto `json:"invoiceInformation" validate:"required"`

	//can be null - no specific customer to be written on the invoice
	CustomerAddress *AddressDto `json:"customerAddress"`

	InvoiceData *InvoiceDto `json:"invoiceData" validate:"required"`

	//string-data-block after the invoice for writing thank you
	InvoiceDataSuffix *string `json:"invoiceDataSuffix" validate:"omitempty"`

	BankPaymentData *BankPaymentDto `json:"bankPaymentData"`
}
