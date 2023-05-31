package qr

import (
	"errors"
	"github.com/skip2/go-qrcode"
)

type TwoDimCodeTransformer struct {
}

func (t *TwoDimCodeTransformer) Encode(input string) (string, error) {
	issue, err := qrcode.Encode(input, qrcode.Medium, 250)
	if err != nil {
		return "", err
	}
	return string(issue), nil
}

func (t *TwoDimCodeTransformer) Decode(input string) (string, error) {
	return "", errors.New("not implemented")
}
