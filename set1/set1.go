package set1

import (
	"bufio"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
	"log"
	"math"
	"os"
)

func HexToBase64(buf string) ([]byte, error) {
	raw_bytes, err := hex.DecodeString(buf)
	if err != nil {
		return nil, err
	}

	base64_bytes := make([]byte, base64.StdEncoding.EncodedLen(len(raw_bytes)))
	base64.StdEncoding.Encode(base64_bytes, raw_bytes)

	return base64_bytes, nil
}

func XORBuffers(buf1, buf2 string) (string, error) {
	if len(buf1) != len(buf2) {
		return "", errors.New("buffers must be of equal length")
	}
	raw_bytes1, err := hex.DecodeString(buf1)
	if err != nil {
		return "", err
	}
	raw_bytes2, err := hex.DecodeString(buf2)
	if err != nil {
		return "", err
	}

	result := make([]byte, len(raw_bytes1))
	for i := range raw_bytes1 {
		result[i] = raw_bytes1[i] ^ raw_bytes2[i]
	}

	res := hex.EncodeToString(result)

	return res, nil
}

func DecryptSingleByteXOR(src string) ([]byte, error) {
	raw_bytes, err := hex.DecodeString(src)
	if err != nil {
		return nil, err
	}

	var bestChar byte
	minLoss := float64(math.MaxFloat64)
	for i := 0; i < 256; i++ {
		char := byte(i)
		if loss := getLossForBuffer(decryptSingleByteXORBuffer(raw_bytes, char)); loss < minLoss {
			minLoss = loss
			bestChar = char
		}
	}

	return decryptSingleByteXORBuffer(raw_bytes, bestChar), nil
}

func findEncryptedStringWithSingleByteXOR(path string) ([]byte, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	scanner := bufio.NewScanner(file)
	bytesToDecrypt := make([]byte, 0)
	var bestChar byte
	minLoss := float64(math.MaxFloat64)

	for scanner.Scan() {
		raw_bytes, err := hex.DecodeString(scanner.Text())
		if err != nil {
			return nil, err
		}

		for i := 0; i < 256; i++ {
			char := byte(i)
			if loss := getLossForBuffer(decryptSingleByteXORBuffer(raw_bytes, char)); loss < minLoss {
				minLoss = loss
				bestChar = char
				bytesToDecrypt = raw_bytes
			}
		}
	}

	return decryptSingleByteXORBuffer(bytesToDecrypt, bestChar), nil
}

func RepeatingKeyXOR(src, key string) string {
	raw_bytes := []byte(src)
	key_bytes := []byte(key)
	result := make([]byte, len(raw_bytes))

	for i := range raw_bytes {
		result[i] = raw_bytes[i] ^ key_bytes[i%len(key_bytes)]
	}

	return hex.EncodeToString(result)
}

func BreakRepeatingKeyXOR(path string) ([]byte, error) {
	KEYSIZE := 2
	bestKeySize := 0
	minDistance := 1000.0

	base64_encoded, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	raw_bytes := make([]byte, base64.StdEncoding.DecodedLen(len(base64_encoded)))
	n, err := base64.StdEncoding.Decode(raw_bytes, base64_encoded)
	if err != nil {
		fmt.Println("decode error:", err)
		return nil, err
	}
	raw_bytes = raw_bytes[:n]

	var distance float64
	for KEYSIZE != 40 {
		iters := (len(raw_bytes) / KEYSIZE) - 1
		for i := 0; i < iters; i++ {
			a := raw_bytes[i*KEYSIZE : (i+1)*KEYSIZE]
			b := raw_bytes[(i+1)*KEYSIZE : (i+2)*KEYSIZE]
			distance += float64(hammingDistance(a, b))
		}

		distance = distance / float64(KEYSIZE) / float64(iters)
		if distance < minDistance {
			minDistance = distance
			bestKeySize = KEYSIZE
		}
		KEYSIZE++
	}

	fmt.Println("Best key size: ", bestKeySize)
	fmt.Println("raw_bytes: ", len(raw_bytes))

	// break the ciphertext into blocks of KEYSIZE length and transpose the blocks
	// make a block that is the first byte of every block, and a block that is the second
	// byte of every block, and so on.

	blocks := make([][]byte, bestKeySize)
	for i := 0; i < len(raw_bytes); i++ {
		blocks[i%bestKeySize] = append(blocks[i%bestKeySize], raw_bytes[i])
	}

	fmt.Println("blocks: ", len(blocks))
	fmt.Println("blocks[0]: ", len(blocks[0]))

	key := make([]byte, bestKeySize)
	for block := range blocks {
		var bestChar byte
		minLoss := float64(math.MaxFloat64)
		for i := 0; i < 256; i++ {
			char := byte(i)
			if loss := getLossForBuffer(decryptSingleByteXORBuffer(blocks[block], char)); loss < minLoss {
				minLoss = loss
				bestChar = char
			}
		}

		key[block] = bestChar
	}

	result := make([]byte, len(raw_bytes))

	for i := range raw_bytes {
		result[i] = raw_bytes[i] ^ key[i%len(key)]
	}

	return result, nil
}
