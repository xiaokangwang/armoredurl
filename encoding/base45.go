package encoding

import "github.com/dasio/base45"

type Base45Transformer struct {
}

func (t *Base45Transformer) Encode(input string) (string, error) {
	return base45.EncodeToString([]byte(input)), nil
}

func (t *Base45Transformer) Decode(input string) (string, error) {
	decoded, err := base45.DecodeString(input)
	if err != nil {
		return "", err
	}

	return string(decoded), nil
}
