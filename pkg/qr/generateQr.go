package qr

import "github.com/skip2/go-qrcode"

func GenerateQrCode(qrContent string) (*[]byte, error) {
	qr, err := qrcode.New(qrContent, qrcode.High)

	if err != nil {
		return nil, err
	}

	qr.DisableBorder = true

	png, err := qr.PNG(1000)

	if err != nil {
		return nil, err
	}

	return &png, nil
}
