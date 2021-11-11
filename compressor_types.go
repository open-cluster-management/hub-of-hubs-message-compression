package compressor

import (
	noop "github.com/open-cluster-management/hub-of-hubs-message-compression/compressors/default"
	"github.com/open-cluster-management/hub-of-hubs-message-compression/compressors/gzip"
)

// CompressionType is the supported compressing types.
//
// Supported types: Default (no op), GZip.
type CompressionType string

const (
	// Default returns a no-op compressor (does nothing).
	Default CompressionType = "default"
	// GZip is used to create a gzip-based Compressor.
	GZip CompressionType = "gzip"
)

// initializationMap maps compressor-types to instances of the compressors.
var initializationMap = map[CompressionType]func() Compressor{
	Default: func() Compressor {
		return noop.NewNoOpCompressor()
	},
	GZip: func() Compressor {
		return gzip.NewGZipCompressor()
	},
}
