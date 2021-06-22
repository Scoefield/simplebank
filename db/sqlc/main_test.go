package db

import (
	"GitCode/simplebank/util"
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

// db driver and postgresql connect db url
// const (
// 	dbDriver = "postgres"
// 	dbSource = "postgresql://root:dys123@localhost:5432/simple_bank?sslmode=disable"
// )

var testQueries *Queries
var testDB *sql.DB

// test main, which is the testing	of entrance.
func TestMain(m *testing.M) {
	config, err := util.LoadConfig("../../")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	
	// testDB, err = sql.Open(dbDriver, dbSource)
	testDB, err = sql.Open(config.DBDriver, config.DBsource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	testQueries = New(testDB)

	os.Exit(m.Run())
}
