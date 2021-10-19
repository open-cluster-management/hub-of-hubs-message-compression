package messagecompression

import supportedcompressors "github.com/open-cluster-management/hub-of-hubs-message-compression/supported-compressors"

// CompressorType is the type of supported compressors.
//
// Supported types: GZip.
type CompressorType string

const (
	// GZip is used to create a gzip-based Compressor.
	GZip CompressorType = "gzip"
)

var compressorsMap = map[CompressorType]func() Compressor{
	GZip: func() Compressor {
		return supportedcompressors.CreateGzipCompressor()
	},
}
