package main

import (
	"fmt"
	"github.com/facundo-alarcon/url-shortener/internal/database"
	"github.com/facundo-alarcon/url-shortener/server"
	"github.com/facundo-alarcon/url-shortener/urls/web"
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate"
	migration "github.com/golang-migrate/migrate/database/mysql"
	_ "github.com/golang-migrate/migrate/source/file"
	"log"
	"os"
)

const (
	migrationsRootFolder     = "file://migrations"
	migrationsScriptsVersion = 1
)

func main() {

	mysqlClient:= newMysqlConnection()
	//mongoClient := database.NewMongoClient(os.Getenv("MONGO_HOST"))
	mongoClient :=newMongoDbConnection()

	//newServer(":3000")
	handler := web.NewCreateUrlShortenerHandler(mysqlClient)
	mongoHandler := web.NewUrlShortenerMongoHandler(mongoClient)

 	mux := server.Routes(handler, mongoHandler)
	//fmt.Printf("%v",mux)
	newServer := server.NewServer(mux)
	newServer.Run()
}

func newMongoDbConnection() *database.Mongo {
	dbUser := os.Getenv("MONGO_INITDB_ROOT_USERNAME")
	dbPass := os.Getenv("MONGO_INITDB_ROOT_PASSWORD")
	//dbName := os.Getenv("MONGO_DATABASE")
	dbHost := os.Getenv("MONGO_HOST")
	dbPort := os.Getenv("MONGO_PORT")

	conn := database.NewMongoClient(dbUser, dbPass, dbHost, dbPort)

	return conn
}

func newMysqlConnection() *database.MySqlClient{

	dbUser := os.Getenv("MYSQL_USER")
	dbPass := os.Getenv("MYSQL_PASSWORD")
	dbName := os.Getenv("MYSQL_DATABASE")
	dbHost := os.Getenv("MYSQL_HOST")
	dbPort := os.Getenv("MYSQL_PORT")

	//db, _ := sql.Open("mysql", "user:password@tcp(host:port)/dbname?multiStatements=true")
	dbSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true",
		dbUser, dbPass, dbHost, dbPort, dbName,
	)

	//log.Printf("%v",dbSource)

	client, err := database.NewSqlClient(dbSource)

	if err != nil {
		log.Printf("Error: %v", err.Error())
	} else {
		dbMigrate(client, dbName)
		log.Printf("Done!")
	}

	return client
}

func dbMigrate(client *database.MySqlClient, dbName string){
	// create database
	driver, _ := migration.WithInstance(client.DB, &migration.Config{})
	m, err := migrate.NewWithDatabaseInstance(
		migrationsRootFolder,
		dbName,
		driver,
	)

	if err != nil {
		//logs.Log().Error(err.Error())
		log.Printf("%v",err.Error())
		return
	}

	current, _, _ := m.Version()
	//logs.Log().Infof("current migrations version in %d", current)
	log.Printf("current migrations version in %d", current)
	err = m.Migrate(migrationsScriptsVersion)
	if err != nil && err.Error() == "no change" {
		//logs.Log().Info("no migration needed")
		log.Printf("no migration needed")
	}
}
