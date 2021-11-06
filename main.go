package main

import (
	"Samples/irlog"
	"fmt"
	"sync"
	"time"
)

func lg(logger *irlog.LogData, w string, wg *sync.WaitGroup) {
	defer wg.Done()
	for k := 0; k <= 10000; k++ {
		logger.Log(1, w)
	}

}

func main() {
	logger, err := irlog.InitializeLogger("/home/irakli/go/src/Samples/Logs/log.txt", "/home/irakli/go/src/Samples/LogArchive/", 2, 1000)
	if err != nil {
		fmt.Println("errror")
	}
	wg := sync.WaitGroup{}
	//just the two of us
	wg.Add(1)
	go lg(logger, "1", &wg)
	time.Sleep(2 * time.Second)
	wg.Add(1)
	go lg(logger, "2", &wg)
	wg.Wait()
	// go lg(logger, "2")
}
