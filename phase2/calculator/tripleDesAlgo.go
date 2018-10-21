package calculator

import (
	"crypto/cipher"
	"crypto/des"
	"fmt"
)

type tripleDES struct {
	cipher1, cipher2, cipher3 string
}

func TriplDESAlgo() {
	var cipher1, cipher2, cipher3, text string
	fmt.Print("Enter cipher 1 : ")
	fmt.Scanln(&cipher1)
	fmt.Print("Enter cipher 2 : ")
	fmt.Scanln(&cipher2)
	fmt.Print("Enter cipher 3 : ")
	fmt.Scanln(&cipher3)
	triplekey := cipher1 + cipher2 + cipher3
	fmt.Println("Enter plain text to convert into unreadable text: ")
	fmt.Scanln(&text)
	plaintext := []byte(text)

	block, err := des.NewTripleDESCipher([]byte(triplekey))

	if err != nil {
		panic(err)
	}

	ciphertext := []byte("abcdef1234567890")
	iv := ciphertext[:des.BlockSize] // const BlockSize = 8

	// encrypt tripleDES

	mode := cipher.NewCBCEncrypter(block, iv)
	encrypted := make([]byte, len(plaintext))
	mode.CryptBlocks(encrypted, plaintext)
	fmt.Println("After Encryption : ", string(encrypted[:]))

	//decrypt tripleDES
	decrypter := cipher.NewCBCDecrypter(block, iv)
	decrypted := make([]byte, len(plaintext))
	decrypter.CryptBlocks(decrypted, encrypted)
	fmt.Println("After Decryption : ", string(decrypted[:]))

}
