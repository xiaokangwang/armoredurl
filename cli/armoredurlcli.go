package main

import (
	"flag"
	"fmt"
	"github.com/xiaokangwang/armoredurl"
	"github.com/xiaokangwang/armoredurl/checksum"
	"github.com/xiaokangwang/armoredurl/compression"
	"github.com/xiaokangwang/armoredurl/encoding"
	"github.com/xiaokangwang/armoredurl/prefix"
	"io"
	"os"
	"strings"
)

func main() {
	transformers := flag.String("transformers", "", "comma separated list of transformers")
	decode := flag.Bool("decode", false, "decode input")
	flag.Parse()
	transformerNames := strings.Split(*transformers, ",")
	var transformersInst []armoredurl.Transformer
	for _, transformerName := range transformerNames {
		var transformer armoredurl.Transformer
		switch transformerName {
		case "choco":
			transformer = &compression.ShocoTransformer{}
		case "chocopath":
			transformer = &compression.ShocoPathTransformer{}
		case "gzip":
			transformer = &compression.GzipTransformer{}
		case "base45":
			transformer = &encoding.Base45Transformer{}
		case "base64":
			transformer = &encoding.Base64Transformer{}
		case "crc32":
			transformer = &checksum.CRC32Transformer{}
		default:
			if strings.HasPrefix(transformerName, "prefix:") {
				transformer = &prefix.StaticPrefix{Prefix: transformerName[7:]}
				break
			}
			fmt.Fprintf(os.Stderr, "unknown transformer: %s\n", transformerName)
		}
		transformersInst = append(transformersInst, transformer)
	}
	chainTransformer := armoredurl.ChainTransformer{Transformers: transformersInst}

	stdinData, err := io.ReadAll(os.Stdin)
	if err != nil && err != io.EOF {
		fmt.Fprintf(os.Stderr, "failed to read from stdin: %w", err)
		os.Exit(1)
	}
	inputString := string(stdinData)
	inputString = strings.Trim(inputString, "\n")
	var output string
	if *decode {
		output, err = chainTransformer.Decode(inputString)
		if err != nil {
			fmt.Fprintf(os.Stderr, "failed to encode input: %w", err)
			os.Exit(1)
		}
	} else {
		output, err = chainTransformer.Encode(inputString)
		if err != nil {
			fmt.Fprintf(os.Stderr, "failed to encode input: %w", err)
			os.Exit(1)
		}
	}

	io.Copy(os.Stdout, strings.NewReader(output+"\n"))
}
