package cryptopals

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"strings"
	"testing"
)

func TestProblem1(t *testing.T) {
	res, err := hexToBase64("49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d")
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
	if err != nil {
		t.Fatal("Failed to decode hex:", s)
	}
	return v
}

func corpusFromFile(name string) map[rune]float64 {
	text, err := ioutil.ReadFile(name)
	if err != nil {
		panic(fmt.Sprintln("failed to read corpus file:", err))
	}
	return buildCorpus(string(text))
}

var corpus = corpusFromFile("_testData/aliceinwonderland.txt")

func TestProblem3(t *testing.T) {
	res, _ := findSingleXORKey(hexDecode(t, "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"), corpus)
	t.Logf("%s", res)
}

func TestProblem4(t *testing.T) {
	text, err := ioutil.ReadFile("4.txt")
	if err != nil {
		t.Fatal("failed to read challenge file:", err)
	}
	var lastScore float64
	var res []byte
	for _, line := range strings.Split(string(text), "\n") {
		out, score := findSingleXORKey(hexDecode(t, line), corpus)
		if score > lastScore {
			res = out
			lastScore = score
		}
	}
	t.Logf("%s", res)
}
