package s2

func PKCS7Padding(plaintext []byte, blockSize int) []byte {
	plaintextLength := len(plaintext)

	if plaintextLength > blockSize || blockSize < 0 {
		panic("Either plaintext is not longer than blockSize or blockSize is less than 0.")
	}

	extraBytesNeeded := blockSize - plaintextLength

	for i := 0; i < extraBytesNeeded; i++ {
		plaintext = append(plaintext, byte(extraBytesNeeded))
	}

	return plaintext
}
