package iam

import (
	"bytes"
	"encoding/base64"
	"math/rand"
	"unicode"

	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

func encrypt(value string) (string, error) {
	encrypted, err := bcrypt.GenerateFromPassword([]byte(value), bcrypt.DefaultCost)
	if err != nil {
		return "", errors.Wrap(err, "An error occurred while using bcrypt to encrypt a value")
	}
	return base64.StdEncoding.EncodeToString(encrypted), nil
}

var (
	digits  = []rune("0123456789")
	letters = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	symbols = []rune("\"`!?$?%^&*()_-+={[}]:;@'~#|\\<,>.?/")
)

const (
	strongThreshold     = 20
	veryStrongThreshold = 40
)

func generateStrongPassword() string {
	generatedPassword := ""
	strong := false
	var password bytes.Buffer
	for !strong {
		opt := rand.Intn(4)
		switch opt {
		case 0:
			index := rand.Intn(len(letters))
			password.WriteRune(letters[index])
			break
		case 1:
			index := rand.Intn(len(letters))
			password.WriteRune(unicode.ToLower(letters[index]))
			break
		case 2:
			index := rand.Intn(len(digits))
			password.WriteRune(digits[index])
			break
		case 3:
			index := rand.Intn(len(symbols))
			password.WriteRune(symbols[index])
			break
		}
		generatedPassword = password.String()
		if len(generatedPassword) > 7 {
			strong = isStrong(generatedPassword)
		}
	}
	return generatedPassword
}

func isStrong(password string) bool {
	return calculatePasswordStrength(password) >= strongThreshold
}

func isVeryStrong(password string) bool {
	return calculatePasswordStrength(password) >= veryStrongThreshold
}

func isWeak(password string) bool {
	return calculatePasswordStrength(password) < strongThreshold
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
	for _, ch := range password {
		if unicode.IsLetter(ch) {
			letterCount++
			if unicode.IsUpper(ch) {
				upperCount++
			} else {
				lowerCount++
			}
		} else if unicode.IsDigit(ch) {
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
