package dto

type AddressDto struct {
	Name    *string `json:"name" validate:"required"`
	Street1 *string `json:"street1,omitempty"`
	Street2 *string `json:"street2,omitempty"`
	Zip     *string `json:"zip,omitempty"`
	City    *string `json:"city,omitempty"`
	Country *string `json:"country,omitempty"`
}
