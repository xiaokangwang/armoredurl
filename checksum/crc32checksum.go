package checksum

import (
	"bytes"
	"encoding/binary"
	"errors"
	"hash/crc32"
	"io"
)

type CRC32Transformer struct {
}

func (t *CRC32Transformer) Encode(input string) (string, error) {
	inputBytes := []byte(input)
	checksum := crc32.ChecksumIEEE(inputBytes)
	writeBuffer := bytes.NewBuffer(nil)
	binary.Write(writeBuffer, binary.BigEndian, checksum)
	io.Copy(writeBuffer, bytes.NewReader(inputBytes))
	return writeBuffer.String(), nil
}

var ErrChecksumMismatch = errors.New("checksum mismatch")

func (t *CRC32Transformer) Decode(input string) (string, error) {
	inputBytes := []byte(input)
	checksum := binary.BigEndian.Uint32(inputBytes[:4])
	decodedBytes := inputBytes[4:]
	if crc32.ChecksumIEEE(decodedBytes) != checksum {
		return "", ErrChecksumMismatch
	}
	return string(decodedBytes), nil
}
