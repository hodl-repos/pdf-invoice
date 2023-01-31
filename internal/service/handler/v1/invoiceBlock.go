package v1

import (
	"fmt"

	go2 "github.com/adam-hanna/arrayOperations"
	"github.com/hodl-repos/pdf-invoice/internal/dto"
	"github.com/hodl-repos/pdf-invoice/pkg/document"
)

func generateInvoiceBlock(data *dto.InvoiceDto, pdf *document.Doc) error {
	table, _ := document.NewDocTable(pdf, prepareInvoiceData(data))
	table.SetAllCellPaddings(document.Padding{1, 1, 1, 1})
	table.SetAllCellTypes(document.CellMulti)

	table.Generate()

	return nil
}

func prepareInvoiceData(data *dto.InvoiceDto) [][]string {
	tmp := make([][]string, 0)

	showDiscountColumn := go2.Reduce(*data.Rows, func(b bool, ird dto.InvoiceRowDto) bool {
		return ird.DiscountFixed != nil || ird.DiscountPercentage != nil
	}, false)

	headerRow := prepareInvoiceLine(data, showDiscountColumn, "Title", "Amount", "Net", "Tax", "Discount", "Gross")
	tmp = append(tmp, headerRow)

	for _, row := range *data.Rows {
		amountString := formatAmount(row.Amount, row.AmountUnit)
		netString := formatMoney(row.Net)
		taxString := formatMoney(row.Tax)
		grossString := formatMoney(row.Gross)
		discountString := formatDiscount(row.TaxPercentage, row.Tax)

		titleString := *row.Name
		if row.Description != nil {
			titleString = fmt.Sprintf("%s\n%s", titleString, *row.Description)
		}

		line := prepareInvoiceLine(data, showDiscountColumn, titleString, amountString, netString, taxString, discountString, grossString)
		tmp = append(tmp, line)
	}

	return tmp
}

func formatAmount(value *float64, unit *string) string {
	if value == nil || unit == nil {
		return "1"
	}

	return fmt.Sprintf("%.2f %s", *value, *unit)
}

func formatMoney(value *float64) string {
	if value == nil {
		return "-"
	}

	return fmt.Sprintf("%.2f EUR", *value)
}

func formatDiscount(percentage, fixed *float64) string {
	if fixed != nil {
		return formatMoney(fixed)
	}

	if percentage != nil {
		return fmt.Sprintf("%.2f%", *percentage)
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
