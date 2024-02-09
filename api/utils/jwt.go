package utils

import (
	"api-auth/api/models"
	"api-auth/config"
	"fmt"
	jwt "github.com/dgrijalva/jwt-go"
	"html"
	"net/http"
	"strings"
	"time"
)

var secretKey = config.JwtSecretKey

func GenerateJWT(user models.User) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["authorized"] = true
	claims["user_email"] = user.Email
	claims["rule"] = user.Rule
	claims["exp"] = time.Now().Add(50 * time.Minute).Unix()

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", fmt.Errorf("Erro ao assinar o token: %v", err)
	}

	tokenString = "Bearer " + tokenString

	return tokenString, nil
}

func JwtExtract(r *http.Request) (map[string]interface{}, error) {
	tokenString := ExtractBearerToken(r)
	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
	if err != nil {
		return nil, fmt.Errorf("Erro ao analisar o token: %v", err)
	}
	if !token.Valid {
		return nil, fmt.Errorf("Token inválido")
	}
	return claims, nil
}

func ExtractBearerToken(r *http.Request) string {
	headerAuthorization := r.Header.Get("Authorization")
	if headerAuthorization == "" {
		return ""
	}

	bearerToken := strings.Split(headerAuthorization, " ")
	if len(bearerToken) != 2 {
		return ""
	}

	return html.EscapeString(bearerToken[1])
}
