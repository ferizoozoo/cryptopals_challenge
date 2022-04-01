package cryptopals

import (
	"encoding/hex"
)

func SingleByteXorCipher(hexStringInput string, key byte) string {
	hexBytes, err := hex.DecodeString(hexStringInput)
	if err != nil {
		panic(err)
	}

	cipherText := make([]byte, len(hexBytes))

	for i := 0; i < len(hexBytes); i++ {
		cipherText[i] = hexBytes[i] ^ key
	}

	return string(cipherText)
}
