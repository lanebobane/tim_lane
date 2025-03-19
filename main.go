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
	Height    int  `json:"height"`
	ID        int  `json:"id"`
	IsDefault bool `json:"is_default"`
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
	fmt.Println("Weight: " + strconv.Itoa(pokemonReponse.Height))
	fmt.Println("ID: " + strconv.Itoa(pokemonReponse.ID))
	var isDefaultString string = fmt.Sprintf("%t", pokemonReponse.IsDefault)
	fmt.Println("OG: " + isDefaultString)
}

// https://pkg.go.dev/net/http
// Requirements
// 1. this package can be imported in another Go repo.
// 2. this package cn be used to communicate with any single Slack workspace.
// 3. messages can be sent to any public? channel.
// 3a. source name
// 4. Messages can be <250 words and supports emojis.
// 5. links or directions to related resources?
