package s1

import "crypto/aes"

func EncryptAES128ECBWith(plaintext, key []byte) []byte {
	cipher, _ := aes.NewCipher(key)
	size := aes.BlockSize
	result := make([]byte, len(plaintext))

	for i := 0; i < len(plaintext); i += size {
		cipher.Encrypt(result[i:i+size], plaintext[i:i+size])
	}

	return result
}
