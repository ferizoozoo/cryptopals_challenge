package main

import (
	"fmt"
	"math"
	"unicode"

	"cryptopals/cryptopals"
)

func main() {
	letterFrequency := map[byte]float64{'e': 12.02, 't': 9.10, 'a': 8.12, 'o': 7.68, 'i': 7.31,
		'n': 6.95, 's': 6.28, 'r': 6.02, 'h': 5.92, 'd': 4.32, 'l': 3.98, 'u': 2.88,
		'c': 2.71, 'm': 2.61, 'f': 2.30, 'y': 2.11, 'w': 2.09, 'g': 2.03, 'p': 1.82,
		'b': 1.49, 'v': 1.11, 'k': 0.69, 'x': 0.17, 'q': 0.11, 'j': 0.10, 'z': 0.07}
	// fmt.Println(cryptopals.FixedXor("1c0111001f010100061a024b53535009181c", "686974207468652062756c6c277320657965"))
	maxScore := math.Inf(-2)
	finalResult := ""
	for i := 0; i < 256; i++ {
		result := cryptopals.SingleByteXorCipher("1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736", byte(i))
		score := 0.0
		for _, char := range result {
			if unicode.IsLetter(char) {
				score += letterFrequency[byte(unicode.ToLower(char))]
			}
		}

		if score > maxScore {
			maxScore = score
			finalResult = result
		}
	}
	fmt.Println(finalResult)
}
