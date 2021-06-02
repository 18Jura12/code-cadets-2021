package main

import (
	"github.com/superbet-group/code-cadets-2021/homework_4/bet_accceptance_api/cmd/bootstrap"
	"github.com/superbet-group/code-cadets-2021/homework_4/bet_accceptance_api/cmd/config"
	"github.com/superbet-group/code-cadets-2021/homework_4/bet_accceptance_api/internal/tasks"
	"log"
)

func main() {
	log.Println("Bootstrap initiated")

	config.Load()

	rabbitMqChannel := bootstrap.RabbitMq()
	signalHandler := bootstrap.SignalHandler()
	api := bootstrap.Api(rabbitMqChannel)

	log.Println("Bootstrap finished. Bet Acceptance API is starting")

	tasks.RunTasks(signalHandler, api)

	log.Println("Bet Acceptance API finished gracefully")
}
