package compressors

// Compressor declares the functionality provided by the different compression logics supported.
type Compressor interface {
	// GetType returns the string type identifier of the compressor.
	GetType() string
	// Compress compresses a slice of bytes.
	Compress([]byte) ([]byte, error)
	// Decompress decompresses a slice of compressed bytes.
	Decompress([]byte) ([]byte, error)
}
