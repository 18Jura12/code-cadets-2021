package bootstrap

import (
	stdhttp "net/http"
	"task_1/internal/infrastructure/http"
)

func NewAxilisOfferFeed() *http.AxilisOfferFeed {
	return http.NewAxilisOfferFeed(stdhttp.Client{})
}
