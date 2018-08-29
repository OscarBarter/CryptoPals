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