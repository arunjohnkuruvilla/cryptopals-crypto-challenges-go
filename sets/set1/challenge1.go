package set1

import (
	"encoding/hex"
	"encoding/base64"
)

func HexToBase64(hexString string) (string, error) {
	bytes, err := hex.DecodeString(hexString)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(bytes), nil
}
