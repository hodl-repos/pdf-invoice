package mcoinvoice

type Customer struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Street1 string `json:"street1"`
	Zip     int    `json:"zip"`
	City    string `json:"city"`
	Country string `json:"country"`
	UID     string `json:"uid"`
	BIC     string `json:"bic"`
	IBAN    string `json:"iban"`
	Mandate string `json:"mandate"`
}

type Invoice struct {
	ID          string        `json:"id"`
	Created     string        `json:"created"`
	DueDays     int           `json:"-"`
	DueDate     string        `json:"due_date"`
	Period      string        `json:"period"`
	Sum         int           `json:"sum"`
	SumWithTax  int           `json:"sum_with_tax"`
	PaymentType string        `json:"payment_type"`
	Items       []InvoiceItem `json:"items"`
	Customer    Customer      `json:"customer"`
}

type InvoiceItem struct {
	Name         string `json:"name"`
	Price        int    `json:"price"`
	TaxRate      int    `json:"tax_rate"`
	Tax          int    `json:"tax"`
	PriceWithTax int    `json:"price_with_tax"`
}
