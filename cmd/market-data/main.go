package main

import (
	"database/sql"
	"flag"
	"log"
	"slices"

	internal "github.com/yourusername/Market-Data/internal/database"
	"github.com/yourusername/Market-Data/internal/pipeline"
	"github.com/yourusername/Market-Data/internal/types"
)

func main() {
	test := flag.Bool("test", false, "Run in test mode (read-only, simulates database changes)")
	flag.Parse()

	db := internal.GetDb(*test)
	defer db.Close()

	instrumentIds, err := getInstrumentIds(db)
	if err != nil {
		log.Fatalf("Failed to get instrument Ids: %v", err)
	}
	log.Printf("Loaded %d instruments", len(instrumentIds))

	tradeStream := pipeline.WatchTrades(instrumentIds)
	candleStream := pipeline.CreateCandles(tradeStream)
	pipeline.WriteCandles(candleStream, db)
}

func getInstrumentIds(db *sql.DB) ([]types.InstrumentId, error) {

	rows, err := db.Query(`
		SELECT exchange_id, base, quote
		FROM data_subscription
		ORDER BY exchange_id, base, quote
	`, nil)
	if err != nil {
		return nil, err
	}

	var instrumentIds []types.InstrumentId

	for rows.Next() {
		var exchange string
		var base string
		var quote string

		if err := rows.Scan(&exchange, &base, &quote); err != nil {
			return nil, err
		}

		instrumentId, parseErr := types.NewInstrumentId(exchange, base, quote)
		if parseErr != nil {
			return nil, err
		}
		if slices.Contains(instrumentIds,instrumentId){
			log.Println("Found duplicate data_subscription for" + instrumentId.String())
			continue
		}
		instrumentIds = append(instrumentIds, instrumentId)
	}

	return instrumentIds, rows.Err()
}
