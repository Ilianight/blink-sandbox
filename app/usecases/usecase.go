package usecases

import (
	"github.com/nightborn-be/blink/blink-demo/app/gateways"
)

type Usecase struct {
	UsecaseBase
}

func Initialise(gateway *gateways.Gateway) Usecase {
	return Usecase{
		UsecaseBase: UsecaseBase{
			PetUsecase: InitialisePetUsecase(gateway),
		},
	}
}
