package pdf

import (
	"testing"

	"github.com/hodl-repos/pdf-invoice/internal/pdf/utils"
)

func TestCreatePdf(t *testing.T) {
	pdf := NewA4()

	pdf.SetFooterFunc(utils.FooterFunc(pdf))

	pdf.AddPage()

	pdf.OutputFileAndClose("test.pdf")
}
