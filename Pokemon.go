package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

// A Response struct to map the Entire Response
type Response struct {
	Name      string `json:"name"`
	PokemonID int    `json:"id"`
}

// A Pokemon Struct to map every pokemon to.
//type Pokemon struct {
//	EntryNo int            `json:"entry_number"`
//	Species PokemonSpecies `json:"pokemon_species"`
//}

// A struct to map our Pokemon's Species which includes it's name
//type PokemonSpecies struct {
//	Name string `json:"name"`
//}

func apiAddress(pokeInput string) string {
	var complete string
	address := "http://pokeapi.co/api/v2/pokemon/"
	complete = address + pokeInput

	return complete
}

func main() {
	//	response, err := http.Get("http://pokeapi.co/api/v2/pokedex/kanto/")

	var userInput string
	fmt.Println("Please name a Pokemon.")
	fmt.Scanln(&userInput)

	response, err := http.Get(apiAddress(userInput))
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

	fmt.Println("Name:" + responseObject.Name)
	fmt.Printf("Pokedex ID: %v\n", responseObject.PokemonID)

}
