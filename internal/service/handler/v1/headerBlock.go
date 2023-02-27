package v1

import (
	"net/http"

	go2 "github.com/adam-hanna/arrayOperations"
	"github.com/hodl-repos/pdf-invoice/internal/dto"
	"github.com/hodl-repos/pdf-invoice/pkg/delimitor"
	"github.com/hodl-repos/pdf-invoice/pkg/document"
	"github.com/hodl-repos/pdf-invoice/pkg/localize"
	"github.com/hodl-repos/pdf-invoice/pkg/standardisedError"
	"github.com/jung-kurt/gofpdf"
)

type headerBlockGeneratorsType struct {
	Name     document.LayoutType
	Function func(*dto.DocumentDto, *document.Doc, *localize.LocalizeClient)
}

var (
	headerBlockGenerators = []headerBlockGeneratorsType{
		{
			Name:     document.LayoutTypeDIN5008A,
			Function: createDIN5008ABlock,
		},
		{
			Name:     document.LayoutTypeDIN5008B,
			Function: createDIN5008BBlock,
		},
	}
)

// as this function is called at first - no checks for site-breaks are made
func generateHeaderBlock(data *dto.DocumentDto, pdf *document.Doc, localizeClient *localize.LocalizeClient) error {
	formatFc, ok := go2.FindOne(headerBlockGenerators, func(abgt headerBlockGeneratorsType) bool { return abgt.Name == *data.Style.Layout })

	if !ok {
		return &standardisedError.StandardisedError{
			Type:   "validation-error",
			Title:  "could not find a correct generator for the address-block",
			Status: http.StatusBadRequest,
			Detail: "only din5008a and din5008b are currently supported",
		}
	}

	formatFc.Function(data, pdf, localizeClient)

	return nil
}

func createDIN5008ABlock(data *dto.DocumentDto, pdf *document.Doc, localizeClient *localize.LocalizeClient) {
	pdf.SetXY(25, 27+17.57)

	//TODO: limit to 27.3 max-height
	pdf.MCell(80, pdf.GetFontLineHeight() /* 27.3 */, data.InvoiceAddress.Format(delimitor.NewLine), "", "LT", false)

	lOld, _, rOld, _ := pdf.GetMargins()

	pdf.SetLeftMargin(125)
	pdf.SetRightMargin(10)
	pdf.SetXY(125, 32)
	table, _ := document.NewDocTable(pdf, prepareInformationCells(data.InvoiceInformation, localizeClient))
	table.SetAllCellPaddings(document.Padding{0, 0, 0, 1})
	table.SetAllCellBorders(false)

	table.Generate()

	pdf.SetLeftMargin(lOld)
	pdf.SetRightMargin(rOld)

	if data.Style.Image != nil {
		drawImage(pdf, data.Style.Image, 125, 10, 200, 27)
	}

	//set to content position
	pdf.SetXY(25, 98.5)
}

func createDIN5008BBlock(data *dto.DocumentDto, pdf *document.Doc, localizeClient *localize.LocalizeClient) {
	pdf.SetXY(25, 45+17.7)

	//TODO: limit to 27.3 max-height
	pdf.MCell(80, pdf.GetFontLineHeight() /* 27.3 */, data.InvoiceAddress.Format(delimitor.NewLine), "", "LT", false)

	lOld, _, rOld, _ := pdf.GetMargins()

	pdf.SetLeftMargin(125)
	pdf.SetRightMargin(10)
	pdf.SetXY(125, 50)
	table, _ := document.NewDocTable(pdf, prepareInformationCells(data.InvoiceInformation, localizeClient))
	table.SetAllCellPaddings(document.Padding{0, 0, 0, 1})
	table.SetAllCellBorders(false)

	table.Generate()

	pdf.SetLeftMargin(lOld)
	pdf.SetRightMargin(rOld)

	if data.Style.Image != nil {
		drawImage(pdf, data.Style.Image, 125, 10, 200, 45)
	}
	//set to content position
	pdf.SetXY(25, 98.5)
}

func drawImage(pdf *document.Doc, dto *document.Image, startX, startY, endX, endY float64) error {
	addedImage, err := pdf.AddImage(dto)

	if err != nil {
		return err
	}

	shouldWidth := endX - startX
	shouldHeight := endY - startY

	drawWidth := shouldWidth
	drawHeight := shouldHeight

	ratioArea := drawWidth / drawHeight
	ratioImage := addedImage.Width() / addedImage.Height()

	if ratioArea >= ratioImage {
		//image gets margin left/right
		drawWidth = drawHeight * ratioImage
	} else {
		//image get margin top/bottom
		drawHeight = drawWidth / ratioImage
	}

	drawWidthOffset := (shouldWidth - drawWidth) / 2.0
	drawHeightOffset := (shouldHeight - drawHeight) / 2.0

	pdf.Fpdf.ImageOptions(*dto.ImageUrl,
		startX+drawWidthOffset,
		startY+drawHeightOffset,
		drawWidth,
		drawHeight,
		false,
		gofpdf.ImageOptions{ReadDpi: true},
		0,
		"")

	return nil
}
