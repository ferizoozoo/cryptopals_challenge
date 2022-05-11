package cryptopals

import (
	"encoding/hex"
	"math"
	"strings"
)

var MAX_KEYSIZE int = 40

func HammingDistance(x1 []byte, x2 []byte) int {
	if len(x1) != len(x2) {
		return -1
	}

	distance := 0

	for index := 0; index < len(x1); index++ {
		byte1 := x1[index]
		byte2 := x2[index]

		for bit_index := 0; bit_index < 7; bit_index++ {
			mask := byte(1 << bit_index)
			bit1 := byte1 & mask
			bit2 := byte2 & mask

			if bit1 != bit2 {
				distance++
			}
		}
	}

	return distance
}

func findKeySize(message []byte) int {
	minDistance := math.Inf(1)
	rightKeysize := 2

	for keysize := 2; keysize <= 40; keysize++ {
		firstChunk := message[:keysize]
		secondChunk := message[keysize : keysize*2]

		distance := HammingDistance(firstChunk, secondChunk)

		if float64(distance) < minDistance && float64(distance) > 0 {
			minDistance = float64(distance) / float64(keysize)
			rightKeysize = keysize
		}
	}

	return rightKeysize
}

func breakCiphertextIntoTransposedBlocks(ciphertext []byte, keysize int) [][]byte {
	var blocks [][]byte

	for byte_index := 0; byte_index < len(ciphertext)-keysize; byte_index += keysize {
		blocks = append(blocks, []byte(ciphertext)[byte_index:byte_index+keysize])
	}

	var transposedBlocks [][]byte

	for b := 0; b < keysize; b++ {
		var transposedBlock []byte
		for block := 0; block < len(blocks); block++ {
			transposedBlock = append(transposedBlock, blocks[block][b])
		}
		transposedBlocks = append(transposedBlocks, transposedBlock)
	}

	return transposedBlocks
}

func FindEncryptionKey(ciphertext []byte) string {
	var correctKeyOfEachBlock []string
	rightKeySize := findKeySize(ciphertext)
	blocks := breakCiphertextIntoTransposedBlocks(ciphertext, rightKeySize)

	for _, block := range blocks {
		//TODO: bring the code for single byte xor cipher key finder here and change it to use just byte arrays
		hexString := hex.EncodeToString(block)
		correctKey := FindCorrectPlainTextWithSingleByteXorCipher(hexString)
		correctKeyOfEachBlock = append(correctKeyOfEachBlock, correctKey)
	}

	return strings.Join(correctKeyOfEachBlock, "")
}
