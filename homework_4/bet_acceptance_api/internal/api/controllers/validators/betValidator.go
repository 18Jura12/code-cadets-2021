package validators

import (
	"github.com/superbet-group/code-cadets-2021/homework_4/bet_accceptance_api/internal/api/controllers/models"
)

type BetValidator struct {}

func NewBetValidator() *BetValidator {
	return &BetValidator{}
}

func (b *BetValidator) IsBetValid(betRequest models.BetRequestDto) bool {
	//if 	betRequest.SelectionId == "" || betRequest.CustomerId == "" ||
	//	betRequest.SelectionCoefficient == 0 || betRequest.Payment == 0 {
	//
	//}
	return false
}
