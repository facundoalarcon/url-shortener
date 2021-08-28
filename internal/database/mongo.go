package database

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"time"
)

type Mongo struct {
	Client *mongo.Client
}

func NewMongoClient(user, password, host, port string) *Mongo {
	connString := fmt.Sprintf("mongodb://%s:%s@%s:%s/?authSource=admin", user, password, host, port)
	//connString := fmt.Sprintf("mongodb://%s:27017", host)
	client, err := mongo.NewClient(options.Client().ApplyURI(connString))

	if err != nil {
		log.Printf("Error: cannot connect with mongo %s", err.Error())
		panic(err)
	}

	// Connect to Mongo
	// se ejecuta en un contexto Background
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)

	if err != nil {
		log.Printf("Error: cannot connect with mongo %s", err.Error())
		panic(err)
	}

	// Test your connection
	err = client.Ping(context.Background(), readpref.Primary())

	if err != nil {
		log.Printf("Error: cannot PING the mongo server %s", err.Error())
		panic(err)
	}

	// esto es un puntero a Mongo porque esta funcion tiene que devolver *Mongo que es un *mongo.Client
	// entonces la unica forma de hacer eso es devolviendo un puntero a la dir de memoria donde se aloca
	// nuestra variable client que guarda ese valor
	return &Mongo{client}
}