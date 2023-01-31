package dto

import "github.com/hodl-repos/pdf-invoice/pkg/document"

type DocumentStyleDto struct {
	LocaleCode   *string `json:"localeCode" validate:"required"`
	LanguageCode *string `json:"languageCode" validate:"required"`

	Layout *document.LayoutType `json:"layout" validate:"required"`

	//only possible when A4-Portrait
	ShowMarkerPuncher *bool `json:"showMarkerPuncher"`

	//only possible when A4-Portrait
	ShowMarkerFolding *bool `json:"showMarkerFolding"`

	//only possible when submitting bank-payment-data in main document
	ShowBankPaymentQrCode *bool `json:"showBankPaymentQrCode"`
}
