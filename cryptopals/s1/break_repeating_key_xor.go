package s1

import (
	"math"
	"strings"
)

var MAX_KEYSIZE int = 40

//correct! w
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

//kinda correct!!
func findKeySize(message []byte) int {
	minDistance := math.Inf(1)
	rightKeysize := 2

	for keysize := 2; keysize <= 40; keysize++ {
		distance := 0.0

		for i := 0; i < (len(message)/keysize)-1; i++ {
			firstChunk := message[i*keysize : (i+1)*keysize]
			secondChunk := message[(i+1)*keysize : (i+2)*keysize]

			if len(secondChunk) < keysize {
				break
			}

			distance += float64(HammingDistance(firstChunk, secondChunk))
		}

		distance /= float64(len(message))

		if distance < minDistance && distance >= 0 {
			minDistance = distance
			rightKeysize = keysize
		}
	}

	return rightKeysize
}

//maybe something here is broken!!
func breakCiphertextIntoTransposedBlocks(ciphertext []byte, keysize int) [][]byte {
	var blocks [][]byte

	for byte_index := 0; byte_index < len(ciphertext); byte_index += keysize {
		if byte_index+keysize > len(ciphertext) {
			blocks = append(blocks, ciphertext[byte_index:])
		} else {
			blocks = append(blocks, ciphertext[byte_index:byte_index+keysize])
		}
	}

	transposedBlocks := make([][]byte, len(blocks[0]))
	for _, block := range blocks {
		for i, b := range block {
			transposedBlocks[i] = append(transposedBlocks[i], b)
		}
	}

	return transposedBlocks
}

//or here!!
func FindPlainText(ciphertext []byte) string {
	var correctPlainTextOfEachBlock []string
	rightKeySize := findKeySize(ciphertext)
	blocks := breakCiphertextIntoTransposedBlocks(ciphertext, rightKeySize)

	for _, block := range blocks {
		correctPlainText := FindCorrectPlainTextWithSingleByteXorCipher(block)
		correctPlainTextOfEachBlock = append(correctPlainTextOfEachBlock, correctPlainText)
	}

	return strings.Join(correctPlainTextOfEachBlock, "")
}
