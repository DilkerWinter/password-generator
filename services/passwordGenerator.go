package services

import (
	"crypto/rand"
	"math/big"
)

func PasswordGenerator(length int,includeSpecialChar bool,includeUpperCaseLetters bool, passwordName string) string {
	
	uppercase := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	lowercase := "abcdefghijklmnopqrstuvwxyz"
	numbers := "0123456789"
	specialChars := "!@#$%^&*()-_=+[]{}|;:,.<>?/~"

	characters := lowercase + numbers
	if includeSpecialChar {
		characters += specialChars
	}
	if includeUpperCaseLetters {
		characters += uppercase
	}

	var password []rune
	for i := 0; i < length; i++ {
		randomIndex, _ := rand.Int(rand.Reader, big.NewInt(int64(len(characters))))
		password = append(password, rune(characters[randomIndex.Int64()]))
	} 

	return string(password)

}
