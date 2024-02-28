package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

// A Response struct to map the Entire Response
type Response struct {
	Name      string `json:"name"`
	PokemonID int    `json:"id"`
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

	var responseObject Response
	json.Unmarshal(responseData, &responseObject)

	// Compares the pokemon ID to 0, 0 means the user input does not make a working API string
	if responseObject.PokemonID == 0 {
		fmt.Println("\nThat's Not a Pokemon.")
	} else {
		fmt.Println("\nName:" + responseObject.Name)
		fmt.Printf("Pokedex ID: %v\n", responseObject.PokemonID)
	}

}
