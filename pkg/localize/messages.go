package localize

import "github.com/nicksnyder/go-i18n/v2/i18n"

// Messages
var (
	Messages []i18n.Message = []i18n.Message{
		{
			ID:    "PageNumberWithTotalCount",
			Other: "Page {{.PageNumber}} from {{.PageCount}}",
		},
		{
			ID:    "InvoiceNumber",
			Other: "Invoice no.",
		},
		{
			ID:    "Date",
			Other: "Date",
		},
		{
			ID:    "OfferNumber",
			Other: "Offer no.",
		},
		{
			ID:    "DueDate",
			Other: "Due date",
		},
		{
			ID:    "CustomerIdentifier",
			Other: "Customer Number",
		},
		{
			ID:    "Name",
			Other: "Name",
		},
		{
			ID:    "Amount",
			Other: "Amount",
		},
		{
			ID:    "Net",
			Other: "Net",
		},
		{
			ID:    "Tax",
			Other: "Tax",
		},
		{
			ID:    "Discount",
			Other: "Discount",
		},
		{
			ID:    "Gross",
			Other: "Gross",
		},
		{
			ID:    "PaymentReference",
			Other: "Payment-Reference",
		},
		{
			ID:    "RemittanceInformation",
			Other: "Transaction-Text",
		},
		{
			ID:    "ContractingParty",
			Other: "Contracting Party",
		},
	}
)

func (client *LocalizeClient) TranslatePageNumberWithTotalCount(pageNumber int, pageCount string) string {
	return client.localizer.MustLocalize(&i18n.LocalizeConfig{
		MessageID: "PageNumberWithTotalCount",
		TemplateData: map[string]interface{}{
			"PageNumber": pageNumber,
			"PageCount":  pageCount,
		},
	})
}

func (client *LocalizeClient) TranslateInvoiceNumber() string {
	return client.localizer.MustLocalize(&i18n.LocalizeConfig{
		MessageID: "InvoiceNumber",
	})
}

func (client *LocalizeClient) TranslateDate() string {
	return client.localizer.MustLocalize(&i18n.LocalizeConfig{
		MessageID: "Date",
	})
}

func (client *LocalizeClient) TranslateOfferNumber() string {
	return client.localizer.MustLocalize(&i18n.LocalizeConfig{
		MessageID: "OfferNumber",
	})
}

func (client *LocalizeClient) TranslateDueDate() string {
	return client.localizer.MustLocalize(&i18n.LocalizeConfig{
		MessageID: "DueDate",
	})
}

func (client *LocalizeClient) TranslateCustomerIdentifier() string {
	return client.localizer.MustLocalize(&i18n.LocalizeConfig{
		MessageID: "CustomerIdentifier",
	})
}

func (client *LocalizeClient) TranslateName() string {
	return client.localizer.MustLocalize(&i18n.LocalizeConfig{
		MessageID: "Name",
	})
}

func (client *LocalizeClient) TranslateAmount() string {
	return client.localizer.MustLocalize(&i18n.LocalizeConfig{
		MessageID: "Amount",
	})
}

func (client *LocalizeClient) TranslateNet() string {
	return client.localizer.MustLocalize(&i18n.LocalizeConfig{
		MessageID: "Net",
	})
}

func (client *LocalizeClient) TranslateTax() string {
	return client.localizer.MustLocalize(&i18n.LocalizeConfig{
		MessageID: "Tax",
	})
}

func (client *LocalizeClient) TranslateDiscount() string {
	return client.localizer.MustLocalize(&i18n.LocalizeConfig{
		MessageID: "Discount",
	})
}

func (client *LocalizeClient) TranslateGross() string {
	return client.localizer.MustLocalize(&i18n.LocalizeConfig{
		MessageID: "Gross",
	})
}

func (client *LocalizeClient) TranslatePaymentReference() string {
	return client.localizer.MustLocalize(&i18n.LocalizeConfig{
		MessageID: "PaymentReference",
	})
}

func (client *LocalizeClient) TranslateRemittanceInformation() string {
	return client.localizer.MustLocalize(&i18n.LocalizeConfig{
		MessageID: "RemittanceInformation",
	})
}

func (client *LocalizeClient) TranslateContractingParty() string {
	return client.localizer.MustLocalize(&i18n.LocalizeConfig{
		MessageID: "ContractingParty",
	})
}
