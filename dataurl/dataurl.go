package dataurl

import "github.com/vincent-petithory/dataurl"

type URLEncoder struct {
}

func (u URLEncoder) Encode(input string) (string, error) {
	return dataurl.EncodeBytes([]byte(input)), nil
}

func (u URLEncoder) Decode(input string) (string, error) {
	decoded, err := dataurl.DecodeString(input)
	if err != nil {
		return "", err
	}

	return string(decoded.Data), nil
}
