package main

import (
	"encoding/base64"
	"encoding/hex"
	"errors"
	"math"
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
	raw_bytes1, err := hex.DecodeString(string(buf1))
  if err != nil {
    return "", err
  }
	raw_bytes2, err := hex.DecodeString(string(buf2))
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

func DecryptSingleByteXOR(src string) (string, error) {
  raw_bytes, err := hex.DecodeString(src)
  if err != nil {
    return "", err
  }

	var bestChar byte
	minLoss := float64(math.MaxFloat64)
	for i := 0; i < 256; i++ {
		char := byte(i)
		if loss := getLossForBuffer(DecryptSingleByteXORBuffer(raw_bytes, char)); loss < minLoss {
			minLoss = loss
			bestChar = char
		}
	}

	return string(DecryptSingleByteXORBuffer(raw_bytes, bestChar)), nil
}
