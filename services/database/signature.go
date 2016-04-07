package database

import (
	"github.com/billinghamj/go-petition-example/models"
	"labix.org/v2/mgo"
)

var signatureIndexes = []mgo.Index{
	mgo.Index{Key: []string{"email"}, Unique: true},
}

// SignatureFetchAll retrieves all Signature instances from the database
func (db *Database) SignatureFetchAll() ([]models.Signature, error) {
	data := []models.Signature{}
	if err := db.mongoDb.C("signatures").Find(nil).All(&data); err != nil {
		return nil, err
	}
	return data, nil
}

// SignatureCreate inserts a new Signature instance into the database
func (db *Database) SignatureCreate(data *models.Signature) error {
	return db.mongoDb.C("signatures").Insert(data)
}
