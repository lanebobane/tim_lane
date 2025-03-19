package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
)

type PokeAPIResponse struct {
	Height         int              `json:"height"`
	Weight         int              `json:"weight"`
	ID             int              `json:"id"`
	IsDefault      bool             `json:"is_default"`
	BaseExperience int              `json:"base_experience"`
	Name           string           `json:"name"`
	Order          int              `json:"order"`
	Abilities      []PokeAPIAbility `json:"abilities"`
}

type PokeAPIAbility struct {
	Ability  PokemonAbility `json:"ability"`
	IsHideen bool           `json:"is_hidden"`
	Slot     int            `json:"int"`
}

type PokemonAbility struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

func main() {
	pokemon := "pikachu" // why don't I have to declare a data type here?
	if len(os.Args) > 1 {
		pokemon = os.Args[1]
	}

	fmt.Println("Finding the pokemon: " + pokemon)
	get_pokemon(pokemon)
}

func get_pokemon(pokemon string) {

	endpoint := "https://pokeapi.co/api/v2/pokemon/"
	endpoint += pokemon

	resp, err := http.Get(endpoint)
	if err != nil {
		// handle error
		panic(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		// handle error
		panic(err)
	}
	string_body := string(body)
	var pokemonReponse PokeAPIResponse
	err = json.Unmarshal([]byte(string_body), &pokemonReponse)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Information on " + pokemon + " retrieved successfully from the PokeAPI.")
}
func printPokemon(pokeAPIReponse PokeAPIResponse) {
	fmt.Println("Weight: " + strconv.Itoa(pokeAPIReponse.Height))
	fmt.Println("ID: " + strconv.Itoa(pokeAPIReponse.ID))
	var isDefaultString string = fmt.Sprintf("%t", pokeAPIReponse.IsDefault)
	fmt.Println("OG: " + isDefaultString)
	fmt.Println(pokeAPIReponse.Abilities[0].Ability.Name)
}

// https://pkg.go.dev/net/http
// Requirements
// 1. this package can be imported in another Go repo.
// 2. this package cn be used to communicate with any single Slack workspace.
// 3. messages can be sent to any public? channel.
// 3a. source name
// 4. Messages can be <250 words and supports emojis.
// 5. links or directions to related resources?
