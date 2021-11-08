package main

import (
	"Samples/irlog"
)

func goru(l *irlog.LogType) {
	for i := 0; i < 100000; i++ {
		l.Log("3", 1)
	}
}

func main() {
	logger := irlog.LogType{}
	bufferSize := 100000
	fileSize := 1000000
	level := 2
	defer logger.ShutDown()
	logger.InitializeLogger("Logs/log.txt", fileSize, level, "LogArchive/", 0, bufferSize)
	for i := 0; i < 10; i++ {
		logger.Log("buffer writed succskgvldfg", 1)
	}

	for i := 0; i < 10; i++ {
		logger.Log("25555555555", 1)
	}
}
