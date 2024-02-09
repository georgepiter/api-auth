package controllers

import (
	"api-auth/api/models"
	"api-auth/api/services/auth"
	"api-auth/api/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Login godoc
// @Summary Faz o login do usuário e retorna um token JWT
// @Description Recebe as credenciais do usuário (user name e senha) e retorna um token JWT se as credenciais forem válidas.
// @ID login
// @Accept json
// @Produce json
// @Tags Auth
// @Param input body models.Login true "Credenciais de login do usuário"
// @Success 200 {object} models.AuthResponse "Token JWT gerado"
// @Failure 400 {object} models.ErrorResponse "Erro de requisição inválida"
// @Failure 401 {object} models.ErrorResponse "Credenciais inválidas"
// @Router /api/login [post]
func Login(c *gin.Context) {
	var login models.Login
	if err := c.ShouldBindJSON(&login); err != nil {
		utils.RespondWithError(c.Writer, "Erro de requisição inválida", http.StatusBadRequest)
		return
	}

	token, err := auth.SignIn(login.Email, login.Password)
	if err != nil {
		utils.RespondWithError(c.Writer, "Usuário ou senha inválido", http.StatusUnauthorized)
		return
	}

	authResponse := models.AuthResponse{Token: token}
	c.JSON(http.StatusOK, authResponse)
}

// ValidateToken valida o token de usuário.
// @SecurityDefinitions jwt
// @SecurityScheme jwt
// @in header
// @name Authorization
// @Summary Valida o token de usuário
// @Description Esse endpoint é responsável por validar o token que foi passado nas requisições das API´s
// @ID validate-token
// @Accept json
// @Produce json
// @Tags Auth
// @Param input body models.TokenValidRequest true "Token a ser validado"
// @Success 200 {object} models.TokenValidResponse "Token validado com sucesso"
// @Failure 400 {string} models.ErrorResponse "Erro de requisição inválida"
// @Failure 404 {string} models.ErrorResponse "token inválido"
// @Router /api/validate-token [post]
func ValidateToken(c *gin.Context) {
	var tokenRequest models.TokenValidRequest
	if err := c.ShouldBindJSON(&tokenRequest); err != nil {
		utils.RespondWithError(c.Writer, "Erro de requisição inválida", http.StatusBadRequest)
		return
	}

	tokenResponse, err := auth.ValidateToken(c, tokenRequest)
	if err != nil {
		utils.RespondWithError(c.Writer, "Token inválido", http.StatusUnauthorized)
		return
	}

	c.JSON(http.StatusOK, tokenResponse)
}
