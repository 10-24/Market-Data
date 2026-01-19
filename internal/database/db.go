package internal

import (
	"database/sql"
	"log"

	"github.com/duckdb/duckdb-go/v2"
	config "github.com/yourusername/Market-Data/internal"
)

func GetDb(testMode bool) *sql.DB {
	dbPath := config.Get().DbPath
	if testMode {
		log.Println("Running in TEST mode - database changes will be simulated")
		dbPath += "?access_mode=READ_ONLY"
	}

	connector, connectorErr := duckdb.NewConnector(dbPath, nil)
	if connectorErr != nil {
		log.Fatalf("could not initialize new connector: %s", connectorErr.Error())
	}

	db := sql.OpenDB(connector)

	if pingErr := db.Ping(); pingErr != nil {
		log.Fatalf("could not ping database: %s", pingErr.Error())
	}


	return db
}
