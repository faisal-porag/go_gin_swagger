package httpserver

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"go_gin_swagger/router_manager"
	"go_gin_swagger/state"
)

func Serve(s *state.State) {
	// TODO:: UNCOMMENT THIS LINE IN PRODUCTION
	//gin.SetMode(gin.ReleaseMode)

	// Create a Gin router instance
	router := gin.Default()

	// CORS middleware
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE"}
	config.AllowHeaders = []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token", "Accept-Language"}
	router.Use(cors.New(config))

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	v1 := router.Group("/api/v1")
	{
		router_manager.MountCustomerRoutes(v1.Group("/customer"))
	}

	// Static files
	router.Static("/uploads", "./uploads")

	// Start the Gin server
	port := s.Cfg.ApplicationPort
	log.Info().Int("port", port).Msg("Starting customer app service http server")
	err := router.Run(fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatal().Err(err).Msg("router.Run err")
	}
}
