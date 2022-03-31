package cryptopals

import (
	"encoding/base64"
	"encoding/hex"
)

func HexToBase64String(hexString string) string {
	hexBytes, err := hex.DecodeString(hexString)
	if err != nil {
		panic(err)
	}

	return base64.StdEncoding.EncodeToString(hexBytes)
}
