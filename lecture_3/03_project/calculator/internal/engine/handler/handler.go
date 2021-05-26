package handler

import (
	"context"
	domainmodels "github.com/superbet-group/code-cadets-2021/lecture_3/03_project/calculator/internal/domain/models"
	rabbitmqmodels "github.com/superbet-group/code-cadets-2021/lecture_3/03_project/calculator/internal/infrastructure/rabbitmq/models"
	"log"
)

type Handler struct {
	betRepository CalcBetRepository
}

func New(betRepository CalcBetRepository) *Handler {
	return &Handler{
		betRepository: betRepository,
	}
}

func (h *Handler) HandleBets(ctx context.Context, bets <-chan rabbitmqmodels.Bet ) {
	go func() {
		for bet := range bets {
			log.Println("Processing bet, betId:", bet.Id, " with selectionId:", bet.SelectionId)

			domainBet := domainmodels.Bet{
				Id:                   bet.Id,
				SelectionId:          bet.SelectionId,
				SelectionCoefficient: bet.SelectionCoefficient,
				Payment:              bet.Payment,
			}

			err := h.betRepository.InsertCalcBet(ctx, domainBet)
			if err != nil {
				log.Println("Failed to insert bet, error: ", err)
				continue
			}

			select {
			case <-ctx.Done():
				return
			default:
				continue
			}
		}
	}()
}

func (h *Handler) HandleEventUpdates(
	ctx context.Context,
	eventUpdates <-chan rabbitmqmodels.EventUpdate,
	) <-chan rabbitmqmodels.BetCalculated {
	betsCalculated := make(chan rabbitmqmodels.BetCalculated)

	go func() {
		defer close(betsCalculated)

		for eventUpdate := range eventUpdates {
			log.Println("Processing event update, Id:", eventUpdate.Id)

			domainBets, exists, err := h.betRepository.GetCalcBetsBySelectionID(ctx, eventUpdate.Id)
			if err != nil {
				log.Println("Failed to fetch a bet which should be updated, error: ", err)
				continue
			}
			if !exists {
				log.Println("A bet which should be updated does not exist, selection id: ", eventUpdate.Id)
				continue
			}

			for _, domainBet := range domainBets {
				err = h.betRepository.DeleteCalcBet(ctx, domainBet.Id)
				if err != nil {
					log.Println("Failed to delete a bet which was updated, error: ", err)
				} else {
					log.Println("Successfully deleted updated bet with id: ", domainBet.Id)
				}

				var payout = 0.0
				if eventUpdate.Outcome == "won" {
					payout = domainBet.SelectionCoefficient * domainBet.Payment
				}

				betCalculated := rabbitmqmodels.BetCalculated{
					Id:     domainBet.Id,
					Status: eventUpdate.Outcome,
					Payout: payout,
				}

				select {
				case betsCalculated <- betCalculated:
				case <-ctx.Done():
					return
				}
			}
		}
	}()

	return betsCalculated
}
