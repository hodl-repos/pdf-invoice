package document

import (
	"bytes"
	"errors"
	"image/jpeg"
	"image/png"
	"io/ioutil"
	"net/http"

	"github.com/jung-kurt/gofpdf"
)

type Image struct {
	ImageUrl *string `json:"imageUrl" validate:"required"`
}

func (doc *Doc) AddImage(dto *Image) (*gofpdf.ImageInfoType, error) {
	rawImg, err := getImageFromUrl(dto.ImageUrl)

	if err != nil {
		return nil, err
	}

	//TEST IF PNG OR JPG
	if _, err = png.Decode(bytes.NewReader(*rawImg)); err == nil {
		return doc.Fpdf.RegisterImageOptionsReader(*dto.ImageUrl,
			gofpdf.ImageOptions{ImageType: "PNG", ReadDpi: true},
			bytes.NewReader(*rawImg)), nil
	}

	if _, err = jpeg.Decode(bytes.NewReader(*rawImg)); err == nil {
		return doc.Fpdf.RegisterImageOptionsReader(*dto.ImageUrl,
			gofpdf.ImageOptions{ImageType: "JPG", ReadDpi: true},
			bytes.NewReader(*rawImg)), nil
	}

	return nil, errors.New("cannot load image - wrong format?")
}

func getImageFromUrl(url *string) (*[]byte, error) {
	// Create a new request using http
	req, err := http.NewRequest("GET", *url, nil)

	if err != nil {
		return nil, err
	}

	// Send req using http Client
	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("wront status @ get request: " + resp.Status)
	}

	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	return &data, nil
}
