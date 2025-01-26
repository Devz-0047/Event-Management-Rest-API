package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	const cost = 12 // Adjust as needed
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	return string(bytes), err
}

func CheckPasswordHash(password, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
