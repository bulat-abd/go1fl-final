package api

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"

	"github.com/bulat-abd/go1fl-final/internal/config"
	"github.com/golang-jwt/jwt/v5"
)

func GeneratePasswordHash(password string) string {
	result, _ := HashPassword(password)
	return result

}

func CreateToken() (string, error) {
	password := config.GetPassword()
	secret := config.TokenSecret
	claims := jwt.MapClaims{
		"hash": GeneratePasswordHash(password),
	}
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := jwtToken.SignedString(secret)
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

func HashPassword(password string) (string, error) {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedBytes), nil
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func ValidateToken(token string) bool {
	password := config.GetPassword()
	secretKey := config.TokenSecret

	jwtToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
	if err != nil {
		fmt.Printf("failed to parse token: %s\n", err)
		return false
	}
	if !jwtToken.Valid {
		fmt.Printf("token is invalid")
		return false
	}
	res, ok := jwtToken.Claims.(jwt.MapClaims)
	if !ok {
		fmt.Printf("failed to type assertion of jwt.MapClaims")
		return false
	}
	hashRaw := res["hash"]
	hash, ok := hashRaw.(string)
	if !ok {
		fmt.Printf("failed to typecast to string")
		return false
	}
	return CheckPasswordHash(password, hash)
}
