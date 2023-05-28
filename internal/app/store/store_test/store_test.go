package store_test

import (
	"os"
	"testing"
)

var (
	databaseURL  string
	databaseType string
)

func TestMain(m *testing.M) {
	databaseURL = os.Getenv("DATABASE_URL")
	databaseType = os.Getenv("DATABASE_TYPE")

	if databaseURL == "" {
		databaseURL = "host=localhost dbname=test-RESTAPISample user=devUser password=devUser sslmode=disable"
	}

	if databaseType == "" {
		databaseType = "postgres"
	}

	os.Exit(m.Run())
}
