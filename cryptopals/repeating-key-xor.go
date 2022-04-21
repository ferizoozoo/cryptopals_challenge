package cryptopals

import "encoding/hex"

func RepeatingKeyXorEncryption(plainText string, key string) string {
	plainTextBytes := []byte(plainText)
	plainTextBytesLength := len(plainTextBytes)
	keyLength := len(key)

	cipherText := make([]byte, len(plainTextBytes))

	for i := 0; i < plainTextBytesLength; i++ {
		cipherText[i] = plainTextBytes[i] ^ byte(key[(i)%keyLength])
	}

	return hex.EncodeToString(cipherText)
}
