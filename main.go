package main

import (
	"fmt"

	"cryptopals/cryptopals"
)

func main() {
	fmt.Println(cryptopals.FixedXor("1c0111001f010100061a024b53535009181c", "686974207468652062756c6c277320657965"))
}