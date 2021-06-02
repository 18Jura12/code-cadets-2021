package bootstrap

import "github.com/superbet-group/code-cadets-2021/homework_4/bet_accceptance_api/internal/tasks"

func SignalHandler() *tasks.SignalHandler {
	return tasks.NewSignalHandler()
}
