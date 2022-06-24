package s1

func RepeatingKeyXorEncryption(plainText []byte, key []byte) []byte {
	plainTextLength := len(plainText)
	keyLength := len(key)

	cipherText := make([]byte, len(plainText))

	for i := 0; i < plainTextLength; i++ {
		cipherText[i] = plainText[i] ^ key[(i)%keyLength]
	}

	return cipherText
}
