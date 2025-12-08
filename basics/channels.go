package main

import (
	"fmt"
	"time"
)

const (
	logInfo    = "INFO"
	logWarning = "WARNING"
	logError   = "ERROR"
)

type logEntry struct {
	time     time.Time
	severity string
	message  string
}

// var wg = sync.WaitGroup{}
var logCh = make(chan logEntry, 50)
var doneCh = make(chan struct{}) // use empty struct to save some memory instead of bool (need to allocate some memory)

func ChannelExample1() {
	ch := make(chan int, 50)
	wg.Add(2)
	go func(ch <-chan int) {
		// for i := range ch {
		// 	fmt.Println(i)
		// }
		for {
			if i, ok := <-ch; ok {
				fmt.Println(i)
			} else {
				break
			}
		}
		wg.Done()
	}(ch)
	go func(ch chan<- int) {
		ch <- 42
		ch <- 27
		close(ch)
		wg.Done()
	}(ch)
	wg.Wait()
}

func ChannelExample2() {
	go logger()
	// defer func() {
	// 	close(logCh)
	// }()
	logCh <- logEntry{time.Now(), logInfo, "App is starting"}
	logCh <- logEntry{time.Now(), logInfo, "App is shutting down"}
	time.Sleep(100 * time.Microsecond)
	doneCh <- struct{}{}
}

func logger() {
	// for entry := range logCh {
	// 	fmt.Printf("%v - [%v]%v\n", entry.time.Format("2006-01-02T15:04:05"), entry.severity, entry.message)
	// }
	for {
		select {
		case entry := <-logCh:
			fmt.Printf("%v - [%v]%v\n", entry.time.Format("2006-01-02T15:04:05"), entry.severity, entry.message)
		case <-doneCh:
			fmt.Println("done")
			break
		}
	}
}
