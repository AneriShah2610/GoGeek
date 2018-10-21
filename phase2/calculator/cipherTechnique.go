package calculator

import "fmt"

func CipherTechnique() {
	var cipherTechnique string
	fmt.Println("Select any one cipher technique from given list : \n 1. Caesar Cipher \n 2. Polyalphabetic Cipher \n 3. DES algo \n 4. Triple DES \n 5. AES algo")
	fmt.Scanln(&cipherTechnique)
	switch {
	case cipherTechnique == "Caesar" || cipherTechnique == "1":
		CaesarCipher()
	case cipherTechnique == "Polyalphabetic" || cipherTechnique == "2":
		PolyalphabeticCipher()
	case cipherTechnique == "DES" || cipherTechnique == "3":
		DESAlgo()
	case cipherTechnique == "TDES" || cipherTechnique == "4":
		TriplDESAlgo()
	case cipherTechnique == "AES" || cipherTechnique == "5":
		AESAlgo()
	}
}
