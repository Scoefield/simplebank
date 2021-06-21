package main

import (
	"GitCode/simplebank/api"
	db "GitCode/simplebank/db/sqlc"
	"database/sql"
	"log"
	_ "github.com/lib/pq"
)

// db driver and postgresql connect db url
const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:dys123@localhost:5432/simple_bank?sslmode=disable"
	serverAddress = "0.0.0.0:8080"
)

func main() {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db: ", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(serverAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}