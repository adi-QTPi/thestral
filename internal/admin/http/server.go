package http

import (
	"fmt"
	"log"

	"github.com/adi-QTPi/thestral/internal/admin/http/controllers"
	"github.com/adi-QTPi/thestral/internal/admin/http/middlewares"
	"github.com/adi-QTPi/thestral/internal/admin/http/response"
	"github.com/adi-QTPi/thestral/internal/config"
	"github.com/adi-QTPi/thestral/internal/store"
	"github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"
)

func InitServer(cfg *config.Env, store store.Service) {

	r := newRouter()

	responder := response.NewResponder(cfg.DEBUG)
	m := middlewares.NewService(cfg, responder)
	c := controllers.NewService(cfg, responder, store)

	initRoutes(r, m, c)

	log.Println("Admin listening on ", cfg.ADMIN_BIND)
	if err := r.Run(fmt.Sprintf("%v", cfg.ADMIN_BIND)); err != nil {
		fmt.Println("Error Running Admin Server")
	}
}

func newRouter() *gin.Engine {
	r := gin.Default()

	r.SetTrustedProxies(nil)

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"}, //[TODO] tighten allowed origins
		AllowMethods:     []string{"GET", "POST", "PATCH", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	return r
}
