package database

import (
	"github.com/billinghamj/go-petition-example/log"
	"github.com/billinghamj/go-petition-example/models"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

var signatureIndexes = []mgo.Index{
	mgo.Index{Key: []string{"email"}, Unique: true},
}

// SignatureFetchAll retrieves all Signature instances from the database
func (db *Database) SignatureFetchAll() ([]models.Signature, *log.Error) {
	data := []models.Signature{}

	if err := db.mongoDb.C("signatures").Find(nil).All(&data); err != nil {
		return nil, log.CreateError("fetch_failed", map[string]interface{}{"error": err})
	}

	return data, nil
}

// SignatureCreate inserts a new Signature instance into the database
func (db *Database) SignatureCreate(data *models.Signature) *log.Error {
	if data.ID == "" {
		data.ID = bson.NewObjectId()
	}

	if err := db.mongoDb.C("signatures").Insert(data); err != nil {
		return log.CreateError("creation_failed", map[string]interface{}{"error": err})
	}

	return nil
}
