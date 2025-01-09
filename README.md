# Pokedex CLI

**Pokedex CLI** is a command-line interface for interacting with the PokeAPI. This project allows you to explore locations, catch Pokémon, view your caught Pokémon, and inspect their details.

---

## Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/iavianm/pokedex.git
   cd pokedex
   ```
2. Install dependencies:
    ```bash
      go mod tidy
    ```
3. Build the application:
    ```bash
     go build -o pokedex
    ```
4. Run the application:
    ```bash
     ./pokedex
    ```
## Features

** Available Commands **

| Command           | Description                                       |
|-------------------|---------------------------------------------------|
| `help`            | Display a help message with all available commands. |
| `catch <name>`    | Attempt to catch a Pokémon by name.               |
| `explore <location>` | Explore a specific location.                   |
| `inspect <name>`  | Inspect a caught Pokémon’s details.               |
| `pokedex`         | Show a list of all caught Pokémon.                |
| `map`             | Go to the next page of locations.                 |
| `mapb`            | Go to the previous page of locations.             |
| `exit`            | Exit the application.                             |