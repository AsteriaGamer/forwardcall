package handler

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

func (h *Handler) CreateScheduleItem(c *gin.Context) {
	logrus.Infoln("Call CreateScheduleItem")
	session := sessions.Default(c)
	user := session.Get("username")
	logrus.Infoln("Username = ", user)
	c.JSON(http.StatusOK, gin.H{"user": user})
}

func (h *Handler) GetAllScheduleItems(c *gin.Context) {

}

func (h *Handler) GetScheduleItemById(c *gin.Context) {

}

func (h *Handler) UpdateScheduleItem(c *gin.Context) {

}

func (h *Handler) DeleteScheduleItem(c *gin.Context) {

}
