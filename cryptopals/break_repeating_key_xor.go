package cryptopals

import (
	"encoding/hex"
	"math"
	"strings"
)

var KEYSIZE int = 2

func HammingDistance(firstString string, secondString string) int {
	if len(firstString) != len(secondString) {
		return -1
	}

	distance := 0

	for index := 0; index < len(firstString); index++ {
		byte1 := firstString[index]
		byte2 := secondString[index]

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

func FindKeySize(message string) {
	minDistance := math.Inf(1)

	for keysize := 2; keysize <= 40; keysize++ {
		firstSubString := message[:KEYSIZE]
		secondSubString := message[KEYSIZE : KEYSIZE*2]

		distance := HammingDistance(firstSubString, secondSubString)

		if float64(distance) < minDistance {
			minDistance = float64(distance) / float64(keysize)
			KEYSIZE = keysize
		}
	}
}

func BreakCiphertextIntoTransposedBlocks(ciphertext string) [][]byte {
	var blocks [][]byte

	for byte_index := 0; byte_index < len(ciphertext)-KEYSIZE; byte_index += KEYSIZE {
		blocks = append(blocks, []byte(ciphertext)[byte_index:byte_index+KEYSIZE])
	}

	var transposedBlocks [][]byte

	for b := 0; b < KEYSIZE; b++ {
		var transposedBlock []byte
		for block := 0; block < len(blocks); block++ {
			transposedBlock = append(transposedBlock, blocks[block][b])
		}
		transposedBlocks = append(transposedBlocks, transposedBlock)
	}

	return transposedBlocks
}

func FindEncryptionKey(ciphertext string) string {
	var correctKeyOfEachBlock []string
	blocks := BreakCiphertextIntoTransposedBlocks(ciphertext)

	for _, block := range blocks {
		//TODO: bring the code for single byte xor cipher key finder here and change it to use just byte arrays
		hexString := hex.EncodeToString(block)
		correctKey := FindCorrectPlainTextWithSingleByteXorCipher(hexString)
		correctKeyOfEachBlock = append(correctKeyOfEachBlock, correctKey)
	}

	return strings.Join(correctKeyOfEachBlock, "")
}
