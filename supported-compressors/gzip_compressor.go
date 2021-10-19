package supportedcompressors

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io/ioutil"
)

// CreateGzipCompressor returns a new instance of gzip-based compressor.
func CreateGzipCompressor() *GZipCompressor {
	return &GZipCompressor{}
}

// GZipCompressor implements Compressor with gzip-based logic.
type GZipCompressor struct{}

const gzipCompressorErrorString = "gzip compressor error"

// Compress compresses a slice of bytes using gzip lib.
func (gzc *GZipCompressor) Compress(data []byte) ([]byte, error) {
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
func (gzc *GZipCompressor) Decompress(compressed []byte) ([]byte, error) {
	// Write gzipped data to the client
	gr, err := gzip.NewReader(bytes.NewBuffer(compressed))
	if err != nil {
		return nil, fmt.Errorf("%s - %w", gzipCompressorErrorString, err)
	}

	data, err := ioutil.ReadAll(gr)
	if err != nil {
		return nil, fmt.Errorf("%s - %w", gzipCompressorErrorString, err)
	}

	return data, nil
}
