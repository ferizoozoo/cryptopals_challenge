package s1

import (
	"encoding/hex"
)

func FixedXor(hexInput1 []byte, hexInput2 []byte) []byte {
	result := make([]byte, len(hexInput1))
	for i := range hexInput1 {
		result[i] = hexInput1[i] ^ hexInput2[i]
	}

	dst := make([]byte, hex.EncodedLen(len(result)))
	hex.Encode(dst, result)

	return dst
}
