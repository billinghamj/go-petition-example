package handler

import (
	"net/http"

	"github.com/billinghamj/go-petition-example/app"
)

// Context contains the app instance and HTTP request
type Context struct {
	App     app.App
	Request *http.Request
}
