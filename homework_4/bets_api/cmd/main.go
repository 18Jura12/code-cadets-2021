package main

import (
	"github.com/superbet-group/code-cadets-2021/homework_4/bets_api/cmd/bootstrap"
	"github.com/superbet-group/code-cadets-2021/homework_4/bets_api/cmd/config"
	"github.com/superbet-group/code-cadets-2021/homework_4/bets_api/internal/tasks"
	"log"
)

func main() {
	log.Println("Bootstrap initiated")

	config.Load()

	db := bootstrap.Sqlite()
	signalHandler := bootstrap.SignalHandler()
	api := bootstrap.Api(db)

	log.Println("Bootstrap finished. Bets API is starting")

	tasks.RunTasks(signalHandler, api)

	log.Println("Bets API finished gracefully")
}
