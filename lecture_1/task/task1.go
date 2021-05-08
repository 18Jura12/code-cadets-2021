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

	for _, application := range applications {
		if application.Passed {
			var hasSkills = false
			for _, skill := range application.Skills {
				if skill == "Java" || skill == "Go" { hasSkills = true }
			}
			if hasSkills {
				var data string
				data = application.Name + " - "
				for idx, skill := range application.Skills {
					data += skill
					if idx != len(application.Skills) - 1 {
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



