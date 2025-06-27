package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// Sua chave secreta (coloque algo forte em produção)
var jwtSecret = []byte("minha-chave-secreta")

func GenerateJWT(userID uint, isAdmin bool) (string, error) {
	claims := jwt.MapClaims{
		"idUsuario": userID,
		"is_admin":  isAdmin,
		"exp":       time.Now().Add(24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

// ValidateJWT valida o token e retorna os claims
func ValidateJWT(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Confere se o método de assinatura é HMAC
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("método de assinatura inválido")
		}
		return jwtSecret, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("token inválido")
}

// AdminOnly retorna true se os claims indicarem que o usuário é admin
func AdminOnly(claims jwt.MapClaims) bool {
	if isAdmin, ok := claims["is_admin"].(bool); ok {
		return isAdmin
	}
	return false
}
