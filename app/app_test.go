package app

import (
	"testing"

	"github.com/billinghamj/go-petition-example/services/database"
)

func TestCreate(t *testing.T) {
	db := &database.Database{}
	app := Create(db)

	if app.database != db {
		t.Error("database doesn't match input")
	}
}
