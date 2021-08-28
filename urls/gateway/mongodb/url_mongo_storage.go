package mongodb

import (
	"context"
	"github.com/facundo-alarcon/url-shortener/internal/database"
	url "github.com/facundo-alarcon/url-shortener/urls/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	//"strconv"
	"time"
)

type UrlMongoStorage interface {
	Add(cmd *url.CreateURLCMD) (string, error)
}

type UrlStg struct {
	*database.Mongo
}

func (s *UrlStg) Add(cmd *url.CreateURLCMD) (string, error) {
	collections := s.Client.Database("urlshortenerDB").Collection("urls")

	// pasar contexto donde mongo esta corriendo
	// documento que queremos insertar con bson (binary json)
	// bson.D = bson ordenado, va a estar ordenado al momento de la insercion: es un mapa de key values
	res, err := collections.InsertOne(context.Background(),
		bson.D{
		{"original_url", cmd.Original},
		{"short", cmd.Short},
		{"short_url", cmd.ShortUrl},
		{"created_at", time.Now() },
		{"url_id", cmd.UrlId },
		})

	if err != nil {
		log.Printf("connot insert url")
		panic(err)
		//return "", err
	}

	id := res.InsertedID.(primitive.ObjectID)
	return id.String(), nil

}
