package main

import (
	"Samples/irlog"
	"time"
)

func goru(l *irlog.LogType) {
	for i := 0; i < 100000; i++ {
		l.Log("3", 1)
	}
}

func main() {
	logger := irlog.LogType{}
	logger.InitializeLogger("Logs/log.txt", 100000, 2, "LogArchive/")
	for i := 0; i <= 3; i++ {
		go goru(&logger)
	}

	time.Sleep(5 * time.Second)
}
