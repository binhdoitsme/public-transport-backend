package server

import (
	"net/http"
	"public-transport-backend/internal/features/identity"
	"public-transport-backend/internal/features/passenger"

	"github.com/gin-gonic/gin"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := gin.Default()

	api := r.Group("/api")
	{
		identity.InitMiddlewares(api, s.dependencies)

		api.GET("/", s.HelloWorldHandler)
		api.GET("/health", s.healthHandler)

		identity.InitAPIHandlers(api, s.dependencies)
		passenger.InitAPIHandlers(api, s.dependencies)
	}

	return r
}

func (s *Server) HelloWorldHandler(c *gin.Context) {
	resp := make(map[string]string)
	resp["message"] = "Hello World"

	c.JSON(http.StatusOK, resp)
}

func (s *Server) healthHandler(c *gin.Context) {
	c.JSON(http.StatusOK, s.db.Health())
}
