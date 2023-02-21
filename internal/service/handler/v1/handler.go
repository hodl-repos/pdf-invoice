package v1

import (
	"bytes"
	"net/http"

	"github.com/hodl-repos/pdf-invoice/internal/dto"
	"github.com/hodl-repos/pdf-invoice/pkg/apihelper"
	"github.com/hodl-repos/pdf-invoice/pkg/localize"
	"github.com/hodl-repos/pdf-invoice/pkg/logging"
)

func Handler(localizationProvider *localize.LocalizeService) apihelper.HandlerFuncWithError {
	return func(w http.ResponseWriter, r *http.Request) error {
		ctx := r.Context()
		logger := logging.FromContext(ctx)

		logger.Debugln("got request for v1-generate")

		//#region unmarshal
		logger.Debugln("deserializing dto")

		var request dto.DocumentDto

		err := apihelper.UnmarshalJsonAndValidateWithError(w, r, &request)
		if err != nil {
			return err
		}

		//#endregion unmarshal

		//starting localization

		localizationClient := localizationProvider.CreateClient(*request.Style.LocaleCode, *request.Style.LanguageCode)

		logger.Debugln("generating pdf")

		pdf, err := Generate(&request, localizationClient)
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
