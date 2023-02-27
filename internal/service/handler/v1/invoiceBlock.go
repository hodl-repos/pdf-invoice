package v1

import (
	"fmt"

	go2 "github.com/adam-hanna/arrayOperations"
	"github.com/hodl-repos/pdf-invoice/internal/dto"
	"github.com/hodl-repos/pdf-invoice/pkg/document"
	"github.com/hodl-repos/pdf-invoice/pkg/localize"
	"github.com/jung-kurt/gofpdf"
)

func generateInvoiceBlock(data *dto.InvoiceDto, pdf *document.Doc, localizeClient *localize.LocalizeClient) error {
	rawData := prepareInvoiceData(data, localizeClient)

	table, _ := document.NewDocTable(pdf, rawData)
	table.SetAllCellBorders(false)
	table.SetHeadType(document.HeadFirstRow)
	table.SetAllCellPaddings(document.Padding{1, 1, 1, 1})
	table.SetAllCellTypes(document.CellMulti)

	//set every col type to calc, but not the first one
	nrCols := len(rawData[0])
	colTypes := make([]document.ColumnType, 0)
	alTypes := make([]document.AlignmentType, 0)
	colTypes = append(colTypes, document.ColDyn)
	alTypes = append(alTypes, document.AlignLeft)
	for i := 1; i < nrCols; i++ {
		colTypes = append(colTypes, document.ColFixed)
		alTypes = append(alTypes, document.AlignRight)
	}
	table.SetColTypes(colTypes)
	table.SetCellAlingsPerColumn(alTypes)
	table.SetAllColFixedWidths(25.0)

	bg := func(fpdf gofpdf.Fpdf) {
		fpdf.SetFillColor(220, 220, 220)
	}
	table.SetCellStyleFuncsPerAlternateRows(&bg, nil)

	table.Generate()

	return nil
}

func prepareInvoiceData(data *dto.InvoiceDto, localizeClient *localize.LocalizeClient) [][]string {
	tmp := make([][]string, 0)

	showDiscountColumn := go2.Reduce(*data.Rows, func(b bool, ird dto.InvoiceRowDto) bool {
		return ird.DiscountFixed != nil || ird.DiscountPercentage != nil
	}, false)

	headerRow := prepareInvoiceLine(
		data,
		showDiscountColumn,
		localizeClient.TranslateName(),
		localizeClient.TranslateAmount(),
		localizeClient.TranslateNet(),
		localizeClient.TranslateTax(),
		localizeClient.TranslateDiscount(),
		localizeClient.TranslateGross())

	tmp = append(tmp, headerRow)

	for _, row := range *data.Rows {
		amountString := formatAmount(row.Amount, row.AmountUnit, localizeClient)
		netString := formatMoney(row.Net, localizeClient)
		taxString := formatMoney(row.Tax, localizeClient)
		grossString := formatMoney(row.Gross, localizeClient)
		discountString := formatDiscount(row.TaxPercentage, row.Tax, localizeClient)

		titleString := *row.Name
		if row.Description != nil {
			titleString = fmt.Sprintf("%s\n%s", titleString, *row.Description)
		}

		line := prepareInvoiceLine(data, showDiscountColumn, titleString, amountString, netString, taxString, discountString, grossString)
		tmp = append(tmp, line)
	}

	return tmp
}

func formatAmount(value *float64, unit *string, localizeClient *localize.LocalizeClient) string {
	if value == nil || unit == nil {
		return "1"
	}

	return fmt.Sprintf("%v %s", localizeClient.FFloat64(*value), *unit)
}

func formatMoney(value *float64, localizeClient *localize.LocalizeClient) string {
	if value == nil {
		return "-"
	}

	return fmt.Sprintf("%v EUR", localizeClient.FFloat64(*value))
}

func formatDiscount(percentage, fixed *float64, localizeClient *localize.LocalizeClient) string {
	if fixed != nil {
		return formatMoney(fixed, localizeClient)
	}

	if percentage != nil {
		return fmt.Sprintf("%v%", localizeClient.FFloat64(*percentage))
	}

	return "-"
}

func prepareInvoiceLine(style *dto.InvoiceDto, showDiscount bool, title, amount, net, tax, discount, gross string) []string {
	row := make([]string, 0)

	row = append(row, title)
	if style.ShowAmountColumn != nil && *style.ShowAmountColumn {
		row = append(row, amount)
	}
	if style.ShowNetColumn != nil && *style.ShowNetColumn {
		row = append(row, net)
	}
	if style.ShowTaxColumn != nil && *style.ShowTaxColumn {
		row = append(row, tax)
	}
	if showDiscount {
		row = append(row, discount)
	}
	if style.ShowGrossColumn != nil && *style.ShowGrossColumn {
		row = append(row, gross)
	}

	return row
}
