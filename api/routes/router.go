package routes

import (
	"api-auth/api/controllers"
	"api-auth/api/middlewares"
	"api-auth/api/services"
	_ "api-auth/docs"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/", controllers.PublicRoute)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.POST("/api/login", controllers.Login)
	r.POST("/api/create-user", controllers.PostUser)

	r.Use(func(c *gin.Context) {
		middlewares.IsAuth(func(c *gin.Context) {
			c.Next()
		})(c)
	})

	r.PUT("/api/update-user/:id", isAdminHandler(controllers.UpdateUser))
	r.DELETE("/api/delete-user/:id", isAdminHandler(controllers.DeleteUser))
	r.POST("/api/validate-token", isAdminHandler(controllers.ValidateToken))

	return r
}

func isAdminHandler(handler gin.HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		if !services.IsAdmin(c) {
			c.JSON(http.StatusForbidden, gin.H{"error": "Acesso negado"})
			return
		}
		handler(c)
	}
}
