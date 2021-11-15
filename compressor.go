package compressor

import (
	"errors"

	"github.com/open-cluster-management/hub-of-hubs-message-compression/compressors"
	"github.com/open-cluster-management/hub-of-hubs-message-compression/compressors/gzip"
	noop "github.com/open-cluster-management/hub-of-hubs-message-compression/compressors/no-op"
)

var errCompressionTypeNotFound = errors.New("compression type not supported")

// CompressionType is the type identifying supported compression methods.
//
// Supported types: NoOp, GZip.
type CompressionType string

const (
	// NoOp is used to create a no-op Compressor.
	NoOp CompressionType = "no-op"
	// GZip is used to create a gzip-based Compressor.
	GZip CompressionType = "gzip"
)

// NewCompressor returns a compressor instance that corresponds to the given CompressionType.
func NewCompressor(compressionType CompressionType) (compressors.Compressor, error) {
	supportedCompressors := map[CompressionType]func() compressors.Compressor{
		NoOp: noop.NewNoOpCompressor,
		GZip: gzip.NewGZipCompressor,
	}

	createCompressorFunc, found := supportedCompressors[compressionType]
	if !found {
		return nil, errCompressionTypeNotFound
	}

	return createCompressorFunc(), nil
}
