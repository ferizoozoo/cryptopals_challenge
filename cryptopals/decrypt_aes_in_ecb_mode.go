package cryptopals

import (
	"crypto/aes"
)

func DecryptAES128ECBWith(input, key []byte, size int) []byte {
	cipher, _ := aes.NewCipher(key)
	result := make([]byte, len(input))

	for i := 0; i < len(input); i += size {
		cipher.Decrypt(result[i:i+size], input[i:i+size])
	}

	return result
}
