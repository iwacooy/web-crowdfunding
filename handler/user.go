package handler

import (
	"fmt"
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

func (h *userHandler) Login(c *gin.Context) {
	var input user.LoginInput

	err := c.ShouldBindBodyWithJSON(&input)
	if err != nil {
		error := helper.FormatValidatorError(err)
		errorMsg := gin.H{"error": error}
		response := helper.ResponseAPI("Login Gagal!", http.StatusUnprocessableEntity, "Failed", errorMsg)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	loginUser, err := h.userService.Login(input)
	if err != nil {
		response := helper.ResponseAPI("Login Gagal!", http.StatusBadRequest, "Failed", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := user.NewUserFormat(loginUser)

	response := helper.ResponseAPI("Login Sukses!", http.StatusOK, "Sukses", formatter)
	c.JSON(http.StatusOK, response)
}

func (h *userHandler) IsEmailAvailable(c *gin.Context) {
	var input user.CheckEmailAvailableInput

	err := c.ShouldBindBodyWithJSON(&input)
	if err != nil {
		error := helper.FormatValidatorError(err)
		errorMsg := gin.H{"error": error}
		response := helper.ResponseAPI("Format email tidak sesuai", http.StatusUnprocessableEntity, "error", errorMsg)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	isAvailable, err := h.userService.IsEmailAvailable(input)
	if err != nil {
		errorMsg := gin.H{"error": err.Error()}
		response := helper.ResponseAPI("Checking email failed", http.StatusBadRequest, "error", errorMsg)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	var metaMsg string

	data := gin.H{
		"is_available": isAvailable,
	}

	if isAvailable {
		metaMsg = "Email Available!"
	} else {
		metaMsg = "Email has been registered!"
	}

	response := helper.ResponseAPI(metaMsg, http.StatusOK, "Sukses", data)
	c.JSON(http.StatusOK, response)

}

func (h *userHandler) UploadAvatar(c *gin.Context) {
	file, err := c.FormFile("avatar")
	if err != nil {
		if err != nil {
			errorMsg := gin.H{"error": err.Error()}
			response := helper.ResponseAPI("Upload Avatar Failed", http.StatusBadRequest, "failed", errorMsg)
			c.JSON(http.StatusBadRequest, response)
			return
		}
	}

	userID := 1
	path := fmt.Sprintf("images/%d-%s", userID, file.Filename)
	err = c.SaveUploadedFile(file, path)
	if err != nil {
		errorMsg := gin.H{"error": err.Error()}
		response := helper.ResponseAPI("Upload Avatar Failed", http.StatusBadRequest, "failed", errorMsg)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	_, err = h.userService.SaveAvatar(1, path)
	if err != nil {
		errorMsg := gin.H{"error": err.Error()}
		response := helper.ResponseAPI("Upload Avatar Failed", http.StatusBadRequest, "failed", errorMsg)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	msg := gin.H{"is_uploaded": true}
	response := helper.ResponseAPI("Avatar successfully uploaded", http.StatusOK, "Sukses", msg)
	c.JSON(http.StatusOK, response)

}
