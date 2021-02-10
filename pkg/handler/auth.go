package handler

import (
	"forwardcall/pkg/entity"
	"forwardcall/pkg/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

func (h *Handler) SignIn(c *gin.Context) {
	var input entity.User

	if err := c.ShouldBind(&input); err != nil {
		//input.GetError(err.(validator.ValidationErrors))
		c.String(http.StatusBadRequest, "Error message: %s", err.(validator.ValidationErrors).Translate(utils.Trans))
		return
	}

	if input.Username != "root" || input.Password != "q1q2q3" {
		c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})

	// Вызов метода из сервиса
	//h.services.Authorization.LoginUser(input)
}

func (h *Handler) LoginPage(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"title": "ForwardCall",
	})
}
