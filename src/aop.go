package pmy

import (
	"log"
	"os"
	"time"
)

func measureElapsedTime(start time.Time, name string) {
	f, err := os.OpenFile(os.ExpandEnv(defaultLogPath), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()

	elapsed := time.Since(start)
	log.SetOutput(f)
	log.Printf("%s took %v", name, elapsed)
}
