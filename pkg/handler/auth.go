package handler

import (
	"forwardcall/pkg/entity"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) SignIn(c *gin.Context) {
	session := sessions.Default(c)

	// Структура entity.User принимает и валидирует входные данные
	var input entity.User

	// Передача данных из запроса в структуру User и их валидация
	if err := c.ShouldBind(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "bad request"})
		return
	}

	// Аутентификация пользователя
	username, err := h.services.Authorization.AuthUser(input.Username, input.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
		return
	}

	session.Set("username", username)
	session.Set("authorization", true)

	if err = session.Save(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save session"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "you are logged in",
	})
}

func (h *Handler) SignOut(c *gin.Context) {
	session := sessions.Default(c)

	username := session.Get("username")
	if username == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid session token"})
		return
	}
	session.Delete("username")
	if err := session.Save(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save session"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Successfully logged out"})
}
