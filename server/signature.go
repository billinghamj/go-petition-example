package server

import (
	"github.com/billinghamj/go-petition-example/models"
	"github.com/mitchellh/mapstructure"
)

func signatureList(ctx context, input map[string]interface{}) (interface{}, error) {
	// todo: validate input

	return ctx.App.ListPetitionSignatures()
}

func signatureCreate(ctx context, input map[string]interface{}) (interface{}, error) {
	// todo: validate input

	data := models.Signature{}
	mapstructure.Decode(input, &data)

	return ctx.App.SignPetition(data)
}
