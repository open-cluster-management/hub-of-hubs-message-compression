package noop

import "github.com/open-cluster-management/hub-of-hubs-message-compression/compressors"

// NewNoOpCompressor returns a new instance of no-op compressor.
func NewNoOpCompressor() compressors.Compressor {
	return &CompressorNoOp{}
}

// CompressorNoOp implements a compressor based on the No-Op pattern.
type CompressorNoOp struct{}

// GetType returns the string identifier for this no-op compressor in the types map.
func (compressor *CompressorNoOp) GetType() string {
	return "no-op"
}

// Compress returns the bytes received as-is.
func (compressor *CompressorNoOp) Compress(data []byte) ([]byte, error) {
	return data, nil
}

// Decompress returns the bytes received as-is.
func (compressor *CompressorNoOp) Decompress(data []byte) ([]byte, error) {
	return data, nil
}
