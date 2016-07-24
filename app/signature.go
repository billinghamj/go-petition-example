package app

import (
	"github.com/billinghamj/go-petition-example/log"
	"github.com/billinghamj/go-petition-example/models"
)

// SignPetition for social justice and safe spaces
func (app *app) SignPetition(signature models.Signature) (*models.Signature, *log.Error) {
	copy := &signature
	if err := app.database.SignatureCreate(copy); err != nil {
		return nil, err
	}
	return copy, nil
}

// ListPetitionSignatures provides data about the signatories
func (app *app) ListPetitionSignatures() ([]models.Signature, *log.Error) {
	return app.database.SignatureFetchAll()
}
