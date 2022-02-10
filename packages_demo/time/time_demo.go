package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	token := RandToken(10)
	fmt.Println("------token----", token)

	fmt.Println("------time----", time.Now())
	fmt.Println("------time.UnixNano()----", time.Now().UnixNano())
	ticker := time.NewTicker(1 * time.Second) // send time to chanel for every tick
	timer := time.NewTimer(10 * time.Second)  // send time to chanel once timer run out
	count := 0
	for {
		select {
		case i := <-ticker.C:
			fmt.Println("---i----", i) //---i---- 2022-02-09 14:37:28.9475511 +0800 CST m=+78.000327101
			count++

			// timer.Stop() will stop timer from firing
			// if count > 5 {
			// 	r := timer.Stop()
			// 	fmt.Println("stoped timer? ", r)
			// }
		case j := <-timer.C:
			fmt.Println("---j----", j) //---j---- 2022-02-09 14:37:28.9475511 +0800 CST m=+78.000327101

		}
	}
}

func RandToken(n int) string {
	rand.Seed(time.Now().UnixNano())

	var letters = []rune("abcdefghijklmnopqrstu1234567890vwxyz-_:!$ABCDEFGHIJKLMNOPQRSTUVWXYZ")

	b := make([]rune, n)

	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}

	return string(b)
}
