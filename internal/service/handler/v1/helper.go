package v1

import (
	"fmt"

	"github.com/hodl-repos/pdf-invoice/internal/dto"
)

func prepareAddressString(data *dto.AddressDto) string {
	//name is required
	tmp := *data.Name

	if data.Street1 != nil {
		tmp = tmp + fmt.Sprintf("\n%s", *data.Street1)
	}

	if data.Street2 != nil {
		tmp = tmp + fmt.Sprintf("\n%s", *data.Street2)
	}

	if data.Zip != nil || data.City != nil {
		if data.Zip == nil {
			tmp = tmp + fmt.Sprintf("\n%s", *data.City)
		} else if data.City == nil {
			tmp = tmp + fmt.Sprintf("\n%s", *data.Zip)
		} else {
			tmp = tmp + fmt.Sprintf("\n%s %s", *data.Zip, *data.City)
		}
	}

	if data.Country != nil {
		tmp = tmp + fmt.Sprintf("\n%s", *data.Country)
	}

	return tmp
}
