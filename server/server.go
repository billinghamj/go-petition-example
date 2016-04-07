package server

import (
	"log"
	"net/http"
	"os"
	"time"

	"gopkg.in/tylerb/graceful.v1"

	"github.com/billinghamj/go-petition-example/app"
	"github.com/codegangsta/negroni"
	"github.com/phyber/negroni-gzip/gzip"
)

// Server uses an App, but exposes its functionality via HTTP
type Server struct {
	app     *app.App
	negroni *negroni.Negroni
}

// Create a Server instance with an App instance
func Create(app *app.App) *Server {
	n := negroni.New()
	server := &Server{app, n}

	initServer(server)

	return server
}

// Run the Server and listen on the specified address
func (server *Server) Run(addr string) {
	l := log.New(os.Stdout, "[server] ", 0)
	l.Printf("listening on %s", addr)
	graceful.Run(addr, 10*time.Second, server.negroni)
}

func initServer(server *Server) {
	server.negroni.Use(negroni.NewRecovery())
	server.negroni.Use(negroni.NewLogger())
	server.negroni.Use(gzip.Gzip(gzip.DefaultCompression)) // order sensitive
	// todo: reject non-post

	handler := createMux(server.app)
	server.negroni.UseHandler(handler)
}

func createMux(app *app.App) http.Handler {
	mux := http.NewServeMux()

	mux.Handle("/1/signature_list", createHandler(app, signatureList))
	mux.Handle("/1/signature_create", createHandler(app, signatureCreate))

	return mux
}
