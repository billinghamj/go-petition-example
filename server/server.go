package server

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	"gopkg.in/tylerb/graceful.v1"

	"github.com/billinghamj/go-petition-example/app"
	"github.com/billinghamj/go-petition-example/models"
	"github.com/codegangsta/negroni"
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

	handler := createHandler(server.app)
	server.negroni.UseHandler(handler)
}

func createHandler(app *app.App) http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/signatures", func(w http.ResponseWriter, req *http.Request) {
		switch req.Method {
		default:
			w.WriteHeader(406)

		case "GET":
			if signatures, err := app.ListPetitionSignatures(); err != nil {
				panic(err)
			} else {
				if data, err := json.Marshal(signatures); err != nil {
					panic(err)
				} else {
					w.WriteHeader(200)
					w.Write(data)
				}
			}

		case "POST":
			signature := models.Signature{}

			if err := json.NewDecoder(req.Body).Decode(&signature); err != nil {
				w.WriteHeader(400)
				return
			}

			if err := app.SignPetition(&signature); err != nil {
				w.WriteHeader(409)
				return
			}

			if data, err := json.Marshal(signature); err != nil {
				panic(err)
			} else {
				w.WriteHeader(201)
				w.Write(data)
			}
		}
	})

	return mux
}
