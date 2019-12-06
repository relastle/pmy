package pmy

import (
	"log"
	"os"
	"path"
	"time"
)

// MeasureElapsedTime measures elapsed time given
// a started time.
func MeasureElapsedTime(start time.Time, name string) {
	logPath := os.ExpandEnv(defaultLogPath)
	err := os.MkdirAll(path.Dir(logPath), os.ModePerm)
	if err != nil {
		return
	}
	f, err := os.OpenFile(logPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()

	elapsed := time.Since(start)
	log.SetOutput(f)
	log.Printf("%s took %v", name, elapsed)
}
