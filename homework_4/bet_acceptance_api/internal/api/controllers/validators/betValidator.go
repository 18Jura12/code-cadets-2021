package validators

import (
	"github.com/superbet-group/code-cadets-2021/homework_4/bet_accceptance_api/internal/api/controllers/models"
)

type BetValidator struct {
	selectionCoefficientUpperBound float64
	paymentLowerBound float64
	paymentUpperBound float64
}

func NewBetValidator(selectionCoefficientUpperBound, paymentLowerBound, paymentUpperBound float64) *BetValidator {
	return &BetValidator{
		selectionCoefficientUpperBound: selectionCoefficientUpperBound,
		paymentLowerBound: paymentLowerBound,
		paymentUpperBound: paymentUpperBound,
	}
}

func (b *BetValidator) IsBetValid(betRequest models.BetRequestDto) bool {
	return 	betRequest.SelectionId != "" && betRequest.CustomerId != "" &&
			b.isSelectionCoefficientValid(betRequest.SelectionCoefficient) &&
			b.isPaymentValid(betRequest.Payment)
}

func (b *BetValidator) isSelectionCoefficientValid(selectionCoefficient float64) bool {
	return 	selectionCoefficient != 0 || selectionCoefficient >= 0 ||
			selectionCoefficient <= b.selectionCoefficientUpperBound
}

func (b *BetValidator) isPaymentValid(payment float64) bool {
	return 	payment != 0 || payment >= b.paymentLowerBound ||
			payment <= b.paymentLowerBound
}
