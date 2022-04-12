package main

import (
	"cryptopals/cryptopals"
	"fmt"
)

func main() {
	// fmt.Println(cryptopals.FixedXor("1c0111001f010100061a024b53535009181c", "686974207468652062756c6c277320657965"))
	key, plainText := cryptopals.FindApropriateKeyWithSingleByteXorCipher("1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736")
	fmt.Printf("%s: %s", string(key), plainText)

	// cryptopals.DetectSingleCharacterXor()
}
