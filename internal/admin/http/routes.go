package http

import (
	"github.com/adi-QTPi/thestral/internal/admin/http/controllers"
	"github.com/adi-QTPi/thestral/internal/admin/http/middlewares"
	"github.com/gin-gonic/gin"
)

func initRoutes(router *gin.Engine, m *middlewares.Service, c *controllers.Service) {
	// [TODO] admin routes for crud

	route := router.Group("/proxy")
	{
		route.POST("", c.CreateProxy)
		route.DELETE("", c.DeleteProxy)
	}
}
