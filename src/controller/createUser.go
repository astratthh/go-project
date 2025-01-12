package controller

import (
	"fmt"

	rest_err "github.com/astratthh/first-go-crud.git/src/configuration"
	"github.com/astratthh/first-go-crud.git/src/controller/model/request"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {

	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		restErr := rest_err.NewBadRequestError(
			fmt.Sprintf("Alguns campos estão incorretos, error=%s\n", err.Error()))

		c.JSON(restErr.Code, restErr)
		return
	}

	fmt.Println(userRequest)

}
