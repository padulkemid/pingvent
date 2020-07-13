package utils

import (
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
)

// time
func JamWaktu() string {
	lokasi, _ := time.LoadLocation("Asia/Jakarta")

	t := time.Now().In(lokasi).Format("2006-01-02 15:04:05")

	return t
}

// hasher
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

// JWT

type TokenData struct {
	Role       string `json:"id"`
	Username string `json:"username"`
}

func getSecretKey() string {

  err := godotenv.Load()

  if err!= nil {
    panic(err)
  }

  JWT_SECRET := os.Getenv("JWT_SECRET")

  return JWT_SECRET
}

func GenerateToken(role, username string) (string, error){
  // get secret
  keyString := getSecretKey()
  key := []byte(keyString)

  token := jwt.New(jwt.SigningMethodHS256)
  claims := token.Claims.(jwt.MapClaims)

  // set claims
  claims["username"] = username
  claims["role"] = role
  claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

  // sign
  tokenString, err := token.SignedString(key)

  if err != nil {
    panic(err)
  }

  return tokenString, nil
}

func ParseToken(tokenString string)(*TokenData, error) {
  // get secret
  keyString := getSecretKey()
  key := []byte(keyString)

  // parse
  token, err := jwt.Parse(
    tokenString,
    func(token *jwt.Token)(interface{}, error) {
      return key, nil
  })

  if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
    username := claims["username"].(string)
    role := claims["role"].(string)

    data := &TokenData{
      Role: role,
      Username: username,
    }

    return data, nil
  } else {
    panic(err)
  }
}



