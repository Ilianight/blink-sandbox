package usecases

import (
	"github.com/nightborn-be/blink/blink-demo/app/gateways"
	"github.com/nightborn-be/blink/blink-demo/app/utils"
	"github.com/samber/lo"
)

// Setup test
func setupTestUsecase() (*Usecase, *gateways.Mocks, error) {
	// Initialise faker generators
	utils.InitialiseFakerGenerators()

	// Setup in-memory gateways
	testGateway, mocks := gateways.InitialiseTest()

	// Setup usecases with gateways
	return lo.ToPtr(Initialise(&testGateway)), mocks, nil
}

// func createContext(sub string) contexts.Context {
// 	return contexts.Context{
// 		ContextBase: contexts.ContextBase{
//			Sub:    &sub,
//			Roles:  nil,
//			TestId: nil,
//		},
//	}
// }

// func createAdminContext(sub string) contexts.Context {
//	return contexts.Context{
//		ContextBase: contexts.ContextBase{
//			Sub:    &sub,
//			Roles:  []string{"admin"},
//			TestId: nil,
//		},
//	}
// }
