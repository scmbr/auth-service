package v1

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}
func (h *Handler) Init(api *gin.RouterGroup) {
	v1 := api.Group("/v1")
	{
		v1.GET("/system/health", func(c *gin.Context) {
			status := http.StatusOK
			c.JSON(status, gin.H{
				"message": "Hello World",
				"service": "ok",
				"time":    time.Now(),
			})
		})

	}
}
