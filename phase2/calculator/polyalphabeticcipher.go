package calculator

import (
	"fmt"
)

type polyalphabeticKey string
type polyalphabeiccipher interface {
	PolyalphabeticEncipher() string
	PolyalphabeticDecipher() string
}

func newVigenère(key string) (polyalphabeticKey, bool) {
	v := polyalphabeticKey(upperOnly(key))
	return v, len(v) > 0 // key length 0 invalid
}
func (k polyalphabeticKey) PolyalphabeticEncipher(pt string) string {
	ct := upperOnly(pt)
	for i, c := range ct {
		ct[i] = 'A' + (c-'A'+k[i%len(k)]-'A')%26
	}
	return string(ct)
}
func (k polyalphabeticKey) PolyalphabeticDecipher(ct string) (string, bool) {
	pt := make([]byte, len(ct))
	for i := range pt {
		c := ct[i]
		if c < 'A' || c > 'Z' {
			return "", false // invalid ciphertext
		}
		pt[i] = 'A' + (c-k[i%len(k)]+26)%26
	}
	return string(pt), true
}
func upperOnly(s string) []byte {
	u := make([]byte, 0, len(s))
	for i := 0; i < len(s); i++ {
		c := s[i]
		if c >= 'A' && c <= 'Z' {
			u = append(u, c)
		} else if c >= 'a' && c <= 'z' {
			u = append(u, c-32)
		}
	}
	return u
}

func PolyalphabeticCipher() {
	var plianText, secretKey string
	fmt.Println("Enter plain text:")
	fmt.Scanln(&plianText)
	fmt.Println("Enter secret key:")
	fmt.Scanln(&secretKey)
	v, ok := newVigenère(secretKey)
	if !ok {
		fmt.Println("Invalid key")
		return
	}
	fmt.Println("Effective key:", v)
	fmt.Println("Plain text:", plianText)
	ct := v.PolyalphabeticEncipher(plianText)
	fmt.Println("Enciphered text:", ct)
	dt, ok := v.PolyalphabeticDecipher(ct)
	if !ok {
		fmt.Println("Invalid ciphertext")
		return
	}
	fmt.Println("Deciphered text:", dt)
}
