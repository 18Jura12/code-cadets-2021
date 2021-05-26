package sqlite

import (
	"context"
	"database/sql"
	"github.com/pkg/errors"
	domainmodels "github.com/superbet-group/code-cadets-2021/lecture_3/03_project/calculator/internal/domain/models"
	storagemodels "github.com/superbet-group/code-cadets-2021/lecture_3/03_project/calculator/internal/infrastructure/sqlite/models"
)

type CalcBetRepository struct {
	dbExecutor    DatabaseExecutor
	calcbetMapper CalcBetMapper
}

func NewCalcBetRepository(dbExecutor DatabaseExecutor, calcbetMapper CalcBetMapper) *CalcBetRepository {
	return &CalcBetRepository{
		dbExecutor:    dbExecutor,
		calcbetMapper: calcbetMapper,
	}
}

func (r *CalcBetRepository) InsertCalcBet(ctx context.Context, bet domainmodels.Bet) error {
	storageCalcBet := r.calcbetMapper.MapDomainBetToStorageBet(bet)
	err := r.queryInsertCalcBet(ctx, storageCalcBet)
	if err != nil {
		return errors.Wrap(err, "calculated bet repository failed to insert a bet with id "+bet.Id)
	}
	return nil
}

func (r *CalcBetRepository) queryInsertCalcBet(ctx context.Context, calcbet storagemodels.CalcBet) error {

	insertBetSQL := "INSERT INTO bets(id, selection_id, selection_coefficient, payment) VALUES (?, ?, ?, ?)"
	statement, err := r.dbExecutor.PrepareContext(ctx, insertBetSQL)
	if err != nil {
		return err
	}

	_, err = statement.ExecContext(ctx, calcbet.Id, calcbet.SelectionId, calcbet.SelectionCoefficient, calcbet.Payment)
	return err
}

func (r *CalcBetRepository) UpdateCalcBet(ctx context.Context, bet domainmodels.Bet) error {
	storageBet := r.calcbetMapper.MapDomainBetToStorageBet(bet)
	err := r.queryUpdateCalcBet(ctx, storageBet)
	if err != nil {
		return errors.Wrap(err, "calculated bet repository failed to update a bet with id "+bet.Id)
	}
	return nil
}

func (r *CalcBetRepository) queryUpdateCalcBet(ctx context.Context, bet storagemodels.CalcBet) error {
	updateBetSQL := "UPDATE bets SET selection_id=?, selection_coefficient=?, payment=? WHERE id=?"

	statement, err := r.dbExecutor.PrepareContext(ctx, updateBetSQL)
	if err != nil {
		return err
	}

	_, err = statement.ExecContext(ctx, bet.SelectionId, bet.SelectionCoefficient, bet.Payment, bet.Id)
	return err
}

func (r *CalcBetRepository) GetCalcBetByID(ctx context.Context, id string) (domainmodels.Bet, bool, error) {
	storageBet, err := r.queryGetCalcBetByID(ctx, id)
	if err == sql.ErrNoRows {
		return domainmodels.Bet{}, false, nil
	}
	if err != nil {
		return domainmodels.Bet{}, false, errors.Wrap(err, "calculated bet repository failed to get a bet with id "+id)
	}

	domainBet := r.calcbetMapper.MapStorageBetToDomainBet(storageBet)
	return domainBet, true, nil
}

func (r *CalcBetRepository) queryGetCalcBetByID(ctx context.Context, id string) (storagemodels.CalcBet, error) {
	row, err := r.dbExecutor.QueryContext(ctx, "SELECT * FROM bets WHERE id='"+id+"';")
	if err != nil {
		return storagemodels.CalcBet{}, err
	}
	defer row.Close()

	// This will move to the "next" result (which is the only result, because a single bet is fetched).
	row.Next()

	var selectionId string
	var selectionCoefficient int
	var payment int

	err = row.Scan(&id, &selectionId, &selectionCoefficient, &payment)
	if err != nil {
		return storagemodels.CalcBet{}, err
	}

	return storagemodels.CalcBet{
		Id:                   id,
		SelectionId:          selectionId,
		SelectionCoefficient: selectionCoefficient,
		Payment:              payment,
	}, nil
}

func (r *CalcBetRepository) GetCalcBetsBySelectionID(ctx context.Context, selectionId string) ([]domainmodels.Bet, bool, error) {
	storageBets, err := r.queryGetCalcBetsBySelectionID(ctx, selectionId)
	if err == sql.ErrNoRows {
		return nil, false, nil
	}
	if err != nil {
		return nil, false, errors.Wrap(err, "calculated bet repository failed to get a bet with selectionId "+selectionId)
	}

	var domainBets []domainmodels.Bet
	for _, storageBet := range storageBets {
		domainBets = append(domainBets,  r.calcbetMapper.MapStorageBetToDomainBet(storageBet))
	}
	return domainBets, true, nil
}

func (r *CalcBetRepository) queryGetCalcBetsBySelectionID(ctx context.Context, selectionId string) ([]storagemodels.CalcBet, error) {
	row, err := r.dbExecutor.QueryContext(ctx, "SELECT * FROM bets WHERE selectionId='"+selectionId+"';")
	if err != nil {
		return nil, err
	}
	defer row.Close()

	var rows []storagemodels.CalcBet
	for row.Next() {
		var id string
		var selectionCoefficient int
		var payment int

		err = row.Scan(&id, &selectionId, &selectionCoefficient, &payment)
		if err != nil {
			return nil, err
		}

		rows = append(
			rows,
			storagemodels.CalcBet{
				Id:                   id,
				SelectionId:          selectionId,
				SelectionCoefficient: selectionCoefficient,
				Payment:              payment,
			},
		)
	}

	return rows, nil
}