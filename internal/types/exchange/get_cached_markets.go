package exchange

import (
	"encoding/gob"
	"log"
	"os"
	"path/filepath"

	config "github.com/yourusername/Market-Data/internal"
)



func getCachedMarkets(exchangeId ExchangeId) (markets,bool) {
	// Register types using interface{} instead of any for gob compatibility
	// gob stores type information differently for interface{} vs any
	gob.Register(map[string]interface{}(nil))
	gob.Register([]any(nil))

	exchangeIdstr := exchangeId.String()
	filePath := filepath.Join(config.GetConfigDirPath(),"markets", exchangeIdstr+".gob")
	file, err := os.Open(filePath)
	if err != nil { // File not found
		return nil,false
	}
	defer file.Close()

	decoder := gob.NewDecoder(file)
	var loadedMarkets map[string]interface{}
	if err := decoder.Decode(&loadedMarkets); err != nil {
		log.Printf("Warning: Failed to decode %s: %v\n", filePath, err)
		return nil,false
	}

	// Fixes Symbol not found error. Apparently gob sometimes disregards type information, and decodes to its own taste 🤷
	ccxtMarkets := make(map[string]interface{})
	for k, v := range loadedMarkets {
		ccxtMarkets[k] = v
	}

	return ccxtMarkets, true
}


type markets = map[string]interface{}