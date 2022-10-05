package usecases

import (
	"errors"

	"github.com/nightborn-be/blink/blink-demo/app/contexts"
	"github.com/nightborn-be/blink/blink-demo/app/gateways"
)

type PetUsecase struct {
	gateway *gateways.Gateway
}

func InitialisePetUsecase(gateway *gateways.Gateway) IPetUsecase {
	return PetUsecase{

		gateway: gateway,
	}
}

func (usecase PetUsecase) AddPet(context *contexts.Context) error {
	return errors.New(METHOD_NOT_IMPLEMENTED + "AddPet")
}
