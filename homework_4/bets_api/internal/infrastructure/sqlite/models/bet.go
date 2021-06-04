package models

type Bet struct {
	Id                   string
	CustomerId           string
	Status               string
	SelectionId          string
	SelectionCoefficient int
	Payment              int
	Payout               int
}
