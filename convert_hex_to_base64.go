package cryptopals

import (
	"encoding/base64"
	"encoding/hex"
)

func hexToBase64String(hexString string) string {
	hexBytes, err := hex.DecodeString(hexString)
	if err != nil {
		panic(err)
	}

	return base64.StdEncoding.EncodeToString(hexBytes)
}
