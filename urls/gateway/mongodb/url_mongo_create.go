package mongodb

import (
	"github.com/facundo-alarcon/url-shortener/internal/database"
	url "github.com/facundo-alarcon/url-shortener/urls/models"
)

// Mongo
type UrlCreateMongoGateway interface {
	AddUrl(cmd *url.CreateURLCMD) (string, error)
}

type UrlMongoCreateGtw struct {
	UrlMongoStorage
}

// Mongo same interface of	 MySql
func NewUrlCreateMongoGateway(mongoClient *database.Mongo) UrlCreateMongoGateway {
	return &UrlMongoCreateGtw{&UrlStg{mongoClient}}
}
