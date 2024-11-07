package main

import (
	"math"
	"strings"
)

// from solution of 0xfe
func getExpectedFreqForChar(char byte) float64 {
	// Default value (helps prevent divide-by-zero)
	value := float64(0.00001)

	freqMap := map[byte]float64{
		' ':  10,
		'\'': 0.1,
		'\n': 0.1,
		',':  0.1,
		'.':  0.1,
		'E':  12.02,
		'T':  9.1,
		'A':  8.12,
		'O':  7.68,
		'I':  7.31,
		'N':  6.95,
		'S':  6.28,
		'R':  6.02,
		'H':  5.92,
		'D':  4.32,
		'L':  3.98,
		'U':  2.88,
		'C':  2.71,
		'M':  2.61,
		'F':  2.3,
		'Y':  2.11,
		'W':  2.09,
		'G':  2.03,
		'P':  1.82,
		'B':  1.49,
		'V':  1.11,
		'K':  0.69,
		'X':  0.17,
		'Q':  0.11,
		'J':  0.10,
		'Z':  0.1,
		'0':  0.1,
		'1':  0.2,
		'2':  0.1,
		'3':  0.1,
		'4':  0.1,
		'5':  0.1,
		'6':  0.1,
		'7':  0.1,
		'8':  0.1,
		'9':  0.1,
	}

	if freq, ok := freqMap[strings.ToUpper(string(char))[0]]; ok {
		value = freq
	}

	return value
}

func getCharFreqMap(buf []byte) map[byte]float64 {
	charFreq := make(map[byte]float64)
	for _, char := range buf {
		charFreq[char]++
	}

	for char, freq := range charFreq {
		charFreq[char] = freq / float64(len(buf))
	}

	return charFreq
}

func getLossForBuffer(buf []byte) float64 {
	charFreqMap := getCharFreqMap(buf)
	loss := 0.0

	for char, observedFreq := range charFreqMap {
		expectedFreq := getExpectedFreqForChar(char)
		loss += math.Pow(expectedFreq-observedFreq, 2) / expectedFreq
	}

	return loss
}

func DecryptSingleByteXORBuffer(buf []byte, key byte) []byte {
	result := make([]byte, len(buf))
	for i := range buf {
		result[i] = buf[i] ^ key
	}

	return result
}
