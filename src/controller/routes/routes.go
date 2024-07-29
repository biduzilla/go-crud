package routes

import (
	"github.com/biduzilla/go-crud/src/controller"
	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.RouterGroup) {
	r.GET("/user/:userId", controller.FindUserById)
	r.GET("/email/:userEmail", controller.FindUserByEmail)
	r.POST("/user", controller.CreateUserById)
	r.PUT("/user", controller.UpdateUserById)
	r.DELETE("/user/:userId", controller.DeleteUserById)
}
