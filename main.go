package main

import (
	"github.com/billinghamj/go-petition-example/app"
	"github.com/billinghamj/go-petition-example/database"
	"github.com/billinghamj/go-petition-example/server"
)

func main() {
	if db, err := database.Create("signatures"); err != nil {
		panic(err)
	} else {
		app := app.Create(db)
		server := server.Create(app)

		server.Run(":3000")
	}
}
