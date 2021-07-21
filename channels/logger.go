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
	time      time.Time
	serverity string
	message   string
}

var logCh = make(chan logEntry, 50)

func main() {
	go logger()
	//alternative, we use a deferred function to close the channel
	defer func() {
		close(logCh)
	}()

	logCh <- logEntry{time.Now(), logInfo, "app is starting"}
	logCh <- logEntry{time.Now(), logInfo, "app is shutting down"}
	time.Sleep(100 * time.Millisecond)

}

//the app shutdown as long as the last statement of main func() finished execution, every resources are cleaned,
//so logger() go routine is shutdown by system, therefore no deadlock occurred

func logger() {
	for entry := range logCh {
		fmt.Printf("%v -[%v]%v\n", entry.time.Format("2006-01-02T15:04:05"), entry.serverity, entry.message)
	}
}
