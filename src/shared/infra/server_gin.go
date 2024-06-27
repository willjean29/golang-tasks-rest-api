package server

import (
	db "app/src/shared/data/gorm"
	"app/src/shared/infra/http/routes"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ServerGin struct {
	port   string
	router *gin.Engine
}

func CreateServer() *ServerGin {
	var defaultPort int64 = 4000
	port := strconv.FormatInt(defaultPort, 10)
	router := gin.Default()
	server := &ServerGin{
		port:   port,
		router: router,
	}
	server.DBConnectionGin()
	server.RoutesGin()
	return server
}

func (s *ServerGin) DBConnectionGin() {
	// connection database
	db.GormSyncDatabase()
}

func (s *ServerGin) RoutesGin() {
	s.router.Static("/api/uploads", "./uploads")
	routes.RoutesGin(s.router)
}

func (s *ServerGin) Run() {
	log.Println("Running on port ", s.port)
	log.Fatal(s.router.Run(":" + s.port))
}
