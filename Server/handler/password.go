package handler

import (
	password "assessment-bootcamp/Server/PassMan"
	"assessment-bootcamp/Server/auth"
	"assessment-bootcamp/Server/entity"
	"net/http"

	"github.com/gin-gonic/gin"
)

type passwordHandler struct {
	passwordService password.PassManService
	authService     auth.Service
}

func NewPasswordhandler(passwordService password.PassManService, authService auth.Service) *passwordHandler {
	return &passwordHandler{passwordService, authService}
}

func (h *passwordHandler) CreatePasswordHandler(c *gin.Context) {
	var inputPassword entity.InputPassword

	if err := c.ShouldBindJSON(&inputPassword); err != nil {
		c.JSON(400, gin.H{
			"error": "input data required",
		})
		return
	}

	userID := int(c.MustGet("currentUser").(int))

	response, err := h.passwordService.SCreateNewPass(userID, inputPassword)

	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(201, response)
}

func (h *passwordHandler) ShowPasswordByIDhandler(c *gin.Context) {
	id := c.Param("id")

	password, err := h.passwordService.SGetPassID(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, password)
}

func (h *passwordHandler) UpdatePasswordByIDHandler(c *gin.Context) {
	id := c.Param("id")

	var updatePasswordInput entity.InputPassword

	if err := c.ShouldBindJSON(&updatePasswordInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err.Error(),
		})
		return
	}

	password, err := h.passwordService.SUpdatePass(id, updatePasswordInput)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	idParam := int(password.UserID)

	userData := int(c.MustGet("currentUser").(int))

	if idParam != userData {
		c.JSON(401, gin.H{
			"error": "unauthorize user",
		})
		return
	}

	c.JSON(http.StatusOK, password)
}

func (h *passwordHandler) DeleteByPasswordIDHandler(c *gin.Context) {
	id := c.Param("id")

	password, _ := h.passwordService.SGetPassID(id)

	idParam := int(password.UserID)

	userData := int(c.MustGet("currentUser").(int))

	if idParam != userData {
		c.JSON(401, gin.H{
			"error": "unauthorize user",
		})
		return
	}

	book, err := h.passwordService.SDelPassID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, book)
}
