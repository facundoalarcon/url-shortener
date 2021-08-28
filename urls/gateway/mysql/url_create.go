package mysql

import (
	"github.com/facundo-alarcon/url-shortener/internal/database"
	url "github.com/facundo-alarcon/url-shortener/urls/models"
)

type UrlCreateGateway interface {
	Create(cmd *url.CreateURLCMD) (*url.Url, error)
}

type UrlCreateGtw struct {
	UrlShortenerStorageGateway
}

func NewUrlCreateGateway(client *database.MySqlClient) UrlCreateGateway {
	return &UrlCreateGtw{&UrlStorage{client}}
}
