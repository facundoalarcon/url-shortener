package web

import (
	"encoding/json"
	"github.com/facundo-alarcon/url-shortener/internal/database"
	"github.com/facundo-alarcon/url-shortener/urls/gateway/mongodb"
	url "github.com/facundo-alarcon/url-shortener/urls/models"
	"net/http"
)

type UrlShortenerMongoHandler struct {
	mongodb.UrlCreateMongoGateway
}

func NewUrlShortenerMongoHandler(mongo *database.Mongo) *UrlShortenerMongoHandler {
	return &UrlShortenerMongoHandler{mongodb.NewUrlCreateMongoGateway(mongo)}
}

func (h *UrlShortenerMongoHandler) AddUrlShortenerMongoHandler(w http.ResponseWriter, r *http.Request) {
	cmd := params(r)
	res, err := h.AddUrl(cmd)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(res))
}

// parsear lo que nos den desde la web
func params(r *http.Request) *url.CreateURLCMD {
	var cmd url.CreateURLCMD

	_ = json.NewDecoder(r.Body).Decode(&cmd)

	return &cmd
}