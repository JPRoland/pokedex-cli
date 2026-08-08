package main

import (
	"fmt"
	"math/rand"
	"os"
)

func commandExit(cfg *config, args ...string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(cfg *config, args ...string) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	commands := getCommands()

	for _, cmd := range commands {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	fmt.Println("")
	return nil
}

func commandMap(cfg *config, args ...string) error {
	//if locationURL is nil we will fetch the first page, otherwise we fetch the page from the config
	res, err := cfg.apiClient.GetLocationAreas(cfg.nextLocationURL)
	if err != nil {
		return err
	}

	fmt.Println("Locations:")
	for _, locationArea := range res.Results {
		fmt.Printf("- %s\n", locationArea.Name)
	}
	cfg.prevLocationURL = res.Previous
	cfg.nextLocationURL = res.Next
	return nil
}

func commandMapb(cfg *config, args ...string) error {
	if cfg.prevLocationURL == nil {
		return fmt.Errorf("No previous page")
	}

	res, err := cfg.apiClient.GetLocationAreas(cfg.prevLocationURL)
	if err != nil {
		return err
	}

	fmt.Println("Locations:")
	for _, locationArea := range res.Results {
		fmt.Printf("- %s\n", locationArea.Name)
	}
	cfg.prevLocationURL = res.Previous
	cfg.nextLocationURL = res.Next
	return nil
}

func commandExplore(cfg *config, args ...string) error {
	fmt.Println("args: ", args)
	if len(args) != 1 {
		return fmt.Errorf("expected 1 argument (location area name)")
	}

	locationAreaName := args[0]

	locationArea, err := cfg.apiClient.GetLocationArea(locationAreaName)
	if err != nil {
		return err
	}

	fmt.Printf("Pokemon in : %s\n", locationAreaName)
	for _, pokemon := range locationArea.PokemonEncounters {
		fmt.Printf("- %s\n", pokemon.Pokemon.Name)
	}
	return nil
}

func commandCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return fmt.Errorf("expected 1 argument (pokemon name)")
	}

	pokemonName := args[0]

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonName)

	pokemon, err := cfg.apiClient.GetPokemon(pokemonName)
	if err != nil {
		return err
	}

	const threshold = 50
	randNum := rand.Intn(pokemon.BaseExperience)
	if randNum > threshold {
		return fmt.Errorf("failed to catch %s", pokemonName)
	}
	fmt.Printf("%s was caught!\n", pokemonName)
	fmt.Println("You may now inspect it with the \"inspect\" command.")
	cfg.caughtPokemon[pokemonName] = pokemon
	return nil
}

func commandInspect(cfg *config, args ...string) error {
	if len(args) != 1 {
		return fmt.Errorf("expected 1 argument (pokemon name)")
	}

	pokemonName := args[0]

	pokemon, ok := cfg.caughtPokemon[pokemonName]
	if !ok {
		return fmt.Errorf("pokemon %s not found in your Pokedex", pokemonName)
	}

	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)

	fmt.Println("Stats: ")
	for _, stat := range pokemon.Stats {
		fmt.Printf("- %s: %v\n", stat.Stat.Name, stat.BaseStat)
	}

	fmt.Println("Types:")
	for _, typeItem := range pokemon.Types {
		fmt.Printf("- %s\n", typeItem.Type.Name)
	}

	fmt.Println("Abilities: ")
	for _, ability := range pokemon.Abilities {
		fmt.Printf("- %s\n", ability.Ability.Name)
	}

	return nil
}

func commandPokedex(cfg *config, args ...string) error {
	if len(cfg.caughtPokemon) == 0 {
		fmt.Println("Your Pokedex is empty.")
		return nil
	}
	fmt.Println("Your Pokedex:")
	for pokemonName := range cfg.caughtPokemon {
		fmt.Printf("- %s\n", pokemonName)
	}
	return nil
}
