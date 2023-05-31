package compression

import (
	"github.com/tmthrgd/shoco"
	"github.com/tmthrgd/shoco/models"
)

type ShocoTransformer struct {
}

func (s *ShocoTransformer) Encode(input string) (string, error) {
	return string(shoco.Compress([]byte(input))), nil
}

func (s *ShocoTransformer) Decode(input string) (string, error) {
	decoded, err := shoco.Decompress([]byte(input))
	if err != nil {
		return "", err
	}

	return string(decoded), nil
}

type ShocoPathTransformer struct {
}

func (s *ShocoPathTransformer) Encode(input string) (string, error) {
	return string(models.FilePath().Compress([]byte(input))), nil
}

func (s *ShocoPathTransformer) Decode(input string) (string, error) {
	decoded, err := models.FilePath().Decompress([]byte(input))
	if err != nil {
		return "", err
	}

	return string(decoded), nil
}
