package mysql

import (
	"github.com/facundo-alarcon/url-shortener/internal/database"
	"github.com/facundo-alarcon/url-shortener/urls/models"
	"log"
	"time"
)

// despues sirve si queremos hacer mocks
type UrlShortenerStorageGateway interface {
	CreateUrl(cmd *url.CreateURLCMD) (*url.Url, error)
	//deleteById(id int)
	//findById(id int) (*url.Url, error)
}

type UrlStorage struct {
	*database.MySqlClient
}

func (s *UrlStorage) CreateUrl(cmd *url.CreateURLCMD) (*url.Url, error) {
	// tx: transaccion
	// bloquear intentos de SQL injection
	tx, err := s.MySqlClient.Begin()

	if err != nil {
		log.Print("cannot create transaction")
		return nil, err
	}

	res, err := tx.Exec(`insert into urlshortener (original_url, short, short_url, created_at) 
    values (?, ?, ?, ?)`, cmd.Original, cmd.Short, cmd.ShortUrl, time.Now())

	if err != nil {
		log.Print("cannot execute statement")
		// Cerrar la transaccion si hubo un error con Rollback
		_ = tx.Rollback()
		return nil, err
	}

	// last id inserted o algun error
	id, err := res.LastInsertId()


	if err != nil {
		log.Print("cannot fetch last id")
		_ = tx.Rollback()
		return nil, err
	}

	// cierra la transaccion si termino bien
	_ = tx.Commit()

	return &url.Url{
		Id:          id,
		OriginalUrl: cmd.Original,
		Short:       cmd.Short,
		ShortUrl:    cmd.ShortUrl,
	}, nil
}