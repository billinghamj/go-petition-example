package handler

import (
	"net/http"

	"github.com/K-Phoen/negotiation"
	"github.com/unrolled/render"
)

func getOutputType(r *http.Request) (string, error) {
	header := r.Header.Get("Accept")

	if header == "" {
		return "application/json", nil
	}

	alt, err := negotiation.NegotiateAccept(header, []string{
		"application/json",
	})

	if err != nil {
		return "", err
	}

	return alt.Name, nil
}

var renderer = render.New()

func writeOutput(w http.ResponseWriter, status int, format string, output interface{}) error {
	switch format {
	default:
		panic("unpossibru!")

	case "application/json":
		return renderer.JSON(w, status, output)
	}
}
