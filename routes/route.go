package routes

import (
	"OIDC/controller"
	"OIDC/utils"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	e := gin.Default()
	e.Use(cors.Default())
	e.POST("/registration", controller.UserRegister)
	e.POST("/login", controller.UserLogin)

	// Use jwtMiddleware for the following routes
	authRoutes := e.Group("/")
	authRoutes.Use(utils.MiddlewareJWTAuthorize())
	{
		authRoutes.POST("userupdate", controller.UserUpdate)
	}

	return e
}
