package controllers

import (
	"net/http"
	"strconv"

	"go-capstone/common"
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
