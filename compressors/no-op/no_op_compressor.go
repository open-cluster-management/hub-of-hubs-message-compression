package noop

import "github.com/stolostron/hub-of-hubs-message-compression/compressors"

const (
	noOpType = "no-op"
)

// NewNoOpCompressor returns a new instance of no-op compressor.
func NewNoOpCompressor() compressors.Compressor {
	return &CompressorNoOp{}
}

// CompressorNoOp implements Compressor based on the No-Op pattern.
type CompressorNoOp struct{}

// GetType returns the string identifier for no-op compressor.
func (compressor *CompressorNoOp) GetType() string {
	return noOpType
}

// Compress returns the bytes received as-is.
func (compressor *CompressorNoOp) Compress(data []byte) ([]byte, error) {
	return data, nil
}

// Decompress returns the bytes received as-is.
func (compressor *CompressorNoOp) Decompress(data []byte) ([]byte, error) {
	return data, nil
}
