package db

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
)

const (
	connString = "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable"
)

var testQueries *Queries

func TestMain(m *testing.M) {
	connPool, err := pgxpool.New(context.Background(), connString)
	if err != nil {
		log.Fatal("cannot connect to db", err)
	}
	defer connPool.Close()

	testQueries = New(connPool)

	os.Exit(m.Run())
}
