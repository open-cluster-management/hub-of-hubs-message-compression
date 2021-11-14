package compressors

// Compressor declares the functionality provided by the different compression logics supported.
type Compressor interface {
	GetType() string
	Compress([]byte) ([]byte, error)
	Decompress([]byte) ([]byte, error)
}
