package iam

import (
	"bytes"
	"encoding/base64"
	"math/rand"
	"strings"
	"unicode"

	"golang.org/x/crypto/bcrypt"
)

var (
	digits  = []rune("0123456789")
	uppers  = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	lowers  = []rune(strings.ToLower("ABCDEFGHIJKLMNOPQRSTUVWXYZ"))
	symbols = []rune("\"`!?$?%^&*()_-+={[}]:;@'~#|\\<,>.?/")
)

const (
	strongThreshold     = 20
	veryStringThreshold = 40
)

func generateStrongPassword() string {
	password := ""
	var buffer bytes.Buffer
	var i int
	strong := false

	for !strong {
		switch rand.Intn(4) {
		case 0:
			i = rand.Intn(len(uppers))
			buffer.WriteRune(uppers[i])
			break
		case 1:
			i = rand.Intn(len(lowers))
			buffer.WriteRune(lowers[i])
			break
		case 2:
			i = rand.Intn(len(digits))
			buffer.WriteRune(digits[i])
			break
		case 3:
			i = rand.Intn(len(symbols))
			buffer.WriteRune(symbols[i])
			break
		}
		password = buffer.String()
		if len(password) >= 7 {
			strong = isStrong(password)
		}
	}
	return password
}

func isStrong(password string) bool {
	return calculateStrength(password) >= strongThreshold
}

func isVeryStong(password string) bool {
	return calculateStrength(password) >= veryStringThreshold
}

func isWeak(password string) bool {
	return calculateStrength(password) < strongThreshold
}

func calculateStrength(password string) int {
	strength := 0
	length := len(password)

	if length > 7 {
		strength += 10 + (length - 7)
	}

	digits := 0
	letters := 0
	lowers := 0
	uppers := 0
	symbols := 0

	for _, ch := range []rune(password) {
		switch {
		case unicode.IsLetter(ch):
			letters++
			if unicode.IsUpper(ch) {
				uppers++
			} else {
				lowers++
			}
			break
		case unicode.IsDigit(ch):
			digits++
		default:
			symbols++
		}
	}
	strength += uppers + lowers + symbols
	if letters >= 2 && digits >= 2 {
		strength += letters + digits
	}

	return strength
}

func encrypt(value string) (string, error) {
	enc, err := bcrypt.GenerateFromPassword([]byte(value), bcrypt.DefaultCost)
	if err != nil {
		return "", &Error{
			Code:    EINTERNAL,
			Message: "An unexpected error occurred while encrypting password.",
			Op:      "encrypt",
			Err:     err,
		}
	}
	return base64.StdEncoding.EncodeToString(enc), nil
}
