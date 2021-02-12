package handler

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

func (h *Handler) userIdentity(c *gin.Context) {
	logrus.Infoln("Exec middleware")
	session := sessions.Default(c)
	username := session.Get("username")
	if username == nil {
		// Abort the request with the appropriate error code
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	logrus.Infoln("Middleware call next handler")
	// Continue down the chain to handler etc
	c.Next()
}
