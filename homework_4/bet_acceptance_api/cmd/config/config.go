package config

import "github.com/kelseyhightower/envconfig"

var Cfg Config

type Config struct {
	Rabbit rabbitConfig `split_words:"true"`
	Api    apiConfig    `split_words:"true"`
	ConstVariables constConfig `split_words:"true"`
}

type constConfig struct {
	SelectionCoefficientUpperBound float64 `split_words:"true" default:"10"`
	PaymentLowerBound float64 `split_words:"true" default:"2"`
	PaymentUpperBound float64 `split_words:"true" default:"100"`
}

type apiConfig struct {
	ReadWriteTimeoutMs int `split_words:"true" default:"10000"`
	Port               int `split_words:"true" default:"8082"`
}

type rabbitConfig struct {
	PublisherBetQueueQueue string `split_words:"true"  default:"bets"`
	PublisherExchange      string `split_words:"true" default:""`
	PublisherMandatory     bool   `split_words:"true" default:"false"`
	PublisherImmediate     bool   `split_words:"true" default:"false"`
}

func Load() {
	err := envconfig.Process("", &Cfg)
	if err != nil {
		panic(err)
	}
}

