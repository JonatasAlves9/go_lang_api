package api

import (
	"tasks_api/src/api/routers"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"tasks_api/src/api/middlewares"
	"tasks_api/src/config"
)

func InitServer(cfg *config.Config) {
	gin.SetMode(cfg.Server.RunMode)
	r := gin.New()

	r.Use(middlewares.Cors(cfg))
	RegisterRoutes(r, cfg)
	r.Run()

}

func RegisterRoutes(r *gin.Engine, cfg *config.Config) {
	api := r.Group("/api")

	v1 := api.Group("/v1")
	{

		// User
		users := v1.Group("/users")

		// Test
		routers.User(users)

		r.Static("/static", "./uploads")

		r.GET("/metrics", gin.WrapH(promhttp.Handler()))
	}

}
