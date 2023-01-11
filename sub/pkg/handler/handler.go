package handler

import (
	"github.com/gin-gonic/gin"
	"nats/pkg/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	router.Static("/static", "/Users/nishaque/GolandProjects/nats/sub/static")
	router.LoadHTMLGlob("html/*.html")

	api := router.Group("/order")
	{

		api.POST("/", h.get)
		api.GET("/", h.mainPage)

	}
	return router
}
