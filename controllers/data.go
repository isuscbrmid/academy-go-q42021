package controllers

import (
	"net/http"
	"strconv"

	"go-capstone/common"
	"go-capstone/entities"
	"go-capstone/repositories"

	"github.com/gin-gonic/gin"
)

func ReadCsv(c *gin.Context) {
	idParam := c.Param("id")
	id, parseErr := strconv.ParseInt(idParam, 10, 64)

	if parseErr != nil {
		c.String(common.BadRequestError.HttpCode, common.BadRequestError.Message)
	}

	user, err := repositories.GetEmployee(id)

	if err.HttpCode != 0 {
		c.String(err.HttpCode, err.Message)
	} else {
		c.String(http.StatusOK, "user with email: "+user.Email+" was found.")
	}
}

func AddCsv(c *gin.Context) {
	var users []entities.User
	c.BindJSON(&users)

	_, err := repositories.SaveUsers(users)

	if err.HttpCode != 0 {
		c.String(err.HttpCode, err.Message)
	} else {
		c.String(http.StatusOK, "saved users")
	}
}
