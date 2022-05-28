package handler

import (
	"net/http"

	"gocrud/user"

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
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := FormatValidationError(err)
		erroMessage := gin.H{"errors": errors}
		resp := ResponseAPI("Register Account Failed", http.StatusBadRequest, "error", erroMessage)
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	newUser, err := h.userService.RegisterUser(input)
	if err != nil {
		resp := ResponseAPI("Register Account Failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, resp)
		return
	}
	formater := user.FormatUser(newUser, "iniadalahtokenyangbelumterisi")
	resp := ResponseAPI("Account has been registered", http.StatusOK, "success", formater)
	c.JSON(http.StatusOK, resp)
}

func (h *userHandler) GetUsers(c *gin.Context) {
	dataUser, err := h.userService.GetUsers()
	if err != nil {
		return
	}

	c.JSON(http.StatusOK, dataUser)
}
