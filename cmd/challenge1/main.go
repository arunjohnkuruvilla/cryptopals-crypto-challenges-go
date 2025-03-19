package main

import (
	"fmt"
	"log"
	"cryptopals-crypto-challenges-go/sets/set1"
)

func main() {
	hexInput := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
	
	base64Output, err := set1.HexToBase64(hexInput)
	
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Base64 Output:", base64Output)
}