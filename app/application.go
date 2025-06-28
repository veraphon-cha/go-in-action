package app

import (
	"go-in-action/config"

	"github.com/gin-gonic/gin"
	"github.com/roonglit/credentials/pkg/credentials"
)

type Server struct {
	Router *gin.Engine
	Config *config.Config
}

func New() *Server {
	router := gin.Default()
	config := loadConfig()

	server := &Server{
		Router: router,
		Config: config,
	}
	return server
}

func (s *Server) Run() {
	s.Router.Run(s.Config.ServerAddress)
}

func loadConfig() *config.Config {
	reader := credentials.NewConfigReader()

	var config config.Config

	if err := reader.Read(gin.Mode(), &config); err != nil {
		panic("Failed to load configuration: " + err.Error())
	}

	return &config
}
