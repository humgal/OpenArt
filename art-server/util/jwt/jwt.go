package jwt

import (
	"errors"
	"log"
	"math/rand"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/humgal/art-server/util/redis"
)

// secret key being used to sign tokens
// random gen secretkey and store it in redis
// var (
// 	SecretKey = []byte("secret")
// )

// data we save in each token
type Claims struct {
	username string
	jwt.StandardClaims
}

func RandomString(n int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	s := make([]rune, n)
	for i := range s {
		s[i] = letters[rand.Intn(len(letters))]
	}
	return string(s)
}

// GenerateToken generates a jwt token and assign a username to it's claims and return it
func GenerateToken(username string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	/* Create a map to store our claims */
	claims := token.Claims.(jwt.MapClaims)
	/* Set token claims */
	claims["username"] = username
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	SecretKey := RandomString(10)
	tokenString, err := token.SignedString([]byte(SecretKey))
	if err != nil {
		log.Fatal("Error in Generating key")
		return "", err
	}
	redis.Rdb.Set(redis.Rdb.Context(), tokenString, SecretKey, time.Hour*24)
	return tokenString, nil
}

// ParseToken parses a jwt token and returns the username it it's claims
func ParseToken(tokenStr string) (string, error) {

	SecretKey := redis.Rdb.Get(redis.Rdb.Context(), tokenStr).Val()
	if SecretKey == "" {
		return "", errors.New("token过期")
	}
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {

		return []byte(SecretKey), nil
	})
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		username := claims["username"].(string)
		return username, nil
	} else {
		return "", err
	}
}
