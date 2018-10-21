package calculator

import "fmt"

type caesarcipher interface {
	Encrypt() string
	Decrypt() string
}

func CaesarCipher() {
	var text string
	var key int
	fmt.Println("Write plain text: ")
	fmt.Scanln(&text)
	fmt.Println("enter key to shift: ")
	fmt.Scanln(&key)
	encrypted := Encrypt(text, -key)
	fmt.Println("Cipher text: ", encrypted)
	decrypted := Decrypt(encrypted, -key)
	fmt.Println("DeCipher text: ", decrypted)
}
func Encrypt(text string, key int) string {
	shift, offset := rune(key), rune(26)
	runes := []rune(text)
	for index, char := range runes {
		if char >= 'a'+shift && char <= 'z' ||
			char >= 'A'+shift && char <= 'Z' {
			char = char - shift
		} else if char >= 'a' && char < 'a'+shift ||
			char >= 'A' && char < 'A'+shift {
			char = char - shift + offset
		}
		runes[index] = char
	}
	return string(runes)
}
func Decrypt(text string, key int) string {
	shift, offset := rune(key), rune(26)
	runes := []rune(text)
	for index, char := range runes {
		if char >= 'a' && char <= 'z'-shift ||
			char >= 'A' && char <= 'Z'-shift {
			char = char + shift
		} else if char > 'z'-shift && char <= 'z' ||
			char > 'Z'-shift && char <= 'Z' {
			char = char + shift - offset
		}
		runes[index] = char
	}
	return string(runes)
}
