package db

import (
	"database/sql"
	"os"
	"testing"

	_ "github.com/lib/pq"
	"github.com/rohitvpatil0810/go-url-shortener-api/internal/config"
	"github.com/rohitvpatil0810/go-url-shortener-api/pkg/logger"
)

var (
	dbDriver = "postgres"
	dbSource string
)

var testQueries *Queries

func TestMain(m *testing.M) {
	cfg, err := config.LoadConfig("../../.env")
	if err != nil {
		logger.Logger.Error(
			"Error loading config",
			"error", err,
		)
	}
	dbSource = cfg.PostgresConfig.Url

	conn, err := sql.Open(dbDriver, dbSource)

	if err != nil {
		logger.Logger.Error("Error opening connection", "error", err)
	}

	testQueries = New(conn)

	os.Exit(m.Run())
}
