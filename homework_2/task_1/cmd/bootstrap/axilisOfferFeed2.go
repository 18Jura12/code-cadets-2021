package bootstrap

import (
	stdhttp "net/http"
	"task_1/internal/infrastructure/http2"
)

func NewAxilisOfferFeed2() *http2.AxilisOfferFeed {
	return http2.NewAxilisOfferFeed(stdhttp.Client{})
}