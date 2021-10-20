package compressor

import (
	"errors"

	"github.com/open-cluster-management/hub-of-hubs-message-compression/compressors"
)

var errCompressorTypeNotFound = errors.New("compressor type not supported")

// NewCompressor returns a compressor instance that corresponds to the given CompressorType.
func NewCompressor(compressorType CompressType) (compressors.Compressor, error) {
	createCompressorFunc, found := compressorsMap[compressorType]
	if !found {
		return nil, errCompressorTypeNotFound
	}

	return createCompressorFunc(), nil
}
