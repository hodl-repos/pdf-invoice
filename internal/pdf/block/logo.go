package block

import (
	"encoding/json"

	"github.com/hodl-repos/pdf-invoice/pkg/pdfhelper"
	"github.com/jung-kurt/gofpdf"
)

// Logo struct represent a logo
type Logo struct {
	Image  *pdfhelper.Image `json:"image,omitempty"`  // Image of the logo
	Width  float64          `json:"width,omitempty"`  // Width of the logo
	Height float64          `json:"height,omitempty"` // Height of the logo
}

// NewLogoFromJSON takes json data represented as []byte and returns a new Logo
// struct if an error occurs while parsing the json, it will return the error.
func NewLogoFromJSON(data []byte) (*Logo, error) {
	var logo Logo
	err := json.Unmarshal(data, &logo)
	if err != nil {
		return nil, err
	}
	logo = setDefaultValues(logo)
	return &logo, nil
}

const LOGO_DEFAULT_PATH = "logo_default.png"
const LOGO_DEFAULT_NAME = "logo_default"

func setDefaultValues(logo Logo) Logo {
	if logo.Image == nil {
		logo.Image = &pdfhelper.Image{
			ImageName:   LOGO_DEFAULT_NAME,
			ImagePath:   LOGO_DEFAULT_PATH,
			ImageString: "",
		}
	}
	return logo
}

// ToJSON returns json string representation of the logo struct
func (l Logo) ToJSON() string {
	jsonBytes, _ := json.Marshal(l)
	return string(jsonBytes)
}

// AddLogoBlock adds a logo at the current position with given width and height.
func AddLogoBlock(pdf *gofpdf.Fpdf, logo *Logo) {
	pdfhelper.AddPNG(pdf, logo.Image)

	x, y := pdf.GetXY()
	pdf.ImageOptions(logo.Image.ImageName, x, y, logo.Width, logo.Height, true, gofpdf.ImageOptions{ReadDpi: true}, 0, "")
}
