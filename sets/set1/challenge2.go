package set1

import (
	"fmt"
)

func XOR(xorString1, xorString2 []byte) ([]byte, error) {
	if len(xorString1) != len(xorString2) {
		return nil, fmt.Errorf("Input buffers must be of equal length")
	}

	result := make([]byte, len(xorString1))

	for i:= range xorString1 {
		result[i] = xorString1[i] ^ xorString2[i]
	}

	return result, nil
}