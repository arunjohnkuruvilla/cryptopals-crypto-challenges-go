package main

import (
	"fmt"
	"log"
	"encoding/hex"

	"cryptopals-crypto-challenges-go/sets/set1"
)

func main() {
	hexString1 := "1c0111001f010100061a024b53535009181c"
	xorString1, err := hex.DecodeString(hexString1)

	if err != nil {
		log.Fatal(err)
	}

	hexString2 := "686974207468652062756c6c277320657965"
	xorString2, err := hex.DecodeString(hexString2)

	if err != nil {
		log.Fatal(err)
	}

	hexOutput := "746865206b696420646f6e277420706c6179"

	result, err := set1.XOR(xorString1, xorString2)

	if err != nil {
		log.Fatal(err)
	}

	xorResult := hex.EncodeToString(result)

	fmt.Println("Output:", xorResult)

	if xorResult == hexOutput {
		fmt.Println("XOR results equal")
	} else {
		fmt.Println("XOR results not equal")
	}
}