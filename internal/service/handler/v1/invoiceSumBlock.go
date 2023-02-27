package v1

import (
	"errors"

	"github.com/hodl-repos/pdf-invoice/internal/dto"
	"github.com/hodl-repos/pdf-invoice/pkg/document"
	"github.com/hodl-repos/pdf-invoice/pkg/localize"
)

func generateInvoiceSumBlock(data *dto.InvoiceDto, pdf *document.Doc, localizeClient *localize.LocalizeClient) error {
	rawData, err := prepareInvoiceSumData(data, localizeClient)

	if err != nil {
		return err
	}

	l, t, r, _ := pdf.Fpdf.GetMargins()
	w, _ := pdf.Fpdf.GetPageSize()
	pdf.SetMargins(w/2.0, t, r)
	pdf.SetX(w / 2.0)

	pdf.SetFontStyle("B")

	table, _ := document.NewDocTable(pdf, *rawData)
	table.SetAllCellBorders(false)
	table.SetAllCellPaddings(document.Padding{1, 1, 1, 1})
	table.SetAllCellTypes(document.CellMulti)

	table.SetCellAlingsPerColumn([]document.AlignmentType{document.AlignLeft, document.AlignRight})

	table.Generate()

	pdf.SetMargins(l, t, r)
	pdf.SetX(l)

	return nil
}

func prepareInvoiceSumData(data *dto.InvoiceDto, localizeClient *localize.LocalizeClient) (*[][]string, error) {
	tmp := make([][]string, 0)

	if data.ShowNetSum != nil && *data.ShowNetSum {
		netSum := 0.0

		for _, item := range *data.Rows {
			if item.Net == nil {
				return nil, errors.New("NET NOT DEFINED IN ONE COLUMN")
			}

			netSum += *item.Net
		}

		tmp = append(tmp, []string{
			localizeClient.TranslateNet(),
			formatMoney(&netSum, localizeClient),
		})
	}

	if data.ShowTaxSum != nil && *data.ShowTaxSum {
		taxSum := 0.0

		for _, item := range *data.Rows {
			if item.Tax == nil {
				return nil, errors.New("TAX-VALUE IS NOT DEFINED IN ONE ROW")
			}

			taxSum += *item.Tax
		}

		tmp = append(tmp, []string{
			localizeClient.TranslateTax(),
			formatMoney(&taxSum, localizeClient),
		})
	}

	if data.ShowGrossSum != nil && *data.ShowGrossSum {
		grossSum := 0.0

		for _, item := range *data.Rows {
			if item.Gross == nil {
				return nil, errors.New("GROSS IS NOT DEFINED IN ONE ROW")
			}

			grossSum += *item.Gross
		}

		tmp = append(tmp, []string{
			localizeClient.TranslateGross(),
			formatMoney(&grossSum, localizeClient),
		})
	}

	return &tmp, nil
}
