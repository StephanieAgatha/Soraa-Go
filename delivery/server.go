package delivery

import (
	"fmt"
	"github.com/StephanieAgatha/Soraa-Go/config"
	"github.com/StephanieAgatha/Soraa-Go/delivery/controller"
	"github.com/StephanieAgatha/Soraa-Go/delivery/middleware"
	"github.com/StephanieAgatha/Soraa-Go/manager"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Server struct {
	um   manager.UsecaseManager
	gin  *gin.Engine
	host string
}

// middleware goes here
func (s *Server) InitMiddleware() {
	// Create a Zap logger
	logger, _ := zap.NewProduction()
	defer logger.Sync() // Ensure logs are flushed

	// Use the logger
	s.gin.Use(middleware.ZapLogger(logger))
}

// controller
func (s *Server) InitController() {
	controller.NewUserController(s.um.UserUsecase(), s.gin).Route()
	controller.NewUserCredentialController(s.um.UserCredUsecase(), s.gin).Route()
}

// run server
func (s *Server) Run() {
	s.InitMiddleware()
	s.InitController()
	err := s.gin.Run(s.host)
	if err != nil {
		fmt.Printf("Failed to run server %v ", err.Error())
	}
}

func NewServer() *Server {
	//define contrusctor from config
	cfg, err := config.NewDbConfig()
	if err != nil {
		fmt.Printf("Failed on config server %v", err.Error())
	}

	//constructor from infra
	im, err := manager.NewInfraManager(cfg)
	if err != nil {
		fmt.Printf("Failed on construct infra %v", err.Error())
	}

	//constructor from repomanager
	rm := manager.NewRepoManager(im)
	//contructor from usecase manager
	um := manager.NewUsecaseManager(rm)

	//set host for gin server
	host := fmt.Sprintf("%s:%s", cfg.ApiConfig.Host, cfg.ApiConfig.Port)
	//return gin instance
	g := gin.Default()
	return &Server{
		um:   um,
		gin:  g,
		host: host,
	}
}
