package database

import "labix.org/v2/mgo"

// Database represents a MongoDB server connection session
type Database struct {
	mongoSession *mgo.Session
	mongoDb      *mgo.Database
}

// Create a database connection to the local MongoDB server
func Create(name string) (*Database, error) {
	session, err := mgo.Dial("mongodb://localhost")

	if err != nil {
		return nil, err
	}

	db := session.DB(name)
	database := &Database{session, db}

	if err := indexesSetup(database); err != nil {
		return nil, err
	}

	return database, nil
}
