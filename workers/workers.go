package main

import (
	"log"

	"asynq-quickstart/task"

	"github.com/hibiken/asynq"
)

// workers.go
func main() {
	srv := asynq.NewServer(
		asynq.RedisClusterClientOpt{Addrs: []string{"localhost:8000"}},
		asynq.Config{Concurrency: 10},
	)

	mux := asynq.NewServeMux()
	mux.HandleFunc(task.TypeWelcomeEmail, task.HandleWelcomeEmailTask)
	mux.HandleFunc(task.TypeReminderEmail, task.HandleReminderEmailTask)

	if err := srv.Run(mux); err != nil {
		log.Fatal(err)
	}
}
