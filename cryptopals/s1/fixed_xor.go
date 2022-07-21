package s1

import (
	"encoding/hex"
)

func HexEncodedFixedXor(hexInput1 []byte, hexInput2 []byte) []byte {
	result := make([]byte, len(hexInput1))
	for i := range hexInput1 {
		result[i] = hexInput1[i] ^ hexInput2[i]
	}

	dst := make([]byte, hex.EncodedLen(len(result)))
	hex.Encode(dst, result)

	return dst
}

func FixedXor(block1, block2 []byte) []byte {
	n1 := len(block1)
	n2 := len(block2)

	if n1 != n2 {
		panic("The two input blocks do not have the same length of byte arrays.")
	}

	result := make([]byte, n1)

	for i := 0; i < n1; i++ {
		result[i] = block1[i] ^ block2[i]
	}

	return result
}
