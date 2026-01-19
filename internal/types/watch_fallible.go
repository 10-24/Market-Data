package types

import (
	"log"
	"time"
)
// Wraps Connect for usage with wait group
func WatchFallibleFn(connector func() error, logId string) func() {
	return func () {
		WatchFallible(connector,logId)
	}
}

func WatchFallible(connector func() error, logId string)  {
	const maxTries uint8 = 128
	const base = time.Duration(2) * time.Second
	const minAcceptableRuntime = time.Duration(2) * time.Minute
	
	var tries uint8 = 0;
	tryConnectTime := time.Now()
	
	for tries <= maxTries {
		tries += 1
		tryConnectTime = time.Now()
		
		err := connector();
		if err == nil {
			return
		}
		log.Printf("Error: Connection %s failed with err: %v. Connection try (%d/%d) \n",logId,err,tries,maxTries)
		
		
		if time.Since(tryConnectTime) >= minAcceptableRuntime {
			tries = 0
			continue
		}
		
		time.Sleep(pow(base,tries))
	}
	
	log.Printf("Error: Exceeded maxTries this is unexpected behavior. %d^%d sec is intended to be arbitrarily large such that this fn never returns",base,maxTries)
	
}



func pow(base time.Duration, exponent uint8) time.Duration {
	n := time.Duration(1)
	for range exponent {
		n *= base
	}
	return n
}

