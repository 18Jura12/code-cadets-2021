package pokemonencounters

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/pkg/errors"
	"github.com/sethgrid/pester"
	"io/ioutil"
)

type Area struct {
	Location struct {
		Name string
	} `json:"location_area"`
}

type Pokemon struct {
	Encounters string `json:"location_area_encounters"`
	Name string `json:"name"`
}

type PokemonEncounters struct {
	Name string
	Places []string
}

func GetAreas() error {

	url := "https://pokeapi.co/api/v2/pokemon/"

	var input string
	flag.StringVar(&input, "pokemon", "1", "name or number of pokemon ( --pokemon _ )")
	flag.Parse()

	url += input

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
		placeList = append(placeList, place.Location.Name)
	}

	pokemonEncounters := PokemonEncounters{
		Name: pokemon.Name,
		Places:placeList,
	}

	output, _ := json.MarshalIndent(pokemonEncounters, "", "\t")
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
