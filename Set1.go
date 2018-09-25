package cryptopals

import (
	"encoding/base64"
	"encoding/hex"
	"log"
	"unicode/utf8"
)

func hexToBase64(hs string) (string, error) {
	v, err := hex.DecodeString(hs)
	if err != nil {
		return "", err
	}
	log.Printf("%s", v)
	return base64.StdEncoding.EncodeToString(v), nil
}

func xor(a, b []byte) []byte {
	if len(a) != len(b) {
		panic("xor: mismatched lengths")
	}
	res := make([]byte, len(a))
	for i := range a {
		res[i] = a[i] ^ b[i]
	}
	return res
}

func buildCorpus(text string) map[rune]float64 {
	c := make(map[rune]float64)
	for _, char := range text {
		c[char]++
	}
	total := utf8.RuneCountInString(text)
	for char := range c {
		c[char] = c[char] / float64(total)
	}
	return c
}

func scoreEnglish(text string, c map[rune]float64) float64 {
	var score float64
	for _, char := range text {
		score += c[char]
	}
	return score / float64(utf8.RuneCountInString(text))
}

func singleXOR(in []byte, key byte) []byte {
	res := make([]byte, len(in))
	for i, c := range in {
		res[i] = c ^ key
	}
	return res
}

func findSingleXORKey(in []byte, c map[rune]float64) []byte {
	var res []byte
	var lastScore float64
	for key := 0; key < 256; key++ {
		out := singleXOR(in, byte(key))
		score := scoreEnglish(string(out), c)
		if score > lastScore {
			res = out
			lastScore = score
		}
	}
	return res
}
