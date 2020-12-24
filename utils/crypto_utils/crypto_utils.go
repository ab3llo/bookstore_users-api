package crypto_utils

import (
	"crypto/rand"

	"golang.org/x/crypto/bcrypt"
)

// Define salt size
const saltSize = 16

func GenerateRandomSalt(saltSize int) []byte {
	var salt = make([]byte, saltSize)

	_, err := rand.Read(salt[:])

	if err != nil {
		panic(err)
	}

	return salt
}

//HashPassword take plaintext password and salt and return hash
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 16)
	return string(bytes), err
}

//Compare compares passsword in db with password supplied
func Compare(hash, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

//ConvertStringToByte takes a string and return byte array
func ConvertStringToByte(s string) []byte {
	return []byte(s)
}
