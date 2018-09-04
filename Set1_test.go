package cryptopals

import (
	"bytes"
	"encoding/hex"
	"testing"
	"io/ioutil"
)

func TestProblem1(t *testing.T) {
	res, err := hexToBase64 ("49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d")
	if err != nil {
		t.Fatal(err)
	}
	if res != "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t" {
		t.Fatal("Wrong String", res)
	}
}

func TestProblem2(t *testing.T) {
	res := xor(hexDecode(t, "1c0111001f010100061a024b53535009181c"),
		hexDecode(t, "686974207468652062756c6c277320657965"))
	if !bytes.Equal(res, hexDecode(t, "746865206b696420646f6e277420706c6179")) {
		t.Errorf("wrong result: %x", res)
	}
}
func hexDecode(t *testing.T, s string) []byte {
	v, err := hex.DecodeString(s)
	if err != nil{
		t.Fatal("Failed to decode hex:", s)
	}
	return v
}

func corpusFromFile(t *testing.T, name string) map[rune]float64 {
	text, err := ioutil.ReadFile(name)
	if err != nil {
		t.Fatal("failed to open corpus file:", err)
	}
	return buildCorpus(string(text))
}

func TestProblem3(t *testing.T) {
	c := corpusFromFile(t, "_testData/aliceinwonderland.txt")
	for char, val := range c {
		t.Logf("%c: %.5f", char, val)
	}
}

func scoreEnglish(text string, c map[rune]float64) float64 {
	var score float64
	for _, char := range text {
		score += c[char]
	}
	return score / float64(utf8.RuneCountInString(text)
}

func singleXOR(in []byte, key byte) []byte { 
	res := make([]byte, len(in))
	for i, c := range in {
		res[i]  =c ^ key
	}
	return res
}
