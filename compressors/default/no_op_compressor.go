package noop

import "github.com/open-cluster-management/hub-of-hubs-message-compression/compressors"

// NewNoOpCompressor returns a new instance of gzip-based compressor.
func NewNoOpCompressor() compressors.Compressor {
	return &CompressorNoOp{}
}

// CompressorNoOp implements Compressor with gzip-based logic.
type CompressorNoOp struct{}

// GetType returns the string representing this no-op compressor in the types map.
func (compressor *CompressorNoOp) GetType() string {
	return "default"
}

// Compress returns the bytes received as-is.
func (compressor *CompressorNoOp) Compress(data []byte) ([]byte, error) {
	return data, nil
}

// Decompress returns the bytes received as-is.
func (compressor *CompressorNoOp) Decompress(data []byte) ([]byte, error) {
	return data, nil
}
