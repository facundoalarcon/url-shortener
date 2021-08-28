package mysql

import url "github.com/facundo-alarcon/url-shortener/urls/models"

func (s *UrlCreateGtw) Create(cmd *url.CreateURLCMD) (*url.Url, error) {
	return s.CreateUrl(cmd)
}

