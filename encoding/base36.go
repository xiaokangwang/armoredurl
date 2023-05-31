package encoding

import "github.com/martinlindhe/base36"

type Base36Transformer struct {
}

func (b Base36Transformer) Encode(input string) (string, error) {
	return base36.EncodeBytes([]byte(input)), nil
}

func (b Base36Transformer) Decode(input string) (string, error) {
	decoded := base36.DecodeToBytes(input)
	return string(decoded), nil
}
