package gzip

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io/ioutil"

	"github.com/open-cluster-management/hub-of-hubs-message-compression/compressors"
)

const (
	gzipCompressorErrorString = "gzip compressor error"
)

// NewGZipCompressor returns a new instance of gzip-based compressor.
func NewGZipCompressor() compressors.Compressor {
	return &CompressorGZip{}
}

// CompressorGZip implements CompressorGZip with gzip-based logic.
type CompressorGZip struct{}

// GetType returns the string representing this GZip compressor in the types map.
func (compressor *CompressorGZip) GetType() string {
	return "gzip"
}

// Compress compresses a slice of bytes using gzip lib.
func (compressor *CompressorGZip) Compress(data []byte) ([]byte, error) {
	var buf bytes.Buffer

	gw := gzip.NewWriter(&buf)
	if _, err := gw.Write(data); err != nil {
		return nil, fmt.Errorf("%s - %w", gzipCompressorErrorString, err)
	}

	if err := gw.Close(); err != nil {
		return nil, fmt.Errorf("%s - %w", gzipCompressorErrorString, err)
	}

	return buf.Bytes(), nil
}

// Decompress decompresses a slice of gzip-compressed bytes using gzip lib.
func (compressor *CompressorGZip) Decompress(compressedData []byte) ([]byte, error) {
	// create a reader for the gzipped data
	gr, err := gzip.NewReader(bytes.NewBuffer(compressedData))
	if err != nil {
		return nil, fmt.Errorf("%s - %w", gzipCompressorErrorString, err)
	}

	data, err := ioutil.ReadAll(gr)
	if err != nil {
		return nil, fmt.Errorf("%s - %w", gzipCompressorErrorString, err)
	}

	_ = gr.Close()

	return data, nil
}
