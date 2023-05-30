package encoding

import (
	"encoding/base64"
)

type Base64Transformer struct {
}

func (t *Base64Transformer) Encode(input string) (string, error) {
	return base64.RawURLEncoding.EncodeToString([]byte(input)), nil
}

func (t *Base64Transformer) Decode(input string) (string, error) {
	decoded, err := base64.RawURLEncoding.DecodeString(input)
	if err != nil {
		return "", err
	}

	return string(decoded), nil
}
