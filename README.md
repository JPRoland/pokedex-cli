# Pokédex CLI

A feature-rich command-line Pokédex application written in Go. Interact with the PokéAPI to explore location areas, catch wild Pokémon, view detailed stats, and manage your personal Pokédex from your terminal.

![Go Version](https://img.shields.io/badge/Go-1.25%2B-00ADD8?style=flat&logo=go)
![License](https://img.shields.io/badge/license-MIT-blue.svg)

---

## Features

- 🗺️ **World Exploration**: Paginate through location areas across the Pokémon world (`map` & `mapb`).
- 🔍 **Area Inspection**: Explore specific location areas to see which wild Pokémon inhabit them (`explore <area>`).
- ⚾ **Catch Mechanics**: Attempt to catch Pokémon with catch rates dynamically based on base experience (`catch <pokemon>`).
- 📖 **Pokédex Management**: Keep track of all Pokémon you have successfully caught (`pokedex`).
- 📊 **Detailed Stats Inspection**: View height, weight, base stats, types, and abilities for any caught Pokémon (`inspect <pokemon>`).
- ⚡ **In-Memory Caching**: Custom thread-safe cache (`pokecache`) with automatic background TTL eviction to minimize redundant PokéAPI network calls.

---

## Project Structure

```text
pokedex/
├── main.go                     # Entry point & global state configuration
├── repl.go                     # Interactive REPL implementation & command routing
├── repl_test.go                # Unit tests for input cleaning logic
├── commands.go                 # Implementation of CLI command callbacks
├── internal/
│   ├── api/                    # PokéAPI client implementation & data structures
│   └── pokecache/              # Thread-safe in-memory cache with background TTL reaping
└── go.mod                      # Go module definition
```

---

## Installation & Setup

### Prerequisites

- [Go](https://go.dev/doc/install) 1.25 or higher installed on your system.

### Building from Source

1. Clone the repository:
   ```bash
   git clone https://github.com/jproland/pokedexcli.git
   cd pokedexcli
   ```

2. Build the binary:
   ```bash
   go build -o pokedexcli .
   ```

3. Run the application:
   ```bash
   ./pokedexcli
   ```

Alternatively, you can run the project directly without building an executable:
```bash
go run .
```

---

## Usage

When launched, the application enters an interactive REPL prompt: `Pokedex > `.

### Available Commands

| Command | Description | Example Usage |
| :--- | :--- | :--- |
| `help` | Displays a list of all available commands and descriptions. | `help` |
| `map` | Displays the names of the next 20 location areas in the Pokémon world. | `map` |
| `mapb` | Displays the names of the previous 20 location areas. | `mapb` |
| `explore <area_name>` | Lists all wild Pokémon encounters in a given location area. | `explore canalave-city-area` |
| `catch <pokemon_name>` | Throws a Pokeball to attempt catching a Pokémon. Harder for higher base XP Pokémon. | `catch pikachu` |
| `inspect <pokemon_name>` | View height, weight, stats, types, and abilities of a caught Pokémon. | `inspect pikachu` |
| `pokedex` | Lists all Pokémon currently caught and stored in your Pokédex. | `pokedex` |
| `exit` | Exits the Pokédex CLI application. | `exit` |

---

## Running Tests

To run the unit tests across all packages:

```bash
go test ./...
```

For verbose test output:

```bash
go test -v ./...
```

---

## License

This project is open source and available under the [MIT License](LICENSE).
