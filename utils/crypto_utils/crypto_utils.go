package crypto_utils

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

//HashPassword takes password as string and returns salted password hash
func HashPassword(password string) string {
	var passwordByte = ConvertStringToByte(password)
	hash, err := bcrypt.GenerateFromPassword(passwordByte, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
		return ""
	}
	return string(hash)
}

//ConvertStringToByte takes a string and return byte array
func ConvertStringToByte(s string) []byte {
	return []byte(s)
}
