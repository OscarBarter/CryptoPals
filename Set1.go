package cryptopals

import (
	"encoding/base64"
	"encoding/hex"
	"log"
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
