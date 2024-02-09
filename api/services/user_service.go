package services

import (
	"api-auth/api/models"
	"api-auth/api/security"
	"api-auth/config"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"strings"
	"time"
)

var secretKey = config.JwtSecretKey

func CreateUser(user models.User) (interface{}, error) {
	db := models.Connect()
	defer func(db *gorm.DB) {
		err := db.Close()
		if err != nil {

		}
	}(db)

	var u models.User
	db.Where("user_name = ?", user.UserName).Find(&u)

	if len(u.UserName) != 0 {
		return nil, errors.New("usuário já cadastrado no sistema")
	}

	hashedPassword, err := security.Hash(user.Password)
	if err != nil {
		return nil, err
	}

	user.Password = string(hashedPassword)
	user.Rule = "USER"
	userCreated := db.Create(&user)
	createdID := userCreated.Value.(*models.User).ID

	userWithoutPassword := models.UserWithoutPassword{
		ID:        createdID,
		UserName:  user.UserName,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	return userWithoutPassword, userCreated.Error
}

func GetUserByEmail(email string) models.User {
	db := models.Connect()
	defer func(db *gorm.DB) {
		err := db.Close()
		if err != nil {
		}
	}(db)

	var user models.User
	db.Where("email = ?", email).Find(&user)

	return user
}

func DeleteUserByID(userID int64) error {
	db := models.Connect()
	defer func(db *gorm.DB) {
		err := db.Close()
		if err != nil {
		}
	}(db)

	var user models.User
	result := db.Where("id = ?", userID).Delete(&user)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("O Usuário não foi encontrado")
	}

	return nil
}

func UpdateUser(userID int64, newUser models.User) (models.User, error) {
	db := models.Connect()
	defer func(db *gorm.DB) {
		err := db.Close()
		if err != nil {

		}
	}(db)

	var existingUser models.User
	if err := db.First(&existingUser, "id = ?", userID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.User{}, err
		}

		return models.User{}, err
	}
	now := time.Now()
	existingUser.UserName = newUser.UserName
	existingUser.Email = newUser.Email
	existingUser.Password = newUser.Password
	existingUser.UpdatedAt = &now

	if err := db.Save(&existingUser).Error; err != nil {
		return models.User{}, err
	}

	return existingUser, nil
}

func GetUserByUserID(userID int64) models.User {
	db := models.Connect()
	defer func(db *gorm.DB) {
		err := db.Close()
		if err != nil {
		}
	}(db)

	var user models.User
	db.Where("id = ?", userID).Find(&user)

	return user
}

func IsAdmin(c *gin.Context) bool {
	authorizationHeader := c.GetHeader("Authorization")
	if authorizationHeader == "" {
		return false
	}

	token := strings.Replace(authorizationHeader, "Bearer ", "", 1)

	email, rule, err := getTokenClaims(token)
	if err != nil {
		return false
	}

	user := GetUserByEmail(email)

	if user.ID == 0 {
		return false
	}

	if user.Rule != rule {
		return false
	}

	if user.Rule != "ADMIN" {
		return false
	}

	return true
}

func getTokenClaims(tokenString string) (email, rule string, err error) {
	tokenString = strings.Replace(tokenString, "Bearer ", "", 1)

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
	if err != nil {
		return email, rule, err
	}

	if !token.Valid {
		return email, rule, fmt.Errorf("Token inválido")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return email, rule, fmt.Errorf("Erro ao obter claims do token")
	}

	email, ok = claims["user_email"].(string)
	if !ok {
		return email, rule, fmt.Errorf("Erro ao obter email do token")
	}

	rule, ok = claims["rule"].(string)
	if !ok {
		return email, rule, fmt.Errorf("Erro ao obter rule do token")
	}

	return email, rule, nil
}
