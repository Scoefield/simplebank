package db

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"os"
	"testing"
)

// db driver and postgresql connect db url
const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:dys123@localhost:5432/simple_bank?sslmode=disable"
)

var testQueries *Queries

// test main, which is the testing	of entrance.
func TestMain(m *testing.M) {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	testQueries = New(conn)

	os.Exit(m.Run())
}
