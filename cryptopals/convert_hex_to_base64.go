package cryptopals

import (
	"encoding/base64"
	"encoding/hex"
)

func HexToBase64String(hexString string) []byte {
	hexBytes, err := hex.DecodeString(hexString)
	if err != nil {
		panic(err)
	}

	dst := make([]byte, base64.StdEncoding.EncodedLen(len(hexBytes)))
	base64.StdEncoding.Encode(dst, hexBytes)

	return dst
}
