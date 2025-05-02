package health

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Controller struct{}

func NewController() *Controller {
	return &Controller{}
}

func (c *Controller) Status(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"status":  "ok",
		"version": "1.0",
	})
}

func SetupRoutes(r *gin.RouterGroup, c *Controller) {
	r.GET("/health", c.Status)
}
