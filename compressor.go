package compressor

import (
	"errors"
)

var errCompressorTypeNotFound = errors.New("compressor type not supported")

// NewCompressor returns a compressor instance that corresponds to the given CompressionType.
func NewCompressor(compressorType CompressionType) (Compressor, error) {
	createCompressorFunc, found := initializationMap[compressorType]
	if !found {
		return nil, errCompressorTypeNotFound
	}

	return createCompressorFunc(), nil
}

// Compressor declares the functionality provided by the different compression logics supported.
type Compressor interface {
	GetType() string
	Compress([]byte) ([]byte, error)
	Decompress([]byte) ([]byte, error)
}
