package tokenutils

import (
	"fmt"

	"github.com/golang-jwt/jwt"
)

func CreateToken(id uint) (string, error) {
	payload := jwt.MapClaims{
		"id": id,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	signedToken, err := token.SignedString([]byte("secret"))
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

func VerifyToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Verify the signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		// Provide the key used for signing
		return []byte("secret"), nil
	})
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}
	return claims, nil
}
