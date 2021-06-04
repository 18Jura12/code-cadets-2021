package services

type BetPublisher interface {
	Publish(customerId, selectionId string, selectionCoefficient, payment float64) error
}

