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
	return err == nil // err is nil if the password and hash match
}

func ValidateToken(token string) bool {
	password := config.GetPassword()
	// для примера возьмём токен, подписанный при помощи секретного ключа secretKey
	secretKey := config.TokenSecret

	// парсим токен
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
	// приводим поле Claims к типу jwt.MapClaims
	res, ok := jwtToken.Claims.(jwt.MapClaims)
	// обязательно используем второе возвращаемое значение ok и проверяем его, потому что
	// если Сlaims вдруг окажется другого типа, мы получим панику
	if !ok {
		fmt.Printf("failed to type assertion of jwt.MapClaims")
		return false
	}
	// Так как jwt.Claims — словарь вида map[string]inteface{}, используем синтаксис получения
	// значения по ключу. Получаем значение ключа "login" и "roles"
	hashRaw := res["hash"]
	// loginRaw — интерфейс, так как тип значения в jwt.Claims — интерфейс.
	// Чтобы получить строку, нужно снова сделать приведение типа к строке.
	hash, ok := hashRaw.(string)
	fmt.Println(hash)
	if !ok {
		fmt.Printf("failed to typecast to string")
		return false
	}
	// TODO: compare hashes
	return CheckPasswordHash(password, hash)
}
