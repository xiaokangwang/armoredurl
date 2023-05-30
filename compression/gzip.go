package compression

import (
	"bytes"
	"compress/gzip"
	"io"
)

type GzipTransformer struct {
}

func (g *GzipTransformer) Encode(input string) (string, error) {
	buf := bytes.NewBuffer(nil)
	w, err := gzip.NewWriterLevel(buf, gzip.BestCompression)
	if err != nil {
		return "", err
	}
	io.Copy(w, bytes.NewReader([]byte(input)))
	w.Close()
	return buf.String(), nil
}

func (g *GzipTransformer) Decode(input string) (string, error) {
	buf := bytes.NewBuffer(nil)
	r, err := gzip.NewReader(bytes.NewReader([]byte(input)))
	if err != nil {
		return "", err
	}
	io.Copy(buf, r)
	return buf.String(), nil
}
