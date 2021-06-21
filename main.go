package main

import (
	"GitCode/simplebank/api"
	db "GitCode/simplebank/db/sqlc"
	"GitCode/simplebank/util"
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

// db driver and postgresql connect db url
// const (
// 	dbDriver = "postgres"
// 	dbSource = "postgresql://root:dys123@localhost:5432/simple_bank?sslmode=disable"
// 	serverAddress = "0.0.0.0:8080"
// )

func main() {
	// load config
	config, err := util.LoadConfig("./")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	// conn, err := sql.Open(dbDriver, dbSource)
	conn, err := sql.Open(config.DBDriver, config.DBsource)
	if err != nil {
		log.Fatal("cannot connect to db: ", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	// err = server.Start(serverAddress)
	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}