package middlewares

import (
	"api-auth/config"
	"fmt"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

var secretKey = config.JwtSecretKey

func IsAuth(endpoint func(c *gin.Context)) gin.HandlerFunc {
	return func(c *gin.Context) {
		authorizationHeader := c.GetHeader("Authorization")
		if authorizationHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Não autorizado: Cabeçalho de Autorização ausente"})
			return
		}

		bearerToken := strings.Split(authorizationHeader, " ")
		if len(bearerToken) != 2 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Não autorizado: Formato inválido no cabeçalho de Autorização"})
			return
		}

		token, err := jwt.Parse(bearerToken[1], func(token *jwt.Token) (interface{}, error) {
			return getKey(token)
		})
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": fmt.Sprintf("Não autorizado: %s", err.Error())})
			return
		}

		if token.Valid {
			endpoint(c)
		} else {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Não autorizado: Token inválido"})
		}
	}
}

func getKey(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(jwt.SigningMethod); !ok {
		return nil, fmt.Errorf("Não autorizado")
	}
	return secretKey, nil
}
