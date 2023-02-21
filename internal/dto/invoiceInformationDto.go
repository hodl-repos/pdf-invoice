package dto

import "time"

type InvoiceInformationDto struct {
	OfferNumber *string    `json:"offerNumber" validate:"required_without=InvoiceNumber"`
	OfferDate   *time.Time `json:"offerDate" validate:"required_with=OfferNumber"`

	DueDate *time.Time `json:"dueDate" validate:"required"`

	InvoiceNumber *string    `json:"invoiceNumber" validate:"required_without=OfferNumber"`
	InvoiceDate   *time.Time `json:"invoiceDate" validate:"required_with=InvoiceNumber"`

	CustomerIdentifier *string `json:"customerIdentifier"`

	AdditionalInformation *[]AdditionalInvoiceInformationDto `json:"additionalInformation"`
}

type AdditionalInvoiceInformationDto struct {
	Title *string `json:"title" validate:"required"`
	Value *string `json:"value" validate:"required"`
}
