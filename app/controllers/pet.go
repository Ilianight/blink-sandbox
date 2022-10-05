/* This file is auto-generated, manual edits in this file will be overwritten! */
package controllers

import (
	"net/http"

	"github.com/getsentry/sentry-go"
	"github.com/gin-gonic/gin"
	"github.com/nightborn-be/blink/blink-demo/app/contexts"
	"github.com/nightborn-be/blink/blink-demo/app/usecases"
)

type PetController struct {
	usecases usecases.Usecase
}

func InitialisePetController(usecases usecases.Usecase) PetController {
	return PetController{usecases: usecases}
}

func (controller PetController) AddPet(c *gin.Context) {
	context, err := contexts.GetContext(c)
	if err != nil {
		sentry.CaptureException(err)
		c.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}

	// Call the usecase
	err = controller.usecases.PetUsecase.AddPet(context)
	if err != nil {
		sentry.CaptureException(err)
		c.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}

	// The response
	c.IndentedJSON(http.StatusNoContent, nil)
}
