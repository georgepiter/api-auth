package auth

import (
	"api-auth/api/models"
	"api-auth/api/security"
	"api-auth/api/services"
	"api-auth/api/utils"
	"api-auth/config"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
	"strings"
	"time"
)

var (
	ErrUserNotFound = errors.New("Usuário não encontrado")
)

var secretKey = config.JwtSecretKey

func SignIn(email, password string) (string, error) {
	user := services.GetUserByEmail(email)
	if user.ID == 0 {
		return "", ErrUserNotFound
	}
	err := security.CheckPassword(user.Password, password)
	if err != nil {
		return "", err
	}
	token, err := utils.GenerateJWT(user)
	if err != nil {
		return "", err
	}

	err = UpdateAuditLogin(user.ID)
	if err != nil {
		return "Erro ao atualizar o horário de login", err
	}

	return token, nil
}

func UpdateAuditLogin(userID uint32) error {
	db := models.Connect()
	defer func(db *gorm.DB) {
		err := db.Close()
		if err != nil {

		}
	}(db)

	user := models.User{}
	if err := db.Model(&user).Where("id = ?", userID).Update("audit_login", time.Now()).Error; err != nil {
		return err
	}

	return nil
}

func ValidateToken(c *gin.Context, tokenRequest models.TokenValidRequest) (tokenResponse models.TokenValidResponse, err error) {
	if tokenRequest.Token == "" || !strings.HasPrefix(tokenRequest.Token, "Bearer ") {
		c.JSON(http.StatusUnauthorized, models.ErrorResponse{Message: "Token inválido"})
		return
	}

	token := strings.TrimPrefix(tokenRequest.Token, "Bearer ")

	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return getKey(token)
	})
	if err != nil || !parsedToken.Valid {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Não autorizado: Token inválido"})
		return
	}

	tokenResponse.Token = tokenRequest.Token
	tokenResponse.IsValid = true

	return tokenResponse, nil
}

func getKey(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(jwt.SigningMethod); !ok {
		return nil, fmt.Errorf("Não autorizado")
	}
	return secretKey, nil
}
