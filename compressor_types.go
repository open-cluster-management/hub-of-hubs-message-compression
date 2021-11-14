package compressor

import (
	"github.com/open-cluster-management/hub-of-hubs-message-compression/compressors"
	noop "github.com/open-cluster-management/hub-of-hubs-message-compression/compressors/default"
	"github.com/open-cluster-management/hub-of-hubs-message-compression/compressors/gzip"
)

// CompressionType is the compressing supported types.
//
// Supported types: Default (no op), GZip.
type CompressionType string

const (
	// Default returns a no-op compressor (does nothing).
	Default CompressionType = "default"
	// GZip is used to create a gzip-based Compressor.
	GZip CompressionType = "gzip"
)

var compressorsMap = map[CompressionType]func() compressors.Compressor{
	Default: noop.NewNoOpCompressor,
	GZip:    gzip.NewGZipCompressor,
}
