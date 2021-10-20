package compressors

// Compressor declares the functionality provided by the different compression logics supported.
type Compressor interface {
	Compress([]byte) ([]byte, error)
	Decompress([]byte) ([]byte, error)
}
