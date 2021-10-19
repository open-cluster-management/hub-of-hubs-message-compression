package messagecompression

import (
	"errors"
)

var errCompressorTypeNotFound = errors.New("compressor type not supported")

// NewCompressor returns a compressor instance that corresponds to the given CompressorType.
func NewCompressor(compressorType CompressorType) (Compressor, error) {
	createCompressorFunc, found := compressorsMap[compressorType]
	if !found {
		return nil, errCompressorTypeNotFound
	}

	return createCompressorFunc(), nil
}

// Compressor declares the functionality provided by the different compression logics supported.
type Compressor interface {
	Compress([]byte) ([]byte, error)
	Decompress([]byte) ([]byte, error)
}
