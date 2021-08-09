package main

import (
	"fmt"
	"sync"
	"time"
)

const (
	logInfo    = "INFO"
	logWarning = "WARNING"
	logError   = "ERROR"
)

type logEntry struct {
	time      time.Time
	serverity string
	message   string
}

var wg = sync.WaitGroup{}

var logCh = make(chan logEntry, 50)

//empty struct means no memory allocation assigned to this struct
//thus, this channel is a signal channel, used to receive or send signal
var doneCh = make(chan struct{})

func main() {
	go logger()
	//alternative, we use a deferred function to close the channel
	wg.Add(2)
	defer func() {
		close(logCh)
	}()

	go func() {
		time.Sleep(1 * time.Second)
		logCh <- logEntry{time.Now(), logInfo, "app is starting"}
		wg.Done()
	}()
	go func() {
		time.Sleep(2 * time.Second)
		logCh <- logEntry{time.Now(), logInfo, "app is shutting down"}
		doneCh <- struct{}{}
		wg.Done()
	}()
	// time.Sleep(100 * time.Millisecond)

	//pass a struct{} with empty elements
	wg.Wait()
}

//the app shutdown as long as the last statement of main func() finished execution, every resources are cleaned,
//so logger() go routine is shutdown by system, therefore no deadlock occurred

func logger() {
	for {
		select {
		case entry := <-logCh:
			fmt.Printf("%v -[%v]%v\n", entry.time.Format("2006-01-02T15:04:05"), entry.serverity, entry.message)
		case <-doneCh:
			fmt.Printf("Done")
			// break

			/*alternatively, you can have a default statement to make select non-blocking
			e.g.
			defaullt:{}
			what it does is when a msg is ready on either channels, select will execute the that code path, if not, then go to deafult block

			without default statement, select is blocked forever until one msg comes in

			P.S be careful when using default statement, sometimes it may cause the high CPU usage problem because if no msg goes into the channel, then default statement will always be executed.
			*/
		}
	}
}
