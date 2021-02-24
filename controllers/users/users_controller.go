package users

import (
	"net/http"
	"strconv"

	"github.com/agusluques/bookstore_users-api/domain/users"
	"github.com/agusluques/bookstore_users-api/services"
	"github.com/agusluques/bookstore_users-api/utils/errors"
	"github.com/gin-gonic/gin"
)

func Get(c *gin.Context) {
	userID, restErr := getUserId(c.Param("user_id"))
	if restErr != nil {
		c.JSON(restErr.Status, restErr)
		return
	}

	result, getErr := services.GetUser(userID)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}

	c.JSON(http.StatusOK, result)
}

func Create(c *gin.Context) {
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

func Update(c *gin.Context) {
	userID, restErr := getUserId(c.Param("user_id"))
	if restErr != nil {
		c.JSON(restErr.Status, restErr)
		return
	}

	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("invalid json object")
		c.JSON(restErr.Status, restErr)
		return
	}

	user.ID = userID

	isPartial := c.Request.Method == http.MethodPatch

	result, updateErr := services.UpdateUser(isPartial, &user)
	if updateErr != nil {
		c.JSON(updateErr.Status, updateErr)
		return
	}

	c.JSON(http.StatusOK, result)
}

func Delete(c *gin.Context) {
	userID, restErr := getUserId(c.Param("user_id"))
	if restErr != nil {
		c.JSON(restErr.Status, restErr)
		return
	}
	if err := services.DeleteUser(userID); err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "deleted"})
}

func Search(c *gin.Context) {
	status := c.Query("status")
	if status == "" {
		c.JSON(http.StatusBadRequest, "status should not be empty")
		return
	}

	result, getErr := services.Search(status)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}

	c.JSON(http.StatusOK, result)
}

func getUserId(userIdParam string) (int64, *errors.RestError) {
	userID, userErr := strconv.ParseInt(userIdParam, 10, 64)
	if userErr != nil {
		return 0, errors.NewBadRequestError("invalid user_id")
	}

	return userID, nil
}
