package controller

import (
	"fmt"

	"github.com/biduzilla/go-crud/src/configuration/rest_err"
	"github.com/biduzilla/go-crud/src/controller/model/request"
	"github.com/gin-gonic/gin"
)

func CreateUserById(c *gin.Context) {
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		restErr := rest_err.NewBadRequestError(
			fmt.Sprintf("Incorrect fields", err),
		)

		c.JSON(restErr.Code, restErr)

		return
	}
}
