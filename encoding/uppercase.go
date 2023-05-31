package encoding

import "strings"

type UppercaseAll struct {
}

func (u UppercaseAll) Encode(input string) (string, error) {
	return strings.ToUpper(input), nil
}

func (u UppercaseAll) Decode(input string) (string, error) {
	return input, nil
}
