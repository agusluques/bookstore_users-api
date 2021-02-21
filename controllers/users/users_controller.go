package users

import (
	"net/http"
	"strconv"

	"github.com/agusluques/bookstore_users-api/domain/users"
	"github.com/agusluques/bookstore_users-api/services"
	"github.com/agusluques/bookstore_users-api/utils/errors"
	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) {
	userId, userErr := strconv.ParseInt(c.Param("user_id"), 10, 64)

	if userErr != nil {
		restErr := errors.NewBadRequestError("invalid user_id")
		c.JSON(restErr.Status, restErr)
		return
	}

	result, getErr := services.GetUser(userId)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}

	c.JSON(http.StatusOK, result)
}

func CreateUser(c *gin.Context) {
	var user users.User

	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("invalid json object")
		c.JSON(restErr.Status, restErr)
		return
	}

	result, saveErr := services.CreateUser(&user)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}

	c.JSON(http.StatusCreated, result)
}

func SearchUser(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{})
}
