package main

import (
	"time"

	"github.com/jproland/pokedexcli/internal/api"
)

type config struct {
	apiClient       api.Client
	nextLocationURL *string
	prevLocationURL *string
	caughtPokemon   map[string]api.Pokemon
}

func main() {
	cfg := config{
		apiClient:     api.NewClient(time.Minute * 5),
		caughtPokemon: make(map[string]api.Pokemon),
	}
	runRepl(&cfg)
}
