package v1

import (
	"net/http"

	go2 "github.com/adam-hanna/arrayOperations"
	"github.com/hodl-repos/pdf-invoice/internal/dto"
	"github.com/hodl-repos/pdf-invoice/pkg/document"
	"github.com/hodl-repos/pdf-invoice/pkg/standardisedError"
)

// as this function is called at first - no checks for site-breaks are made
func generateBankBlock(data *dto.DocumentDto, pdf *document.Doc) error {
	formatFc, ok := go2.FindOne(headerBlockGenerators, func(abgt headerBlockGeneratorsType) bool { return abgt.Name == *data.Style.Layout })

	if !ok {
		return &standardisedError.StandardisedError{
			Type:   "validation-error",
			Title:  "could not find a correct generator for the address-block",
			Status: http.StatusBadRequest,
			Detail: "only din5008a and din5008b are currently supported",
		}
	}

	formatFc.Function(data, pdf)

	return nil
}
