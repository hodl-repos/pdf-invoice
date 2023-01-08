package mcoinvoice

import (
	"fmt"

	"github.com/hodl-repos/pdf-invoice/internal/pdf/utils"
	"github.com/jung-kurt/gofpdf"
)

func (mi *mcoInvoice) writeLetterInfo() {
	mi.pdf.Ln(17.8)

	mi.pdf.SetFont("Arial", "", 8)
	addressStr := "medicloud one OG, Thaliastraße 53/15, 1160 Wien"
	mi.pdf.CellFormat(colW1, 2.5, mi.trUTF8(addressStr), borderStr, 1, "", false, 0, "")
	mi.pdf.Ln(9.6)
	mi.pdf.SetFont("Arial", "", 11)
	mi.pdf.CellFormat(colW1, 4.2, mi.trUTF8(mi.Invoice.Customer.Name), borderStr, 1, "", false, 0, "")
	mi.pdf.CellFormat(colW1, 4.2, mi.trUTF8(mi.Invoice.Customer.Street1), borderStr, 1, "", false, 0, "")
	mi.pdf.CellFormat(colW1, 4.2, fmt.Sprintf("%d %s", mi.Invoice.Customer.Zip, mi.trUTF8(mi.Invoice.Customer.City)), borderStr, 1, "", false, 0, "")
	mi.pdf.CellFormat(colW1, 4.2, mi.trUTF8(mi.Invoice.Customer.Country), borderStr, 1, "", false, 0, "")
}

func (mi *mcoInvoice) writeContractInfo() {
	mi.pdf.Ln(32.3)

	mi.pdf.SetFont("Arial", "", 10)
	mi.pdf.CellFormat(colW1, 3.8, "Vertragspartner:", borderStr, 1, "", false, 0, "")
	mi.pdf.Ln(3.6)
	mi.pdf.CellFormat(colW1, 3.8, mi.trUTF8(mi.Invoice.Customer.Name), borderStr, 1, "", false, 0, "")
	mi.pdf.CellFormat(colW1, 3.8, mi.trUTF8(mi.Invoice.Customer.Street1), borderStr, 1, "", false, 0, "")
	mi.pdf.CellFormat(colW1, 3.8, fmt.Sprintf("%d %s", mi.Invoice.Customer.Zip, mi.trUTF8(mi.Invoice.Customer.City)), borderStr, 1, "", false, 0, "")
}

