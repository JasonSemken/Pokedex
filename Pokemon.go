package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// Global variable to take user input and spread to other functions
var FormattedInput string

// A Response struct to map the Entire Response
type Response struct {
	Name               string               `json:"name"`
	PokemonID          int                  `json:"id"`
	PokemonStat        []PokeStats          `json:"stats"`
	PokemonDescription []DescriptionVersion `json:"flavor_text_entries"`
}

// Struct to hold pokemon stat and value
type PokeStats struct {
	StatName Stat `json:"stat"`
	Value    int  `json:"base_stat"`
}

// Collects the name of the Pokemon stat
type Stat struct {
	Name string `json:"name"`
}

type DescriptionVersion struct {
	Descripton  string `json:"flavor_text"`
	GameVersion string `json:"version"`
}

type GameDescription struct {
	Game string `json:"game"`
}

// Requests users input and converts to lowercase to ensure compatibility with api
func userInput() string {
	var uI string

	fmt.Println("Please name a Pokemon.")
	fmt.Scanln(&uI)
	FormattedInput := strings.ToLower(uI)

	return FormattedInput

}

// Appends the users input onto the api pokemon string
func pokemonAddress() string {
	var cA string

	iA := "http://pokeapi.co/api/v2/pokemon/"
	cA = iA + FormattedInput

	return cA
}

// Appends the users input onto the api pokemon string
func pokemonSpeciesAddress() string {
	var cA string

	iA := "http://pokeapi.co/api/v2/pokemon-species/"
	cA = iA + FormattedInput

	return cA
}

// Compares the pokemon ID to 0, 0 means the user input does not make a working API string. Returns requested Pokemon if not 0
func pokemonRequestReturn(rD Response, sD Response) {

	if rD.PokemonID == 0 {
		fmt.Printf("%v", FormattedInput)
		fmt.Println("\nThat's Not a Pokemon.")
	} else {
		fmt.Printf("\nName: %v\n", cases.Title(language.Und, cases.NoLower).String(rD.Name))
		fmt.Printf("Pokedex ID: %v\n", rD.PokemonID)
		fmt.Printf("Description: \n%v\n\n", sD.PokemonDescription[13].Descripton)
		fmt.Printf("HP: %v\n", rD.PokemonStat[0].Value)
		fmt.Printf("Attack: %v\n", rD.PokemonStat[1].Value)
		fmt.Printf("Defense: %v\n", rD.PokemonStat[2].Value)
		fmt.Printf("Speed: %v\n", rD.PokemonStat[5].Value)
		fmt.Printf("Special Attack: %v\n", rD.PokemonStat[3].Value)
		fmt.Printf("Special Defense: %v\n", rD.PokemonStat[4].Value)
	}
}

// Reusable function to call api based on string input
func callAPI(api string) Response {

	response, err := http.Get(api)
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var responseObject Response
	json.Unmarshal(responseData, &responseObject)

	return responseObject
}

func main() {

	var selection string

	fmt.Println("What would you like to search?")
	fmt.Println("Pokemon nothing, or reset")
	fmt.Scan(&selection)

	switch selection {
	case "pokemon", "Pokemon":
		FormattedInput = userInput()
		pokemonRequestReturn(callAPI(pokemonAddress()), callAPI(pokemonSpeciesAddress()))
	case "nothing", "Nothing":
	default:

	}

}
