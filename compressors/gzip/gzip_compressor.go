package gzip

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io/ioutil"

	"github.com/open-cluster-management/hub-of-hubs-message-compression/compressors"
)

const gzipCompressorErrorString = "gzip compressor error"

// NewGzipCompressor returns a new instance of gzip-based compressor.
func NewGzipCompressor() compressors.Compressor {
	return &CompressorGZip{}
}

// CompressorGZip implements Compressor with gzip-based logic.
type CompressorGZip struct{}

// Compress compresses a slice of bytes using gzip lib.
func (gzc *CompressorGZip) Compress(data []byte) ([]byte, error) {
	var buf bytes.Buffer

	gw, err := gzip.NewWriterLevel(&buf, gzip.BestSpeed)
	if err != nil {
		return nil, fmt.Errorf("%s - %w", gzipCompressorErrorString, err)
	}

	if _, err := gw.Write(data); err != nil {
		return nil, fmt.Errorf("%s - %w", gzipCompressorErrorString, err)
	}

	return buf.Bytes(), nil
}

// Decompress decompresses a slice of gzip-compressed bytes using gzip lib.
func (gzc *CompressorGZip) Decompress(compressedData []byte) ([]byte, error) {
	// create a reader for the gzipped data
	gr, err := gzip.NewReader(bytes.NewBuffer(compressedData))
	if err != nil {
		return nil, fmt.Errorf("%s - %w", gzipCompressorErrorString, err)
	}

	data, err := ioutil.ReadAll(gr)
	if err != nil {
		return nil, fmt.Errorf("%s - %w", gzipCompressorErrorString, err)
	}

	return data, nil
}
