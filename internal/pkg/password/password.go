package password

import (
	"bytes"
	"math/rand"
	"unicode"

	"golang.org/x/crypto/bcrypt"
)

const (
	pwdDigits           = "0123456789"
	pwdUpperLetters     = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	pwdLowerLetters     = "abcdefghijklmnopqrstuvwxyz"
	pwdSymbols          = "\"`!?$?%^&*()_-+={[}]:;@'~#|\\<,>.?/"
	strongThreshold     = 20
	veryStringThreshold = 40
)

// Encrypt will encrypt a password, returning the encrypted password or an error.
func Encrypt(password string) (string, error) {
	plain := []byte(password)
	hash, err := bcrypt.GenerateFromPassword(plain, 0)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

// Generate will generate a new strong password.
func Generate() string {
	var generatedPassword string
	strongPassword := false
	buffer := bytes.NewBufferString("")
	for !strongPassword {
		opt := rand.Intn(4)
		switch opt {
		case 0:
			index := rand.Intn(len(pwdUpperLetters))
			buffer.WriteByte(pwdUpperLetters[index])
		case 1:
			index := rand.Intn(len(pwdLowerLetters))
			buffer.WriteByte(pwdLowerLetters[index])
		case 2:
			index := rand.Intn(len(pwdDigits))
			buffer.WriteByte(pwdDigits[index])
		case 3:
			index := rand.Intn(len(pwdSymbols))
			buffer.WriteByte(pwdSymbols[index])
		}
		generatedPassword = buffer.String()
		if len(generatedPassword) >= 7 {
			strongPassword = IsStrong(generatedPassword)
		}
	}

	return generatedPassword
}

// IsWeak check if the supplied password is weak.
func IsWeak(password string) bool {
	return calculatePasswordStrength(password) < strongThreshold
}

// IsStrong check if the supplied password is strong.
func IsStrong(password string) bool {
	return calculatePasswordStrength(password) >= strongThreshold
}

// IsVeryStrong will check if the supplied password is very strong.
func IsVeryStrong(password string) bool {
	return calculatePasswordStrength(password) >= veryStringThreshold
}

func calculatePasswordStrength(password string) int {
	strength := 0
	length := len(password)
	if length > 7 {
		strength += 10 + (length - 7)
	}
	digitCount := 0
	letterCount := 0
	lowerCount := 0
	upperCount := 0
	symbolCount := 0

	for _, chr := range password {
		if unicode.IsLetter(chr) {
			letterCount++
			if unicode.IsUpper(chr) {
				upperCount++
			} else {
				lowerCount++
			}
		} else if unicode.IsDigit(chr) {
			digitCount++
		} else {
			symbolCount++
		}
	}
	strength += upperCount + lowerCount + symbolCount
	if letterCount >= 2 && digitCount >= 2 {
		strength += letterCount + digitCount
	}
	return strength
}
