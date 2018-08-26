package main

import (
	"crypto/sha256"
	"encoding/hex"
)

func doubleSHA256(s []byte) []byte {
	h := sha256.New()
	h.Write(s)
	h2 := sha256.New()
	h2.Write(h.Sum(nil))
	return h2.Sum(nil)
}

//reverse array in place
func reverseArrayInPlace(a []byte) {
	for i := len(a)/2 - 1; i >= 0; i-- {
		opp := len(a) - 1 - i
		a[i], a[opp] = a[opp], a[i]
	}
}

func reverseArray(a []byte) []byte{
	b := make([]byte, len(a))
	for i,v := range a {
		b[i] = v
	}
	return b
}

func hexStringToByteArray(s *string) []byte {
	ba, _ := hex.DecodeString(*s)
	return ba
}



