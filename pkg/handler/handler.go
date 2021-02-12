package handler

import (
	"forwardcall/pkg/service"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/memstore"
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

func (h *Handler) InitRoutes(mode string, sessionKey string) *gin.Engine {

	// Выбор режима запуска приложения
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

	// Создание нового экземпляра фреймворка
	router := gin.New()

	// Регистрация пути "/assets" в качестве стандартного URL'a для файлов frontend'a
	router.Static("/assets", "./public/assets")
	// Парсинг html файлов шаблонов
	router.LoadHTMLFiles(listTemplates("public/view")...)

	// Инициализация менеджера сессий.
	sessionStorage := memstore.NewStore([]byte(sessionKey))
	router.Use(sessions.Sessions("session", sessionStorage))

	// Main page with login panel
	router.GET("/", h.MainPage)

	auth := router.Group("/auth")
	{
		auth.POST("/sign-in", h.SignIn)
		auth.POST("/sign-out", h.SignOut)
	}

	schedule := router.Group("/schedule", h.userIdentity)
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