func (mi *mcoInvoice) writeInvoiceInfo() {
	mi.pdf.Ln(4.4)

	mi.pdf.SetFont("Arial", "B", 14)
	mi.pdf.CellFormat(colW2+colW3, 4.6, "Ihre Rechnung", borderStr, 1, "", false, 0, "")
	mi.pdf.Ln(4.4)
	mi.pdf.SetFont("Arial", "B", 10)
	if mi.Invoice.Customer.ID != "" {
		mi.pdf.CellFormat(colW2+colW3, 3.8, "Kundenummer:", borderStr, 0, "", false, 0, "")
		mi.pdf.CellFormat(colW4, 3.8, mi.Invoice.Customer.ID, borderStr, 1, "R", false, 0, "")
	}
	mi.pdf.CellFormat(colW2+colW3, 3.8, "Rechnungsnummer:", borderStr, 0, "", false, 0, "")
	mi.pdf.CellFormat(colW4, 3.8, mi.Invoice.ID, borderStr, 1, "R", false, 0, "")
	mi.pdf.SetFont("Arial", "", 10)
	if mi.Invoice.Customer.UID != "" {
		mi.pdf.CellFormat(colW2+colW3, 3.8, "UID:", borderStr, 0, "", false, 0, "")
		mi.pdf.CellFormat(colW4, 3.8, mi.Invoice.Customer.UID, borderStr, 1, "R", false, 0, "")
	}
	mi.pdf.Ln(9.6)
	mi.pdf.CellFormat(colW2+colW3, 3.8, "Belegdatum:", borderStr, 0, "", false, 0, "")
	mi.pdf.CellFormat(colW4, 3.8, utils.IsoDateFormatter(mi.Invoice.Created, "dd.mm.yyyy"), borderStr, 1, "R", false, 0, "")
	mi.pdf.CellFormat(colW2+colW3, 3.8, mi.trUTF8("Fälligkeit:"), borderStr, 0, "", false, 0, "")
	mi.pdf.CellFormat(colW4, 3.8, utils.IsoDateFormatter(mi.Invoice.DueDate, "dd.mm.yyyy"), borderStr, 1, "R", false, 0, "")
	mi.pdf.CellFormat(colW2+colW3, 3.8, "Abrechnungsperiode:", borderStr, 0, "", false, 0, "")
	mi.pdf.CellFormat(colW4, 3.8, mi.Invoice.Period, borderStr, 1, "R", false, 0, "")
	mi.pdf.Ln(5.2)
	mi.pdf.CellFormat(colW2+colW3, 3.8, "Betrag:", borderStr, 0, "", false, 0, "")
	mi.pdf.CellFormat(colW4, 3.8, utils.CentToString(mi.Invoice.SumWithTax), borderStr, 1, "R", false, 0, "")
	mi.pdf.CellFormat(colW2+colW3, 3.8, "Zahlungsart:", borderStr, 0, "", false, 0, "")
	mi.pdf.CellFormat(colW4, 3.8, mi.trUTF8(mi.Invoice.PaymentType), borderStr, 2, "R", false, 0, "")
	x := mi.pdf.GetX()
	mi.pdf.Ln(5.2)
	mi.pdf.SetX(x)
	mi.pdf.CellFormat(colW4, 3.8, "Kontakt:", borderStr, 2, "R", false, 0, "")
	mi.pdf.CellFormat(colW4, 3.8, "01 / 375 0 222", borderStr, 2, "R", false, 0, "")
	mi.pdf.CellFormat(colW4, 3.8, "office@medicloudone.com", borderStr, 2, "R", false, 0, "")
	mi.pdf.Ln(46.4)
}

func (mi *mcoInvoice) writeSummaryTable() {
	drawVerticalLine(mi.pdf, 0.07)
	mi.pdf.Ln(5.8)

	mi.pdf.SetFont("Arial", "B", 10)
	mi.pdf.CellFormat(colW1, 3.8, "Betrag in EUR exkl. USt.", borderStr, 0, "R", false, 0, "")
	mi.pdf.CellFormat(colW2, 3.8, "USt.", borderStr, 0, "R", false, 0, "")
	mi.pdf.CellFormat(colW3, 3.8, "Betrag USt.", borderStr, 0, "R", false, 0, "")
	mi.pdf.CellFormat(colW4, 3.8, "Betrag in EUR inkl. USt.", borderStr, 1, "R", false, 0, "")
	mi.pdf.Ln(3.6)
	drawVerticalLine(mi.pdf, 0.07)

	// TABLE DATA
	mi.pdf.SetFont("Arial", "", 10)
	for _, x := range mi.Invoice.Items {
		mi.pdf.CellFormat(colW1-15, 6.5, mi.trUTF8(x.Name), borderStr, 0, "", false, 0, "")
		mi.pdf.CellFormat(15, 6.5, utils.CentToString(x.Price), borderStr, 0, "R", false, 0, "")
		mi.pdf.CellFormat(colW2, 6.5, fmt.Sprintf("%d%%", x.TaxRate), borderStr, 0, "R", false, 0, "")
		mi.pdf.CellFormat(colW3, 6.5, utils.CentToString(x.Tax), borderStr, 0, "R", false, 0, "")
		mi.pdf.CellFormat(colW4, 6.5, utils.CentToString(x.PriceWithTax), borderStr, 1, "R", false, 0, "")
	}
	drawVerticalLine(mi.pdf, 0.07)

	// TABLE SUMMARY
	mi.pdf.SetFont("Arial", "B", 10)
	mi.pdf.CellFormat(colW1-15, 5.8, "Gesamtsumme", borderStr, 0, "", false, 0, "")
	mi.pdf.CellFormat(15, 5.8, utils.CentToString(mi.Invoice.Sum), borderStr, 0, "R", false, 0, "")
	mi.pdf.CellFormat(colW2+colW3+colW4, 5.8, utils.CentToString(mi.Invoice.SumWithTax), borderStr, 1, "R", false, 0, "")
	drawVerticalLine(mi.pdf, 0.2)
}

