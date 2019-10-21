package pmy

import (
	"log"
	"os"
	"path"
	"time"
)

func measureElapsedTime(start time.Time, name string) {
	logPath := os.ExpandEnv(defaultLogPath)
	os.MkdirAll(path.Dir(logPath), os.ModePerm)
	f, err := os.OpenFile(logPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()

	elapsed := time.Since(start)
	log.SetOutput(f)
	log.Printf("%s took %v", name, elapsed)
}
