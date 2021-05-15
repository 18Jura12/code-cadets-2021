package pokemonencounters

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"

	"github.com/pkg/errors"
	"github.com/sethgrid/pester"
)

type Area struct {
	LocationArea Location `json:"location_area"`
}

type Location struct {
	Name string
}

type Pokemon struct {
	Encounters string `json:"location_area_encounters"`
	Name string `json:"name"`
}

type PokemonEncounters struct {
	Name string
	Places []string
}

const URL = "https://pokeapi.co/api/v2/pokemon/"

func GetAreas() error {

	var input string
	flag.StringVar(&input, "pokemon", "", "name or number of pokemon ( --pokemon _ )")

	flag.Parse()

	if input == "" {
		return errors.New("pokemon name/number not provided")
	}

	url := URL + input

	bodyContentPokemon, err := getRequest(url)
	if err != nil {
		return err
	}

	var pokemon Pokemon
	err = json.Unmarshal(bodyContentPokemon, &pokemon)
	if err != nil {
		return errors.WithMessage(err, "unmarshalling the JSON body content")
	}

	bodyContentPlaces, err := getRequest(pokemon.Encounters)
	if err != nil {
		return err
	}

	var places []Area
	err = json.Unmarshal(bodyContentPlaces, &places)
	if err != nil {
		return errors.WithMessage(err, "unmarshalling the JSON body content")
	}

	var placeList []string
	for _, place := range places {
		placeList = append(placeList, place.LocationArea.Name)
	}

	pokemonEncounters := PokemonEncounters{
		Name: pokemon.Name,
		Places:placeList,
	}

	output, err := json.MarshalIndent(pokemonEncounters, "", "\t")
	if err != nil {
		return err
	}

	fmt.Println(string(output))

	return nil

}

func getRequest(url string) ([]byte, error) {

	httpClient := pester.New()

	httpResponse, err := httpClient.Get(url)
	if err != nil {
		return nil, errors.WithMessage(err, "HTTP get towards API")
	}

	bodyContent, err := ioutil.ReadAll(httpResponse.Body)
	if err != nil {
		return nil, errors.WithMessage(err, "reading body of API response")
	}

	return bodyContent, nil

}
