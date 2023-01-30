package v1

import (
	"bytes"
	"net/http"

	"github.com/hodl-repos/pdf-invoice/internal/dto"
	"github.com/hodl-repos/pdf-invoice/pkg/apihelper"
)

func Handler() apihelper.HandlerFuncWithError {
	return func(w http.ResponseWriter, r *http.Request) error {
		// ctx := r.Context()

		//#region unmarshal
		var request dto.DocumentDto

		err := apihelper.UnmarshalJsonAndValidateWithError(w, r, &request)
		if err != nil {
			return err
		}

		//#endregion unmarshal

		pdf, err := Generate(&request)
		if err != nil {
			return err
		}

		var buf bytes.Buffer
		if err := pdf.Output(&buf); err != nil {
			return err
		}

		w.Header().Set("content-type", "application/pdf")
		w.WriteHeader(http.StatusOK)
		w.Write(buf.Bytes())

		return nil
	}
}
