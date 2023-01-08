package errorhandling

import "github.com/hodl-repos/pdf-invoice/pkg/standardisedError"

type GetStandardisedErrorInterface interface {
	GetStandardisedError() *standardisedError.StandardisedError
}
