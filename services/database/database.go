package database

import (
	"github.com/billinghamj/go-petition-example/log"
	"github.com/billinghamj/go-petition-example/models"
	"labix.org/v2/mgo"
)

// Database represents a MongoDB server connection session
type Database interface {
	SignatureFetchAll() ([]models.Signature, *log.Error)
	SignatureCreate(data *models.Signature) *log.Error
}

type database struct {
	mongoSession *mgo.Session
	mongoDb      *mgo.Database
}

// Create a database connection to the local MongoDB server
func Create(name string) (Database, error) {
	session, err := mgo.Dial("mongodb://localhost")

	if err != nil {
		return nil, err
	}

	db := session.DB(name)
	database := &database{session, db}

	if err := indexesSetup(database); err != nil {
		return nil, err
	}

	return database, nil
}
