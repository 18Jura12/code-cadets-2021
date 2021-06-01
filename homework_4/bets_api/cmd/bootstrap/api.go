package bootstrap

import (
	"github.com/superbet-group/code-cadets-2021/homework_4/bets_api/cmd/config"
	"github.com/superbet-group/code-cadets-2021/homework_4/bets_api/internal/api"
	"github.com/superbet-group/code-cadets-2021/homework_4/bets_api/internal/api/controllers"
	"github.com/superbet-group/code-cadets-2021/homework_4/bets_api/internal/api/controllers/validators"
	"github.com/superbet-group/code-cadets-2021/homework_4/bets_api/internal/domain/mappers"
	"github.com/superbet-group/code-cadets-2021/homework_4/bets_api/internal/domain/services"
	"github.com/superbet-group/code-cadets-2021/homework_4/bets_api/internal/infrastructure/sqlite"
)

func newIdValidator() *validators.IdValidator {
	return validators.NewIdValidator()
}

func newStatusValidator() *validators.StatusValidator {
	return validators.NewStatusValidator()
}

func newBetMapper() *mappers.BetMapper {
	return mappers.NewBetMapper()
}

func newBetRepository(dbExecutor sqlite.DatabaseExecutor, betMapper sqlite.BetMapper) *sqlite.BetRepository {
	return sqlite.NewBetRepository(dbExecutor, betMapper)
}

func newBetService(betRepository services.BetRepository) *services.BetService {
	return services.NewBetService(betRepository)
}

func newController(
	idValidator validators.IdValidator,
	statusValidator validators.StatusValidator,
	betService services.BetService,
) *controllers.Controller {
	return controllers.NewController(&idValidator, &statusValidator, &betService)
}

func Api(dbExecutor sqlite.DatabaseExecutor) *api.WebServer {
	idValidator := newIdValidator()
	statusValidator := newStatusValidator()

	betMapper := newBetMapper()
	betRepository := newBetRepository(dbExecutor, betMapper)
	betService := newBetService(betRepository)

	controller := newController(*idValidator, *statusValidator, *betService)

	return api.NewServer(config.Cfg.Api.Port, config.Cfg.Api.ReadWriteTimeoutMs, controller)
}
