package utils

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

func JamWaktu() string {
	lokasi, _ := time.LoadLocation("Asia/Jakarta")

	t := time.Now().In(lokasi).Format("2006-01-02 15:04:05")

	return t
}

func HashPassword(pwd string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(pwd), 10)

	return string(bytes), err
}

func CheckPassword(pwd, dbPwd string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(dbPwd), []byte(pwd))

	if err != nil {
		panic(err)
	}

	return true
}
