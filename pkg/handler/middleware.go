package handler

import (
	"errors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) userIdentity(c *gin.Context) {
	session := sessions.Default(c)
	username := session.Get("username")
	if username == nil {
		// Abort the request with the appropriate error code
		//c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized from middleware"})
		//c.AbortWithStatus(http.StatusUnauthorized)
		c.AbortWithError(http.StatusUnauthorized, errors.New("User not authorized"))
		return
	}
}
