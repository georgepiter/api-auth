package controllers

import (
	"api-auth/api/models"
	"api-auth/api/services"
	"api-auth/api/utils"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// PostUser godoc
// @Summary Cria um novo usuário
// @Description Cria um novo usuário com base nos dados fornecidos
// @ID create-user
// @Accept json
// @Produce json
// @Tags User
// @Param input body models.User true "Credenciais de login do usuário"
// @Success 201 {object} models.User "Usuário criado com sucesso"
// @Failure 400 {string} models.ErrorResponse "Erro de requisição inválida"
// @Failure 404 {string} models.ErrorResponse "Recurso não encontrado"
// @Router /api/create-user [post]
func PostUser(c *gin.Context) {
	request := c.Request
	writer := c.Writer

	bodyParser, err := utils.BodyParser(request)
	if err != nil {
		return
	}

	var user models.User
	if err := json.Unmarshal(bodyParser, &user); err != nil {
		utils.RespondWithError(writer, err.Error(), http.StatusBadRequest)
		return
	}

	result, err := services.CreateUser(user)
	if err != nil {
		utils.RespondWithError(writer, err.Error(), http.StatusNotFound)
		return
	}

	utils.RespondWithJSON(writer, result, http.StatusCreated)
}

// DeleteUser godoc
// @SecurityDefinitions jwt
// @SecurityScheme jwt
// @in header
// @name Authorization
// @Summary Deleta um usuário
// @Description Deleta um usuário pelo ID
// @ID delete-user
// @Accept json
// @Produce json
// @Tags User
// @Param id path int true "ID do usuário a ser deletado"
// @Success 204 {string} string "Usuário deletado com sucesso"
// @Failure 404 {string} models.ErrorResponse "Usuário não encontrado"
// @Failure 500 {string} models.ErrorResponse "Erro interno do servidor"
// @Router /api/delete-user/{id} [delete]
func DeleteUser(c *gin.Context) {
	writer := c.Writer

	userIDStr := c.Param("id")
	userID, err := strconv.ParseInt(userIDStr, 10, 64)
	if err != nil {
		utils.RespondWithError(writer, "ID de usuário inválido", http.StatusBadRequest)
		return
	}

	err = services.DeleteUserByID(userID)
	if err != nil {
		utils.RespondWithError(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	utils.RespondWithJSON(writer, map[string]string{"message": "Usuário excluído com sucesso"}, http.StatusNoContent)
}

// UpdateUser godoc
// @SecurityDefinitions jwt
// @SecurityScheme jwt
// @in header
// @name Authorization
// @Summary Atualiza um usuário
// @Description Atualiza um usuário pelo ID
// @ID update-user
// @Accept json
// @Produce json
// @Tags User
// @Param id path int true "ID do usuário a ser atualizado"
// @Param input body models.User true "Dados do usuário a serem atualizados"
// @Success 200 {object} models.User "Usuário atualizado com sucesso"
// @Failure 400 {string} models.ErrorResponse "Erro de requisição inválida"
// @Failure 404 {string} models.ErrorResponse "Usuário não encontrado"
// @Failure 422 {string} models.ErrorResponse "Entidade não processável"
// @Router /api/update-user/{id} [put]
func UpdateUser(c *gin.Context) {
	request := c.Request
	writer := c.Writer
	idStr := c.Param("id")

	userID, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	user, requestBody, err := getUserByUserID(writer, request, userID)
	if err != nil {
		return
	}

	if len(requestBody) == 0 {
		utils.RespondWithError(writer, "Corpo da requisição vazio", http.StatusBadRequest)
		return
	}

	var updateUser models.User
	if err := json.Unmarshal(requestBody, &updateUser); err != nil {
		utils.RespondWithError(writer, err.Error(), http.StatusUnauthorized)
		return
	}

	userSaved, err := services.UpdateUser(int64(user.ID), updateUser)
	if err != nil {
		utils.RespondWithError(writer, err.Error(), http.StatusUnprocessableEntity)
		return
	}
	utils.ToJson(writer, userSaved, http.StatusOK)
}

func getUserByUserID(writer http.ResponseWriter, request *http.Request, userID int64) (user models.User, requestBody []byte, err error) {
	requestBody, err = utils.BodyParser(request)
	if err != nil {
		utils.RespondWithError(writer, err.Error(), http.StatusBadRequest)
		return user, requestBody, err
	}

	if err := json.Unmarshal(requestBody, &user); err != nil {
		utils.RespondWithError(writer, err.Error(), http.StatusUnauthorized)
		return user, requestBody, err
	}

	userResult := services.GetUserByUserID(userID)
	if userResult.ID == 0 {
		utils.RespondWithError(writer, "Usuário não encontrado", http.StatusNotFound)
		return user, requestBody, nil
	}

	return userResult, requestBody, nil
}
