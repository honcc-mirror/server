package security

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) []byte {
	passwordBytes := []byte(password)
	hash, err := bcrypt.GenerateFromPassword(passwordBytes, bcrypt.DefaultCost)
	if err != nil {
		log.Fatalf("Could not hash password, got error: %s", err)
	}

	return hash
}

func PasswordMatchesHash(password string, hash []byte) bool {
	passwordBytes := []byte(password)
	err := bcrypt.CompareHashAndPassword(hash, passwordBytes)
	return err == nil
}
