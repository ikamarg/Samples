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
	bufferSize := 1000
	fileSize := 10000

	asyncLogger := irlog.LogType{}
	defer asyncLogger.ShutDown()
	asyncLogger.InitializeLogger("Logs/asyncLog.txt", fileSize, irlog.AsynchronousLogLevel, "AsynchronousLogArchive/", irlog.AsynchronousLog, bufferSize)

	syncLogger := irlog.LogType{}
	defer syncLogger.ShutDown()
	syncLogger.InitializeLogger("Logs/syncLog.txt", fileSize, irlog.SynchronousLogLevel, "SynchronousLogArchive/", irlog.SynchronousLog, bufferSize)

	for i := 0; i < 100; i++ {
		asyncLogger.Log("This Is Production Log For Debug", irlog.Debug)
		asyncLogger.Log("This Is Production Log For Info", irlog.Info)
		asyncLogger.Log("This Is Production Log For Warning", irlog.Warning)
		asyncLogger.Log("This Is Production Log For Error", irlog.Error)
		asyncLogger.Log("This Is Production Log For Fatal", irlog.Fatal)
	}

	for i := 0; i < 100; i++ {
		syncLogger.Log("This Is Developement Log For Debug", irlog.Debug)
		syncLogger.Log("This Is Developement Log For Info", irlog.Info)
		syncLogger.Log("This Is Developement Log For Warning", irlog.Warning)
		syncLogger.Log("This Is Developement Log For Error", irlog.Error)
		syncLogger.Log("This Is Developement Log For Fatal", irlog.Fatal)
	}
}
