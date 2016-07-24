package app

import (
	"github.com/billinghamj/go-petition-example/log"
	"github.com/billinghamj/go-petition-example/models"
	"github.com/billinghamj/go-petition-example/services/database"
)

// App handles business logic, does not involve HTTP
type App interface {
	SignPetition(signature models.Signature) (*models.Signature, *log.Error)
	ListPetitionSignatures() ([]models.Signature, *log.Error)
}

type app struct {
	database database.Database
}

// Create an instance of App with a database instance
func Create(db database.Database) App {
	return &app{db}
}
