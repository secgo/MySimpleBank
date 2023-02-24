package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

var (
	testQueries *Queries
	testDB      *sql.DB
)

const (
	DBDrive  = "postgres"
	DBSource = "postgresql://root:secret@127.0.0.1/bank?sslmode=disable"
)

func TestMain(m *testing.M) {
	var err error
	testDB, err = sql.Open(DBDrive, DBSource)
	if err != nil {
		log.Fatal("cannot connect database: ", err)
	}

	testQueries = New(testDB)
	os.Exit(m.Run())
}
