package mathsets

import (
	"crypto/sha256"
)

// Hash256Twice returns double hash of array of bytes.
func Hash256Twice(b []byte) []byte {
	h1 := sha256.Sum256(b)
	h2 := sha256.Sum256(h1[:])
	return h2[:]
}

// Reversebytes reverses the  order of bytes in a byte array.
func Reversebytes(b []byte) []byte {
	d := make([]byte, len(b))
	j := 0
	for i := (len(b) - 1); i >= 0; i-- {
		d[j] = b[i]
		j++
	}
	return d
}
