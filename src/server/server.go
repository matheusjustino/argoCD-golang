package server

import (
	"argoCD-golang/src/routes"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

type Server struct {
	port   string
	server *gin.Engine
}

func NewServer() Server {
	fmt.Println("New Server")
	return Server{
		port:   "8080",
		server: gin.Default(),
	}
}

func (s *Server) Run() {
	router := routes.ConfigRoutes(s.server)

	log.Print("Server is running at port: " + s.port)
	log.Fatal(router.Run(":" + s.port))
}
