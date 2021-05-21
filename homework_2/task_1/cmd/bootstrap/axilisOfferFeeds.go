package bootstrap

import (
	stdhttp "net/http"
	"task_1/internal/infrastructure/http"
)

func NewAxilisOfferFeed(client stdhttp.Client) *http.AxilisOfferFeed {
	return http.NewAxilisOfferFeed(client)
}

func NewAxilisOfferFeed2(client stdhttp.Client) *http.AxilisOfferFeed2 {
	return http.NewAxilisOfferFeed2(client)
}
