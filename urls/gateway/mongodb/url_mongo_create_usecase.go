package mongodb

import url "github.com/facundo-alarcon/url-shortener/urls/models"

func (u *UrlMongoCreateGtw) AddUrl(cmd *url.CreateURLCMD) (string, error) {
	return u.Add(cmd)
}