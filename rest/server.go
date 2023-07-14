package rest

import (
	"net/http"

	"github.com/Tonmoy404/hello-world/service"
	"github.com/gin-gonic/gin"
)

type Server struct {
	router *gin.Engine
	svc    service.Service
}

func NewServer(svc service.Service) (*Server, error) {
	server := &Server{
		svc: svc,
	}

	server.setupRouter()

	return server, nil
}

func (server *Server) setupRouter() {
	router := gin.Default()

	router.GET("/api/health-check", server.checkHealth)
	router.GET("/api/students/:id", server.getStudent)

	server.router = router
}

func (server *Server) Start() error {
	return server.router.Run("0.0.0.0:3000")
}

func (server *Server) checkHealth(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "OK")
}
