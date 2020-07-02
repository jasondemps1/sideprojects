package vigenere

import (
	"strings"
)

// Encrypt - Encrypt a plaintext string with a keyword
func Encrypt(plainText, keyword string) (cipherText string) {
	plainText = strings.ToUpper(strings.ReplaceAll(plainText, " ", ""))
	keyword = strings.ToUpper(strings.ReplaceAll(keyword, " ", ""))
	keywordRepeat := ""

	for i := 0; i < len(plainText); i++ {
		keywordRepeat += string(keyword[i%len(keyword)])
	}

	for i := 0; i < len(plainText); i++ {
		plainByte := plainText[i] - 'A'
		keywordByte := keywordRepeat[i] - 'A'
		result := ((plainByte + keywordByte) % 26) + 'A'

		cipherText += string(result)
	}

	return
}

// Decrypt - Decipher a ciphertext string with a keyword
func Decrypt(cipherText, keyword string) (decipherText string) {
	cipherText = strings.ToUpper(strings.ReplaceAll(cipherText, " ", ""))
	keyword = strings.ToUpper(strings.ReplaceAll(keyword, " ", ""))
	keywordRepeat := ""

	for i := 0; i < len(cipherText); i++ {
		keywordRepeat += string(keyword[i%len(keyword)])
	}

	for i := 0; i < len(cipherText); i++ {
		cipherByte := cipherText[i] - 'A'
		keywordByte := keywordRepeat[i] - 'A'
		result := ((cipherByte - keywordByte + 26) % 26) + 'A'

		decipherText += string(result)
	}

	return
}