func (mi *mcoInvoice) writePaymentInfo() {
	c := mi.Invoice.Customer
	mi.pdf.Ln(16.8)

	mi.pdf.SetFont("Arial", "B", 9)
	mi.pdf.CellFormat(fullWidth, 3.6, "Unsere Kontodaten: AT68 2027 2000 0072 0979, BIC: SPZWAT21XXX", borderStr, 1, "", false, 0, "")
	if c.BIC != "" && c.IBAN != "" && c.Mandate != "" {
		mi.pdf.CellFormat(fullWidth, 3.6, mi.trUTF8("Die angeführte Rechnungssumme wird von folgendem Konto eingezogen:"), borderStr, 1, "", false, 0, "")
		mi.pdf.SetFont("Arial", "", 9)
		mi.pdf.CellFormat(fullWidth, 3.6, fmt.Sprintf("BIC: %s IBAN: %s Mandat: %s", c.BIC, c.IBAN, c.Mandate), borderStr, 1, "", false, 0, "")
	} else {
		mi.pdf.CellFormat(fullWidth, 3.6, mi.trUTF8("Bei Überweisung führen Sie die Rechnungsnummer als Zahlungsreferenz an."), borderStr, 1, "", false, 0, "")
	}
	mi.pdf.Ln(3.9)
	mi.pdf.SetFont("Arial", "", 9)
	mi.pdf.CellFormat(fullWidth, 3.6, "Bei Lastschrift bitte nicht einzahlen.", borderStr, 1, "", false, 0, "")
	mi.pdf.CellFormat(fullWidth, 3.6, mi.trUTF8("Einwände gegen diese Rechnung sind bis spätestens 3 Monate nach Rechnungserhalt schriftlich möglich, ansonsten gilt sie"), borderStr, 1, "", false, 0, "")
	mi.pdf.CellFormat(fullWidth, 3.6, "als anerkannt.", borderStr, 1, "", false, 0, "")
	mi.pdf.CellFormat(fullWidth, 3.6, fmt.Sprintf("Zahlung: %d Tage netto, 12%% p. a. Verzugszinsen.", mi.Invoice.DueDays), borderStr, 1, "", false, 0, "")
}

func (mi *mcoInvoice) writeFooter() {
	mi.pdf.SetXY(25, 279)
	mi.pdf.SetTextColor(80, 80, 80)
	mi.pdf.SetFont("Arial", "", 7)
	mi.pdf.CellFormat(fullWidth-5, 2.6, mi.trUTF8("medicloud one OG, Thaliastraße 53/15, 1160 Wien"), borderStr, 2, "", false, 0, "")
	mi.pdf.CellFormat(fullWidth-5, 2.6, "Waldviertler Sparkasse Bank AG, IBAN: AT68 2027 2000 0072 0979, BIC: SPZWAT21XXX", borderStr, 2, "", false, 0, "")
	mi.pdf.CellFormat(fullWidth-5, 2.6, "FN 476968z, Handelsgericht Wien, Firmensitz: Wien, UID-Nr. ATU 72559509", borderStr, 2, "", false, 0, "")
}

// *****************************************************************************
// HELPER FUNCTIONS
// *****************************************************************************
func drawVerticalLine(pdf *gofpdf.Fpdf, thickness float64) {
	pdf.SetLineWidth(thickness)

	y := pdf.GetY()
	pdf.Line(colX1, y, colX4+colW4, y)
}
