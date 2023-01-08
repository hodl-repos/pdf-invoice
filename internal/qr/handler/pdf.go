package handler

import (
	"bytes"
	"net/http"
	"strconv"

	"github.com/hodl-repos/pdf-invoice/internal/pdf"
	"github.com/hodl-repos/pdf-invoice/pkg/jsonutil"
	"github.com/hodl-repos/pdf-invoice/pkg/logging"
)

// v1/qr?content=lkjljk&size=mm
func HandlePdfGetV1() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		logger := logging.FromContext(ctx).Named("qr: handle-pdf-get")

		logger.Debug("qr-pdf: get invoked")

		q := r.URL.Query()

		logger.Debug("qr-pdf: parsing query params")

		content := q.Get("content")

		if len(content) < 1 {
			jsonutil.MarshalResponse(w, http.StatusBadRequest, "qr-pdf: cannot read content of query")
			return
		}

		size, err := strconv.ParseFloat(q.Get("size"), 64)

		if err != nil {
			jsonutil.MarshalResponse(w, http.StatusBadRequest, "qr-pdf: cannot read size of query")
			return
		}

		logger.Debug("qr-pdf: gen pdf")

		pdfFile := pdf.New(size, size)

		// err = pdf.AddQr(content, pdfFile)

		if err != nil {
			jsonutil.MarshalResponse(w, http.StatusBadRequest, "qr-pdf: cannot read size of query")
			return
		}

		var buf bytes.Buffer
		err = pdfFile.Output(&buf)

		if err != nil {
			jsonutil.MarshalResponse(w, http.StatusBadRequest, "qr-pdf: cannot serialize pdf")
			return
		}

		w.Header().Set("content-type", "application/pdf")
		w.WriteHeader(http.StatusOK)
		w.Write(buf.Bytes())
	}
}
