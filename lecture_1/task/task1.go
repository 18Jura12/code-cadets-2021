package main

import (
	"encoding/json"
	"github.com/pkg/errors"
	"github.com/sethgrid/pester"
	"io/ioutil"
	"log"
	"os"
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
	//httpClient.Backoff = linearBackoff

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

	var decodedContent []application
	err = json.Unmarshal(bodyContent, &decodedContent)
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

	for _, val := range decodedContent {
		if val.Passed {
			var HasSkills = false
			for _, skill := range val.Skills {
				if skill == "Java" || skill == "Go" { HasSkills = true }
			}
			if HasSkills {
				var data string
				data = val.Name + " - "
				for idx, skill := range val.Skills {
					data += skill
					if idx != len(val.Skills) - 1 {
						data += ", "
					} else {
						data += "\n"
					}
				}
				f.WriteString(data)
			}
		}
	}

}



