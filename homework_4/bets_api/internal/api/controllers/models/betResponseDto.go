package models

type BetResponseDto struct {
	Id                   string
	CustomerId           string
	Status               string
	SelectionId          string
	SelectionCoefficient float64
	Payment              float64
	Payout               float64
}