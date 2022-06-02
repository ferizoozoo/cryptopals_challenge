package cryptopals

import (
	"crypto/aes"
)

func DetectAES128ECB(inputs [][]byte) []byte {
	key := []byte("YELLOW SUBMARINE")
	size := len(key)

	cipher, _ := aes.NewCipher(key)
	var foundInput []byte

	for _, input := range inputs {
		blockRepeatCount := make(map[string]int, len(input)/size)
		decrypted_result := make([]byte, len(input))

		for i := 0; i < len(input); i += size {
			cipher.Decrypt(decrypted_result[i:i+size], input[i:i+size])
			blockRepeatCount[string(decrypted_result[i:i+size])]++
		}

		isBlockFound := false
		for _, v := range blockRepeatCount {
			if v > 1 {
				foundInput = input
				isBlockFound = true
				break
			}
		}

		if isBlockFound {
			break
		}
	}
	return foundInput
}
