package main

import (
	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
	"go_gin_swagger/config"
	_ "go_gin_swagger/docs"
	"go_gin_swagger/httpserver"
	"go_gin_swagger/state"
)

// @title MY Sample API
// @version 1.0
// @description Sample API description
// @BasePath /api/v1
func main() {
	envLoadErr := godotenv.Load()
	if envLoadErr != nil {
		log.Fatal().Err(envLoadErr).Msg("env load error")
	}

	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatal().Err(err).Msg("Config parsing failed")
	}

	appState := state.NewState(cfg)
	httpserver.Serve(appState)
}
