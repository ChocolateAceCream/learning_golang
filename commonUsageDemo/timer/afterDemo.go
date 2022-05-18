package main

import (
	"log"
	"time"
)

func main() {
	AfterDemo()
	AfterFuncDemo()
}

func AfterDemo() {
	log.Println("AfterDemo start ")
	<-time.After(5 * time.Second)
	log.Println("AfterDemo end")
}
func AfterFuncDemo() {
	log.Println("AfterFuncDemo start")
	time.AfterFunc(5*time.Second, func() {
		AfterDemo()
		log.Println("AfterFuncDemo end")
	})
	time.Sleep(20 * time.Second)

}
