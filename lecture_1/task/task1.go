package main

import (
	"encoding/json"
	"github.com/pkg/errors"
	"github.com/sethgrid/pester"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type application struct {
	Name string
	Age int
	Passed bool
	Skills []string
}

const URL = "https://run.mocky.io/v3/f7ceece5-47ee-4955-b974-438982267dc8"

func main() {
	httpClient := pester.New()

	httpResponse, err := httpClient.Get(URL)
	if err != nil {
		log.Fatal(
			errors.WithMessage(err, "HTTP get towards API"),
		)
	}

	bodyContent, err := ioutil.ReadAll(httpResponse.Body)
	if err != nil {
		log.Fatal(
			errors.WithMessage(err, "reading body of API response"),
		)
	}

	var applications []application
	err = json.Unmarshal(bodyContent, &applications)
	if err != nil {
		log.Fatal(
			errors.WithMessage(err, "unmarshalling the JSON body content"),
		)
	}

	f, err := os.Create("people.txt")
	if err != nil {
		log.Fatal(
			errors.WithMessage(err, "opening a file"),
		)
	}
	defer f.Close()

	for _, userApplication := range applications {
		if !userApplication.Passed { continue }
		if 	!contains(userApplication.Skills, "Java") &&
			!contains(userApplication.Skills, "Go") { continue }
		var data string
		data = 	userApplication.Name + " - " +
				strings.Join(userApplication.Skills, ", ") + "\n"
		f.WriteString(data)
	}

}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}



