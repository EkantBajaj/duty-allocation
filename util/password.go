package util

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"math/rand"
)

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("failed to hash password: %w", err)
	}
	return string(hashedPassword), nil
}

// CheckPassword that checks if the password is correct
func CheckPassword(hashedPassword string, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

// create RandomString function which returns a random string of length n
// this function is used to generate a random password
// this function is used to generate a random email
// this function is used to generate a random mobile number
func RandomString(n int) string {
	var letter = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	b := make([]rune, n)
	for i := range b {
		b[i] = letter[RandomInt(len(letter))]
	}
	return string(b)
}

// create RandomInt function which returns a random integer between 0 and n
func RandomInt(n int) int {
	return rand.Intn(n)
}
