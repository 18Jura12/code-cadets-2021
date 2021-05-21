package http

import (
	"context"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/pkg/errors"

	"task_1/internal/domain/models"
)

const secondFeedURL = "http://18.193.121.232/axilis-feed-2"

type AxilisOfferFeed2 struct {
	httpClient http.Client
	updates    chan models.Odd
}

func NewAxilisOfferFeed2(
	httpClient http.Client,
) *AxilisOfferFeed2 {
	return &AxilisOfferFeed2{
		httpClient: httpClient,
		updates:    make(chan models.Odd),
	}
}

func (a *AxilisOfferFeed2) Start(ctx context.Context) error {
	defer close(a.updates)

	for {
		timeout := time.After(time.Second)
		select {
			case <-ctx.Done():
				return nil
			case <-timeout:
				httpResponse, err := a.httpClient.Get(secondFeedURL)
				if err != nil {
					return errors.WithMessage(err, "response 2")
				}

				bodyContent, err := ioutil.ReadAll(httpResponse.Body)
				if err != nil {
					return errors.WithMessage(err, "body 2")
				}

				input := string(bodyContent)
				inputs := strings.Split(input, "\n")
				for _, item := range inputs {
					content := strings.Split(item, ",")
					coef, err := strconv.ParseFloat(content[3], 64)
					if err != nil {
						return errors.WithMessage(err,"parse float")
					}

					a.updates <- models.Odd{
						Id:          content[0],
						Name:        content[1],
						Match:       content[2],
						Coefficient: coef,
						Timestamp:   time.Now(),
					}
				}
		}
	}
}

func (a *AxilisOfferFeed2) GetUpdates() chan models.Odd {
	return a.updates
}

func (a *AxilisOfferFeed2) String() string {
	return "axilis offer feed 2"
}
