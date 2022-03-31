package cryptopals

import "encoding/hex"

func FixedXor(hexStringInput1 string, hexStringInput2 string) string {
	hexBytes1, err1 := hex.DecodeString(hexStringInput1)
	if err1 != nil {
		panic(err1)
	}

	hexBytes2, err2 := hex.DecodeString(hexStringInput2)
	if err2 != nil {
		panic(err2)
	}

	result := make([]byte, len(hexBytes1))
	for i := range hexBytes1 {
		result[i] = hexBytes1[i] ^ hexBytes2[i]
	}

	return hex.EncodeToString(result)
}
