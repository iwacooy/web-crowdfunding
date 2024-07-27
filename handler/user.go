package handler

import (
	"net/http"
	"web-crowdfunding/helper"
	"web-crowdfunding/user"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService user.Service
}

func NewUserHandler(userService user.Service) *userHandler {
	return &userHandler{userService}
}

func (h *userHandler) RegisterUser(c *gin.Context) {
	var input user.RegisterUserInput

	err := c.ShouldBindBodyWithJSON(&input)

	if err != nil {
		error := helper.FormatValidatorError(err)
		errorMsg := gin.H{"error": error}
		response := helper.ResponseAPI("Register Account Failed!", http.StatusUnprocessableEntity, "Failed", errorMsg)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newUser, err := h.userService.RegisterUser(input)
	if err != nil {
		response := helper.ResponseAPI("Register Account Failed!", http.StatusBadRequest, "Failed", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := user.NewUserFormat(newUser)

	response := helper.ResponseAPI("Your Account Has Been Created Sir!", http.StatusOK, "Sukses", formatter)
	c.JSON(http.StatusOK, response)
}
