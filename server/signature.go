package server

import (
	"github.com/billinghamj/go-petition-example/log"
	"github.com/billinghamj/go-petition-example/models"
	"github.com/billinghamj/go-petition-example/server/handler"
)

func signatureList(ctx handler.Context, input map[string]interface{}) (interface{}, *log.Error) {
	// todo: validate input

	return ctx.App.ListPetitionSignatures()
}

func signatureCreate(ctx handler.Context, input map[string]interface{}) (interface{}, *log.Error) {
	// todo: validate input

	data := models.Signature{}
	if err := mapInput(input, &data); err != nil {
		return nil, err
	}

	return ctx.App.SignPetition(data)
}
