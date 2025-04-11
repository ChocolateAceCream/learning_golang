package main

import (
	"bytes"
	"fmt"
	"sync"
)

var bufferPool = sync.Pool{
	New: func() interface{} {
		return new(bytes.Buffer)
	},
}

func main() {
	buf := bufferPool.Get().(*bytes.Buffer)
	defer bufferPool.Put(buf)
	buf.Reset()
	buf.WriteString("Hello world")
	buf.WriteString("\n!!!")
	fmt.Println(buf.String())
}
