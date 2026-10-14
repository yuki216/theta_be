package http

import (
	"github.com/gin-gonic/gin"
	"theta-be/delivery/http/healthhandler"
	"theta-be/delivery/http/middleware"
	"theta-be/delivery/http/recordhandler"
)

type RouterConfig struct {
	HealthHandler *healthhandler.Handler
	RecordHandler *recordhandler.Handler
}

func SetupRouter(cfg RouterConfig) *gin.Engine {
	r := gin.New()

	r.Use(middleware.Logger())
	r.Use(middleware.Recovery())
	r.Use(middleware.CORS())

	// Health Route
	r.GET("/healthz", cfg.HealthHandler.HealthCheck)

	// API v1 Routes
	v1 := r.Group("/api/v1")
	{
		records := v1.Group("/records")
		{
			records.POST("", cfg.RecordHandler.CreateRecord)
			records.GET("", cfg.RecordHandler.ListRecords)
			records.GET("/:id", cfg.RecordHandler.GetRecord)
			records.DELETE("/:id", cfg.RecordHandler.DeleteRecord)
			records.PATCH("/:id/activate", cfg.RecordHandler.ActivateRecord)
		}
	}

	return r
}
