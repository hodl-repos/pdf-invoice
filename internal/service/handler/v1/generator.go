package v1

import (
	"github.com/hodl-repos/pdf-invoice/internal/dto"
	"github.com/hodl-repos/pdf-invoice/pkg/document"
)

func Generate(data *dto.DocumentDto) (*document.Doc, error) {
	pdf := document.NewA4()

	if err := generateHeaderBlock(data, pdf); err != nil {
		return nil, err
	}

	return pdf, nil
}
