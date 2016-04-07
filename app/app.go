package app

import "github.com/billinghamj/go-petition-example/services/database"

// App handles business logic, does not involve HTTP
type App struct {
	database *database.Database
}

// Create an instance of App with a database instance
func Create(db *database.Database) *App {
	return &App{db}
}
