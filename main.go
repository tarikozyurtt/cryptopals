package main

import "fmt"

func main() {
  input := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"

	resbuf, err := HexToBase64(input)
  if err != nil {
    fmt.Println(err)
  }
  fmt.Printf("%s\n", resbuf)

  input1 := "1c0111001f010100061a024b53535009181c"
  input2 := "686974207468652062756c6c277320657965"


	res, err := XORBuffers(input1, input2)
  if err != nil {
    fmt.Println(err)
  }
	fmt.Printf("%s\n", res)

  res, err = DecryptSingleByteXOR("1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736")
  if err != nil {
    fmt.Println(err)
  }
	fmt.Printf("%s\n", res)
}