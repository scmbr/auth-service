package handler

import (
	"github.com/gin-gonic/gin"

	v1 "github.com/scmbr/auth-service/internal/delivery/http/handler/v1"
)

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) Init() *gin.Engine {
	router := gin.Default()
	router.Use(gin.Recovery(), gin.Logger())
	h.initAPI(router)
	return router
}
func (h *Handler) initAPI(router *gin.Engine) {
	handlerV1 := v1.NewHandler()
	api := router.Group("/api")
	{
		handlerV1.Init(api)
	}
}
