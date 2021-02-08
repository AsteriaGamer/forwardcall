package handler

import (
	"forwardcall/pkg/entity"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) SignIn(c *gin.Context) {
	var input entity.User

	if err := c.Bind(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	// Вызов метода из сервиса
	h.services.Authorization.LoginUser(input)
}

func (h *Handler) LoginPage(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"msg": "Hello World",
	})
}
