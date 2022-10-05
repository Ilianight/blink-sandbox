/* This file is auto-generated, manual edits in this file will be overwritten! */
package controllers

import "github.com/nightborn-be/blink/blink-demo/app/usecases"

type Controller struct {
	PetController PetController
}

func Initialise(usecases usecases.Usecase) Controller {
	return Controller{
		PetController: InitialisePetController(usecases),
	}
}
