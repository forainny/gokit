package password

import (
	"golang.org/x/crypto/bcrypt"
)

// Encrypt passwords using the bcrypt algorithm.
func Encrypt(password string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes)
}

// Verify the password is correct.
func IsValid(hash, password string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)); err != nil {
		return false
	}
	return true
}
