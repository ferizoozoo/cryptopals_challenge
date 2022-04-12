package cryptopals

import (
	"encoding/hex"
	"math"
	"regexp"
	"unicode/utf8"
)

func SingleByteXorCipher(hexStringInput string, key byte) string {
	hexBytes, err := hex.DecodeString(hexStringInput)
	if err != nil {
		panic(err)
	}

	cipherText := make([]byte, len(hexBytes))

	for i := 0; i < len(hexBytes); i++ {
		cipherText[i] = hexBytes[i] ^ key
	}

	return string(cipherText)
}

var IsLetter = regexp.MustCompile(`^[a-zA-Z]+$`).MatchString

func FindApropriateKeyWithSingleByteXorCipher(hexStringInput string) (byte, string) {
	letterFrequency := map[byte]float64{
		'e': 0.1202,
		't': 0.0910,
		'a': 0.0812,
		'o': 0.0768,
		'i': 0.0731,
		'n': 0.0695,
		's': 0.0628,
		'r': 0.0602,
		'h': 0.0592,
		'd': 0.0432,
		'l': 0.0398,
		'u': 0.0288,
		'c': 0.0271,
		'm': 0.0261,
		'f': 0.0230,
		'y': 0.0211,
		'w': 0.0209,
		'g': 0.0203,
		'p': 0.0182,
		'b': 0.0149,
		'v': 0.0111,
		'k': 0.0069,
		'x': 0.0017,
		'q': 0.0011,
		'j': 0.0010,
		'z': 0.0007,
	}

	minScore := math.Inf(1)
	var key byte
	correctPlainText := ""

	for i := 0; i < 256; i++ {
		// get cipher text and check for letter frequencies
		plainText := SingleByteXorCipher(hexStringInput, byte(i))
		plainTextLength := utf8.RuneCountInString(plainText)

		letterCountsInPlainText := make(map[byte]int)
		for _, char := range plainText {
			if IsLetter(string(char)) {
				if _, ok := letterCountsInPlainText[byte(char)]; ok {
					letterCountsInPlainText[byte(char)]++
				} else {
					letterCountsInPlainText[byte(char)] = 1
				}
			}
		}

		score := 0.0
		for _, char := range plainText {
			score += letterFrequency[byte(char)] - float64(letterCountsInPlainText[byte(char)])/float64(plainTextLength)
		}

		if score <= minScore {
			correctPlainText = plainText
			minScore = score
			key = byte(i)
		}

	}
	// return the text with min score and it's key
	return key, correctPlainText
}
