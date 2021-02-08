package handler

import (
	"forwardcall/pkg/service"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"os"
	"path/filepath"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes(mode string) *gin.Engine {
	switch mode {
	case "dev":
		gin.SetMode(gin.DebugMode)
		logrus.Infoln("App loading in debug mode")
	case "release":
		gin.SetMode(gin.ReleaseMode)
		logrus.Infoln("App loading in release mode")
	default:
		logrus.Errorf("Unexpected run mode, please check mode variable in config.")
	}

	router := gin.New()

	router.LoadHTMLFiles(listTemplates("templates/")...)

	// Main page with login panel
	router.GET("/", h.LoginPage)
	// Handler used for authorization
	router.POST("/", h.SignIn)

	schedule := router.Group("/schedule")
	{
		// Get schedule items
		schedule.GET("/", h.GetAllScheduleItems)
		// Create new schedule item
		schedule.POST("/", h.CreateScheduleItem)
		// Get schedule information by ID
		schedule.GET("/:id", h.GetScheduleItemById)
		// Update schedule item information
		schedule.PUT("/:id", h.UpdateScheduleItem)
		// Delete schedule item
		schedule.DELETE("/:id", h.DeleteScheduleItem)
	}

	contact := router.Group("/contact")
	{
		// Get contact list
		contact.GET("/", h.GetAllContacts)
		// Create new contact
		contact.POST("/", h.CreateContact)
		// Get contact information by ID
		contact.GET("/:id", h.GetContactById)
		// Update contact information
		contact.PUT("/:id", h.UpdateContact)
		// Delete contact
		contact.DELETE("/:id", h.DeleteContact)
	}

	api := router.Group("/api")
	{
		// Return next duty engineer
		api.POST("/next_duty", h.GetNextDuty)
	}

	return router
}

func listTemplates(pathTemplate string) []string {
	var files []string

	filepath.Walk(pathTemplate, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() && filepath.Ext(path) == ".html" {
			files = append(files, filepath.ToSlash(path))
		}

		if err != nil {
			return err
		}

		return nil
	})

	return files
}
