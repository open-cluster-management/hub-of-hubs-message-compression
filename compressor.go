package compressor

import (
	"errors"

	"github.com/stolostron/hub-of-hubs-message-compression/compressors"
	"github.com/stolostron/hub-of-hubs-message-compression/compressors/gzip"
	noop "github.com/stolostron/hub-of-hubs-message-compression/compressors/no-op"
)

var errCompressionTypeNotFound = errors.New("compression type not supported")

// CompressionType is the type identifying supported compression methods.
type CompressionType string

const (
	// NoOp is used to create a no-op Compressor.
	NoOp CompressionType = "no-op"
	// GZip is used to create a gzip-based Compressor.
	GZip CompressionType = "gzip"
)

// NewCompressor returns a compressor instance that corresponds to the given CompressionType.
func NewCompressor(compressionType CompressionType) (compressors.Compressor, error) {
	switch compressionType {
	case NoOp:
		return noop.NewNoOpCompressor(), nil
	case GZip:
		return gzip.NewGZipCompressor(), nil
	default:
		return nil, errCompressionTypeNotFound
	}
}
