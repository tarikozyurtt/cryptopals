package set1

import "fmt"

func ApplyFirstChallenge() {
	fmt.Println("Challenge - 1")
	input := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"

	res, err := HexToBase64(input)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%s\n", res)
}

func ApplySecondChallenge() {
	fmt.Println("Challenge - 2")
	input1 := "1c0111001f010100061a024b53535009181c"
	input2 := "686974207468652062756c6c277320657965"

	res, err := XORBuffers(input1, input2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%s\n", res)
}

func ApplyThirdChallenge() {
	fmt.Println("Challenge - 3")
	input := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"

	res, err := DecryptSingleByteXOR(input)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%s\n", res)
}

func ApplyFourthChallenge() {
	fmt.Println("Challenge - 4")
	file_name := "set1-ch4.txt"

	res, err := findEncryptedStringWithSingleByteXOR(file_name)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%s", res)
}

func ApplyFifthChallenge() {
	fmt.Println("Challenge - 5")
	src := `Burning 'em, if you ain't quick and nimble
I go crazy when I hear a cymbal`
	key := "ICE"

	res := RepeatingKeyXOR(src, key)
	fmt.Printf("%s\n", res)
}

func ApplySixthChallenge() {
	fmt.Println("Challenge - 6")
	file_name := "set1-ch6.txt"

	res, err := BreakRepeatingKeyXOR(file_name)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%s", res)
}
