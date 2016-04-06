package app

import "github.com/billinghamj/go-petition-example/models"

// SignPetition for social justice and safe spaces
func (app *App) SignPetition(signature *models.Signature) error {
	return app.database.SignatureCreate(signature)
}

// ListPetitionSignatures provides data about the signatories
func (app *App) ListPetitionSignatures() ([]models.Signature, error) {
	return app.database.SignatureFetchAll()
}
