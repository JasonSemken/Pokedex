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

var responseObject Response

// A Response struct to map the Entire Response
type Response struct {
	Name        string      `json:"name"`
	PokemonID   int         `json:"id"`
	PokemonStat []PokeStats `json:"stats"`
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

// Appends the users input onto the api string
func apiAddress() string {
	var cA string
	iA := "http://pokeapi.co/api/v2/pokemon/"
	cA = iA + userInput()

	return cA
}

// Requests users input and converts to lowercase to ensure compatibility with api
func userInput() string {
	var uI string

	fmt.Println("Please name a Pokemon.")
	fmt.Scanln(&uI)
	fI := strings.ToLower(uI)

	return fI
}

// Compares the pokemon ID to 0, 0 means the user input does not make a working API string. Returns requested Pokemon if not 0
func requestReturn() {

	if responseObject.PokemonID == 0 {
		fmt.Println("\nThat's Not a Pokemon.")
	} else {
		fmt.Printf("\nName: %v\n", cases.Title(language.Und, cases.NoLower).String(responseObject.Name))
		fmt.Printf("Pokedex ID: %v\n", responseObject.PokemonID)
		fmt.Printf("HP: %v\n", responseObject.PokemonStat[0].Value)
		fmt.Printf("Attack: %v\n", responseObject.PokemonStat[1].Value)
		fmt.Printf("Defense: %v\n", responseObject.PokemonStat[2].Value)
		fmt.Printf("Speed: %v\n", responseObject.PokemonStat[5].Value)
		fmt.Printf("Special Attack: %v\n", responseObject.PokemonStat[3].Value)
		fmt.Printf("Special Defense: %v\n", responseObject.PokemonStat[4].Value)
	}
}

func main() {

	response, err := http.Get(apiAddress())
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	json.Unmarshal(responseData, &responseObject)

	requestReturn()

}
