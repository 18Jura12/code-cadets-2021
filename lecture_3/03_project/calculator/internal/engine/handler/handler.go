package handler

import (
	"context"
	"github.com/mattn/go-sqlite3"
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
			if err == sqlite3.ErrConstraint {
				log.Println("bet already exists in calc_bets")
				continue
			} else if err != nil {
				log.Println("Failed to insert bet, error: ", err)
				continue
			}
			log.Println("Inserted bet, betId: ", domainBet.Id, ", selectionId: ", domainBet.SelectionId)

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
				log.Println("Failed to fetch bets which should be updated, error: ", err)
				continue
			}
			if !exists {
				log.Println("Event update references 0 bets, selection id: ", eventUpdate.Id)
				continue
			}

			for _, domainBet := range domainBets {
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
					log.Println("Updated bet, betId: ", betCalculated.Id, ", status: ", betCalculated.Status, ", payout: ", betCalculated.Payout)
				case <-ctx.Done():
					return
				}
			}
		}
	}()

	return betsCalculated
}
