package database

import "labix.org/v2/mgo"

var indexesMap = map[string][]mgo.Index{
	"signatures": signatureIndexes,
}

func indexesSetup(db *Database) error {
	for k, v := range indexesMap {
		if err := indexesIndex(db.mongoDb.C(k), v); err != nil {
			return err
		}
	}

	return nil
}

func indexesIndex(collection *mgo.Collection, indexes []mgo.Index) error {
	for _, v := range indexes {
		if err := collection.EnsureIndex(v); err != nil {
			return err
		}
	}

	return nil
}
