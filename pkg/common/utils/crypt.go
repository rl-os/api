package utils

import (
	"golang.org/x/crypto/bcrypt"
)

// GetHash from string
func GetHash(pwd string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}

// CompareHash and return error if not equil
func CompareHash(hashed, plain string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(plain))
	if err != nil {
		return false
	}

	return true
}
