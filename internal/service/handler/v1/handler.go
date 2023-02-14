package v1

import (
	"bytes"
	"net/http"

	"github.com/hodl-repos/pdf-invoice/internal/dto"
	"github.com/hodl-repos/pdf-invoice/pkg/apihelper"
	"github.com/hodl-repos/pdf-invoice/pkg/logging"
)

func Handler() apihelper.HandlerFuncWithError {
	return func(w http.ResponseWriter, r *http.Request) error {
		ctx := r.Context()
		logger := logging.FromContext(ctx)

		logger.Debugln("got request for v1-generate")

		logger.Debugln("deserializing dto")

		//#region unmarshal
		var request dto.DocumentDto

		err := apihelper.UnmarshalJsonAndValidateWithError(w, r, &request)
		if err != nil {
			return err
		}

		logger.Debugln("generating pdf")

		//#endregion unmarshal

		pdf, err := Generate(&request)
		if err != nil {
			return err
		}

		logger.Debugln("serializing pdf")

		var buf bytes.Buffer
		if err := pdf.Output(&buf); err != nil {
			return err
		}

		logger.Debugln("sending response")

		w.Header().Set("content-type", "application/pdf")
		w.WriteHeader(http.StatusOK)
		w.Write(buf.Bytes())

		return nil
	}
}
