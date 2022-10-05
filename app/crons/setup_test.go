package crons

import (
	"github.com/joho/godotenv"
	"github.com/nightborn-be/blink/blink-demo/app/gateways"
	"github.com/nightborn-be/blink/blink-demo/app/usecases"
	"github.com/samber/lo"
)

// Setup test cron
func SetupTestCron() (*CronJob, error) {
	// Load env
	if err := godotenv.Load(); err != nil {
		return nil, err
	}

	// Setup in-memory gateways
	testGateway, _ := gateways.InitialiseTest()

	// Setup usecases with gateways
	usecase := usecases.Initialise(&testGateway)

	// Setup crons
	return lo.ToPtr(Initialise(usecase)), nil
}
