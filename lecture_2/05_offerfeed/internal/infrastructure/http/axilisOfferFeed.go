package http

import (
	"code-cadets-2021/lecture_2/05_offerfeed/internal/domain/models"
	"context"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
	"time"
)

const axilisFeedURL = "http://18.193.121.232/axilis-feed"

type AxilisOfferFeed struct {
	httpClient http.Client
	updates    chan models.Odd
}

func NewAxilisOfferFeed(
	httpClient http.Client,
) *AxilisOfferFeed {
	return &AxilisOfferFeed{
		httpClient: httpClient,
		updates:    make(chan models.Odd),
	}
}

func (a *AxilisOfferFeed) Start(ctx context.Context) error {
	// repeatedly:
	// - get odds from HTTP server
	// - write them to updates channel
	// - if context is finished, exit and close updates channel
	// (test your program from cmd/main.go)
	defer close(a.updates)

	for {
		select {
			case <-ctx.Done():
				fmt.Println("finsihed")
				return nil
			case <-time.After(time.Second):
				httpResponse, err := a.httpClient.Get(axilisFeedURL)
				if err != nil {
					return errors.WithMessage(err, "response")
				}

				bodyContent, err := ioutil.ReadAll(httpResponse.Body)
				if err != nil {
					return errors.WithMessage(err, "body")
				}

				var input []axilisOfferOdd
				err = json.Unmarshal(bodyContent, &input)
				if err != nil {
					return errors.WithMessage(err, "unmarshal")
				}
				for _, odd := range input {
					a.updates <- models.Odd{
						Id:          odd.Id,
						Name:        odd.Name,
						Match:       odd.Match,
						Coefficient: odd.Details.Price,
						Timestamp:   time.Time{},
					}
				}
		}
	}
}

func (a *AxilisOfferFeed) GetUpdates() chan models.Odd {
	return a.updates
}

type axilisOfferOdd struct {
	Id      string
	Name    string
	Match   string
	Details axilisOfferOddDetails
}

type axilisOfferOddDetails struct {
	Price float64
}
