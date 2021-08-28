package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

type MySqlClient struct {
	*sql.DB
}

func NewSqlClient(source string) (*MySqlClient, error) {
	log.Printf("%s", source)

	db, err := sql.Open("mysql", source)

	if err != nil {
		//logs.Log().Errorf("cannot create db tentat: %s", err.Error())
		log.Printf("cannot create db tentat: %s", err.Error())
		panic(err)
	}

	if err := db.Ping(); err != nil {
		//logs.Log().Warn("cannot connect to mysql!")
		log.Printf("cannot connect to mysql!")
		log.Fatal()
	}

	return &MySqlClient{db}, nil
}

func (c *MySqlClient) ViewStats() sql.DBStats{
	return c.Stats()
}