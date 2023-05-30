package prefix

import "strings"

type StaticPrefix struct {
	Prefix string
}

func (s *StaticPrefix) Encode(input string) (string, error) {
	return s.Prefix + input, nil
}

func (s *StaticPrefix) Decode(input string) (string, error) {
	if strings.HasPrefix(input, s.Prefix) {
		return strings.TrimPrefix(input, s.Prefix), nil
	}
	return input, nil
}
