package main

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"log"
)

func Challenge1() {
	src := []byte("49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d")

	raw_bytes := make([]byte, hex.DecodedLen(len(src)))
	n, err := hex.Decode(raw_bytes, src)
	if err != nil {
		log.Fatal(err)
	}

	raw_bytes = raw_bytes[:n]

	base64_formatted := base64.StdEncoding.EncodeToString(raw_bytes)
	fmt.Printf("%s\n", base64_formatted)
}