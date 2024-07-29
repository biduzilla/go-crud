package routes

import (
	controller "github.com/biduzilla/go-crud/src/controller/routes"
	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.RouterGroup) {
	r.GET("/user/:userId", controller.FindUserById)
	r.GET("/email/:userEmail", controller.FindUserByEmail)
	r.POST("/user", controller.CreateUser)
	r.PUT("/user", controller.UpdateUser)
	r.DELETE("/user/:userId", controller.DeleteUser)
}
