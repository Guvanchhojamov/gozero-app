package repository

import (
	"crypto/sha1"
	"fmt"
	"time"
)

const (
	//salt      = "123"
	tokenTLL = time.Hour * 12
	//signedKey = "abc"
)

func generateHashPassword(password, salt string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(salt))) //Sum()  is hashedString+salt
}
