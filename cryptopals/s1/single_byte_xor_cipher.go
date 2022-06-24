package s1

import (
	"cryptopals/cryptopals/utils"
	"math"
	"regexp"
)

func SingleByteXorCipher(hexInput []byte, key byte) []byte {
	cipherText := make([]byte, len(hexInput))

	for i := 0; i < len(hexInput); i++ {
		cipherText[i] = hexInput[i] ^ key
	}

	return cipherText
}

func FindCorrectPlainTextWithSingleByteXorCipher(hexInput []byte) string {
	plainTextScores := findScores(hexInput)
	correctPlainText, _ := minOfPlainTextScoreMap(plainTextScores)
	return correctPlainText
}

func minOfPlainTextScoreMap(m map[string]float64) (string, float64) {
	minVal := math.Inf(1)
	minKey := "a"

	for k, v := range m {
		if v <= minVal {
			minVal = v
			minKey = k
		}
	}
	return minKey, minVal
}

func findScores(hexInput []byte) map[string]float64 {
	isLetter := regexp.MustCompile(`^[a-zA-Z]+$`).MatchString
	letterFrequency := utils.LetterFrequency()

	var plainTextScore = make(map[string]float64)

	for i := 0; i < 256; i++ {
		// get cipher text and check for letter frequencies
		plainText := SingleByteXorCipher(hexInput, byte(i))
		plainTextLength := len(plainText)

		letterCountsInPlainText := make(map[byte]int)
		for _, b := range plainText {
			if isLetter(string(b)) {
				if _, ok := letterCountsInPlainText[b]; ok {
					letterCountsInPlainText[b]++
				} else {
					letterCountsInPlainText[b] = 1
				}
			}
		}

		score := 0.0
		for _, b := range plainText {
			score += letterFrequency[b] - float64(letterCountsInPlainText[b])/float64(plainTextLength)
		}

		plainTextScore[string(plainText)] = score
	}
	// return the text with min score and it's key
	return plainTextScore
}
