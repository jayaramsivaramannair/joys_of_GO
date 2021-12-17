package routes


import (
	controller "github.com/jayaramsivaramannair/joys_of_GO/Project_JWT-Authentication/controllers"
	"github.com/gin-gonic/gin"
)

func AuthRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("users/signup", controller.Signup())
	incomingRoutes.POST("users/login", controller.Login())
}