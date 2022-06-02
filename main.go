package main

import (
	"bufio"
	"cryptopals/cryptopals"
	"encoding/hex"
	"fmt"
	"os"
)

func main() {
	// fmt.Println(string(cryptopals.HexToBase64String("49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d")))
	// msg1, _ := hex.DecodeString("1c0111001f010100061a024b53535009181c")
	// msg2, _ := hex.DecodeString("686974207468652062756c6c277320657965")
	// fmt.Println(string(cryptopals.FixedXor(msg1, msg2)))
	// msg, _ := hex.DecodeString("1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736")
	// fmt.Printf("%s", cryptopals.FindCorrectPlainTextWithSingleByteXorCipher(msg))

	var messages [][]byte
	file, err := os.Open("C:/Users/Asus/Desktop/8.txt")
	if err != nil {
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		plaintext := scanner.Text()
		if plaintext == "" {
			break
		}
		hexPT, _ := hex.DecodeString(plaintext)
		messages = append(messages, hexPT)
	}

	// fmt.Printf("%s", cryptopals.DetectSingleCharacterXor(messages))
	// res := cryptopals.RepeatingKeyXorEncryption([]byte("Burning 'em, if you ain't quick and nimble\nI go crazy when I hear a cymbal"), []byte("ICE"))
	// fmt.Println(hex.EncodeToString(res))
	// fmt.Println(cryptopals.HammingDistance([]byte("this is a test"), []byte("wokka wokka!!!")))

	// file, err := os.Open("C:/Users/Asus/Desktop/8.txt")
	// if err != nil {
	// 	return
	// }
	// defer file.Close()

	// bytes, err := ioutil.ReadAll(hex.NewDecoder(hex.StdEncoding, file))
	// if err != nil {
	// 	return
	// }

	// fmt.Print(cryptopals.FindEncryptionKey(bytes))

	// fmt.Println(string(cryptopals.DecryptAES128ECBWith(bytes, []byte("YELLOW SUBMARINE"), 16)))
	fmt.Println(hex.EncodeToString(cryptopals.DetectAES128ECB(messages)))
}
