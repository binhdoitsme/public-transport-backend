package server

import (
	"log"
	"net/http"
	"public-transport-backend/internal/features/identity"
	"public-transport-backend/internal/features/passenger"
	"public-transport-backend/internal/infrastructure/eventhub/passengerhub"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := gin.Default()

	api := r.Group("/api")
	{
		api.GET("/health", s.healthHandler)

		identity.InitMiddlewares(api, s.dependencies)
		identity.InitAPIHandlers(api, s.dependencies)

		passenger.InitAPIHandlers(api, s.dependencies)
	}

	upgrader := &websocket.Upgrader{}
	r.GET("/ws", func(ctx *gin.Context) {
		c, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
		if err != nil {
			log.Print("upgrade:", err)
			return
		}
		defer c.Close()

		hub := s.dependencies.CreateDependenciesFactory().EventPublisher.(*passengerhub.PassengerEventHub)
		id := hub.Subscribe(passengerhub.PassengerCreated, func(data any) {
			c.WriteJSON(data)
		})
		defer hub.Unsubscribe(id)

		for {
			_, _, err := c.ReadMessage()

			if err != nil {
				log.Println("err:", err)
				break
			}
		}
	})

	return r
}

func (s *Server) healthHandler(c *gin.Context) {
	c.AbortWithStatusJSON(http.StatusOK, s.db.Health())
}
