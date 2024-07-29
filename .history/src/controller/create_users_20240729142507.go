package controller

import (
	"github.com/biduzilla/go-crud/src/configuration/rest_err/validation"
	"github.com/biduzilla/go-crud/src/controller/model/request"
	"github.com/gin-gonic/gin"
)

func CreateUserById(c *gin.Context) {
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		restErr := validation.ValidateUserError(err)

		c.JSON(restErr.Code, restErr)

		return
	}
}
