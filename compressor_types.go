package compressor

import (
	"github.com/open-cluster-management/hub-of-hubs-message-compression/compressors"
	"github.com/open-cluster-management/hub-of-hubs-message-compression/compressors/gzip"
)

// CompressType is the compressing supported types.
//
// Supported types: GZip.
type CompressType string

const (
	// GZip is used to create a gzip-based Compressor.
	GZip CompressType = "gzip"
)

var compressorsMap = map[CompressType]func() compressors.Compressor{
	GZip: gzip.NewGzipCompressor,
}
