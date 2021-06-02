package main

import (
	"bytes"
	"encoding/json"
	"github.com/pkg/errors"
	"github.com/sethgrid/pester"
	"io/ioutil"
	"log"
	"math/rand"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

const getSelectionIds = "http://127.0.0.1:8081/bets?status=active"
const postEventUpdate = "http://127.0.0.1:8080/event/update"

type eventUpdateDto struct {
	Id      string `json:"id"`
	Outcome string `json:"outcome"`
}

type betDto struct {
	SelectionId          string `json:"SelectionId"`
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
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

func postRequest(url string, data []byte) error {

	httpClient := pester.New()

	_, err := httpClient.Post(url, "application/json", bytes.NewBuffer(data))
	if err != nil {
		return errors.WithMessage(err, "HTTP get towards API")
	}

	return nil

}

func main() {

	rand.Seed(time.Now().UnixNano())

	bodyContentSelectionIds, err := getRequest(getSelectionIds)
	failOnError(err, "failed to fetch bets")

	var selectionIds []betDto
	err = json.Unmarshal(bodyContentSelectionIds, &selectionIds)
	failOnError(err, "unmarshalling the JSON body content: " + string(bodyContentSelectionIds))

	unique := make(map[string]bool)
	for _, bet := range selectionIds {
		unique[bet.SelectionId] = true
	}

	for selectionId := range unique {
		var outcome string
		if rand.Float64() > 0.5 {
			outcome = "lost"
		} else {
			outcome = "won"
		}

		eventUpdate := &eventUpdateDto{
			Id: selectionId,
			Outcome: outcome,
		}

		eventUpdateJson, err := json.Marshal(eventUpdate)
		failOnError(err, "failed to marshal an event update")

		err = postRequest(postEventUpdate, eventUpdateJson)
		failOnError(err, "failed to post event update request")

		log.Printf("Sent %s", eventUpdateJson)
	}
}
