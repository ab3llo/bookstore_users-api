package crypto_utils

import (
	"log"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

//HashPassword take plaintext password and salt and return hash
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(strings.TrimSpace(password)), bcrypt.DefaultCost)
	log.Println("Password hash:" + string(bytes))
	return string(bytes), err
}

//Compare compares passsword in db with password supplied
func Compare(hash, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(strings.TrimSpace(password)))
	if err != nil {
		log.Println(err.Error())
	}
	return err == nil
}

//ConvertStringToByte takes a string and return byte array
func ConvertStringToByte(s string) []byte {
	return []byte(s)
}
