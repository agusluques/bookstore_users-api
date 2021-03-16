package users

import (
	"net/http"
	"strconv"

	"github.com/agusluques/bookstore_oauth-go/oauth"
	"github.com/agusluques/bookstore_users-api/domain/users"
	"github.com/agusluques/bookstore_users-api/services"
	"github.com/agusluques/bookstore_utils-go/rest_errors"
	"github.com/gin-gonic/gin"
)

// Get an user
func Get(c *gin.Context) {
	if err := oauth.AuthenticateRequest(c.Request); err != nil {
		c.JSON(err.Status, err)
		return
	}

	userID, restErr := getUserID(c.Param("user_id"))
	if restErr != nil {
		c.JSON(restErr.Status, restErr)
		return
	}

	result, getErr := services.UsersService.Get(userID)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}

	if oauth.GetCallerId(c.Request) == result.ID {
		c.JSON(http.StatusOK, result.Marshall(false))
		return
	}

	c.JSON(http.StatusOK, result.Marshall(oauth.IsPublic(c.Request)))
}

// Create an user
func Create(c *gin.Context) {
	var user users.User

	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := rest_errors.NewBadRequestError("invalid json object")
		c.JSON(restErr.Status, restErr)
		return
	}

	result, saveErr := services.UsersService.Create(&user)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}

	c.JSON(http.StatusCreated, result.Marshall(c.GetHeader("X-Public") == "true"))
}

// Update an user
func Update(c *gin.Context) {
	userID, restErr := getUserID(c.Param("user_id"))
	if restErr != nil {
		c.JSON(restErr.Status, restErr)
		return
	}

	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := rest_errors.NewBadRequestError("invalid json object")
		c.JSON(restErr.Status, restErr)
		return
	}

	user.ID = userID

	isPartial := c.Request.Method == http.MethodPatch

	result, updateErr := services.UsersService.Update(isPartial, &user)
	if updateErr != nil {
		c.JSON(updateErr.Status, updateErr)
		return
	}

	c.JSON(http.StatusOK, result.Marshall(c.GetHeader("X-Public") == "true"))
}

// Delete an user
func Delete(c *gin.Context) {
	userID, restErr := getUserID(c.Param("user_id"))
	if restErr != nil {
		c.JSON(restErr.Status, restErr)
		return
	}
	if err := services.UsersService.DeleteUser(userID); err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "deleted"})
}

// Search an user
func Search(c *gin.Context) {
	status := c.Query("status")
	if status == "" {
		c.JSON(http.StatusBadRequest, "status should not be empty")
		return
	}

	results, getErr := services.UsersService.Search(status)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}

	c.JSON(http.StatusOK, results.Marshall(c.GetHeader("X-Public") == "true"))
}

// Login an user
func Login(c *gin.Context) {
	var request users.LoginRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		restErr := rest_errors.NewBadRequestError("invalid json object")
		c.JSON(restErr.Status, restErr)
		return
	}

	result, err := services.UsersService.LoginUser(request)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusOK, result.Marshall(c.GetHeader("X-Public") == "true"))
}

func getUserID(userIDParam string) (int64, *rest_errors.RestError) {
	userID, userErr := strconv.ParseInt(userIDParam, 10, 64)
	if userErr != nil {
		return 0, rest_errors.NewBadRequestError("invalid user_id")
	}

	return userID, nil
}
