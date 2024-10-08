package main

import (
	"bufio"
	"fmt"
	"net"
	"net/http"
	"strings"
)

func main() {
	http.Handle("GET /test", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Trailer", "my-trailer,my-trailer2")
		w.Write([]byte("Hello"))
		w.Header().Add("asdasd", "asdfasdf") // won't work, has to put it before any write
		w.Write([]byte(strings.Repeat("!", 3000)))
		w.Header().Add("my-trailer", "asdfasdf")
		w.Header().Add("my-trailer2", "22asdfasdf")
	}))
	go http.ListenAndServe(":8000", nil)

	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	_, err = conn.Write([]byte("GET /test HTTP/1.1\r\nHost: localhost:8000\r\n\r\n"))
	if err != nil {
		panic(err)
	}
	// buf := make([]byte, 1024)
	reader := bufio.NewReader(conn)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("------breakk------")
			break // Break the loop if there's an error or end of response
		}
		fmt.Print(line) // Print each chunk of the body
		if line == "\r\n" {
			fmt.Print("\r\n: ", line) // Print each chunk of the body
			break
		}
	}
	fmt.Println("reader.Buffered()", reader.Buffered())
	// var buf []byte
	// for {
	// 	n, err := conn.Read(buf)
	// 	if err != nil {
	// 		panic(err)
	// 		break
	// 	}
	// 	fmt.Println("n:", n)
	// 	fmt.Println(string(buf[:n]))
	// }
	// Read the trailers (headers sent after the body)
	for {
		trailer, err := reader.ReadString('\n')
		if err != nil || trailer == "\r\n" {
			break // Exit if there's no more data or empty line signaling end of trailers
		}
		fmt.Print(trailer)
	}
}
