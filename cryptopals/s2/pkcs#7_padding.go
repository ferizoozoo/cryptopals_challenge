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

func PKCS7UnPadding(plaintext []byte, blockSize int) []byte {
	plaintextLength := len(plaintext)

	if plaintextLength%blockSize != 0 || blockSize == 0 {
		panic("Plaintext was not padded correctly or blockSize is 0.")
	}

	last_index_of_plaintext := plaintextLength - 1
	last_byte := plaintext[last_index_of_plaintext]

	// Check if plaintext was padded using PKCS7
	for i := last_index_of_plaintext; i > last_index_of_plaintext-int(last_byte); i-- {
		if plaintext[i] != last_byte {
			panic("This plaintext was not padded using PKCS7.")
		}
	}

	// Recreate the unpadded plaintext
	unpadded_plaintext := make([]byte, plaintextLength-int(last_byte))
	for i := 0; i < len(unpadded_plaintext); i++ {
		unpadded_plaintext[i] = plaintext[i]
	}

	return unpadded_plaintext
}
