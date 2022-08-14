package server

import (
	"log"

	"github.com/gin-gonic/gin"

	"github.com/snow-sr/learningGo/server/routes"
)

type Server struct {
	port   string
	server *gin.Engine
}

func NewServer() Server {
	return Server{
		port:   "8087",
		server: gin.Default(),
	}
}

func (s *Server) Run() {
	routes := routes.Config(s.server)

	log.Print("Server sucessfully started")
	log.Fatal(routes.Run("localhost:" + s.port))
}
