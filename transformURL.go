package armoredurl

type ChainTransformer struct {
	Transformers []Transformer
}

func (t *ChainTransformer) Encode(input string) (string, error) {
	var err error
	for _, transformer := range t.Transformers {
		input, err = transformer.Encode(input)
		if err != nil {
			return "", err
		}
	}
	return input, nil
}

func (t *ChainTransformer) Decode(input string) (string, error) {
	var err error
	for i := len(t.Transformers) - 1; i >= 0; i-- {
		transformer := t.Transformers[i]
		input, err = transformer.Decode(input)
		if err != nil {
			return "", err
		}
	}
	return input, nil
}

type Transformer interface {
	Encode(input string) (string, error)
	Decode(input string) (string, error)
}
