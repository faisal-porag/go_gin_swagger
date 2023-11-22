package state

import (
	"go_gin_swagger/config"
)

type State struct {
	Cfg *config.Config
}

func NewState(cfg *config.Config) *State {
	return &State{
		Cfg: cfg,
	}
}
