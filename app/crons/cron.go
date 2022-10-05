package crons

import "github.com/nightborn-be/blink/blink-demo/app/usecases"

type CronJob struct{}

func Initialise(usecase usecases.Usecase) CronJob {
	return CronJob{}
}

func (jobs CronJob) Run() error {
	return nil
}
