package s2

import (
	"crypto/aes"
	"cryptopals/cryptopals/s1"
)

func CBCModeEncryption(plaintext, iv, key []byte) []byte {
	blocksize := aes.BlockSize
	previous_block := iv
	var ciphertext []byte

	plaintext = PKCS7Padding(plaintext, blocksize)

	for i := 0; i < len(plaintext); i += blocksize {
		block := plaintext[i : i+blocksize]
		xored_result := s1.FixedXor(block, previous_block)

		encrypted_block := s1.EncryptAES128ECBWith(xored_result, key)

		ciphertext = append(ciphertext, encrypted_block...)
		previous_block = encrypted_block
	}

	return ciphertext
}

func CBCModeDecryption(ciphertext, iv, key []byte) []byte {
	blocksize := aes.BlockSize
	previous_block := iv
	var plaintext []byte

	for i := 0; i < len(ciphertext); i += blocksize {
		block := ciphertext[i : i+blocksize]
		decrypted_block := s1.DecryptAES128ECBWith(block, key, blocksize)

		xored_result := s1.FixedXor(decrypted_block, previous_block)

		plaintext = append(plaintext, xored_result...)
		previous_block = block
	}

	return plaintext
}
