package main

import (
	"bytes"
	"fmt"
	"io"
)

func main() {
	var wc WriterCloser = NewBufferedWriterClose()
	wc.Write([]byte("Z123456789"))
	// wc2.Write([]byte("wc2 hello, this is a test"))
	// wc2.Close()
	fmt.Println("buffer is:", wc.(*BufferedWriterCloser).buffer)
	wc.Close()
	fmt.Println("buffer is:", wc.(*BufferedWriterCloser).buffer)

	r, ok := wc.(io.Reader)
	if !ok {
		fmt.Println("---not ok---")
	} else {
		fmt.Println("r", r)
	}

}

type Writer interface {
	Write([]byte) (int, error)
}

type Closer interface {
	Close() error
}

type WriterCloser interface {
	Writer
	Closer
}

type BufferedWriterCloser struct {
	buffer *bytes.Buffer
}

func (bwc *BufferedWriterCloser) Write(data []byte) (int, error) {
	n, err := bwc.buffer.Write(data) // n is the length of buffer
	fmt.Println("------after write----")
	fmt.Println("length of buffer: ", n)
	fmt.Println("------after read ----")
	k := make([]byte, 1)
	bwc.buffer.Read(k)
	fmt.Println(k) //read first byte of buffer and assgin the value to k
	fmt.Println(string(k))

	if err != nil {
		return 0, err
	}
	v := make([]byte, 8)
	for bwc.buffer.Len() > 8 {
		_, err := bwc.buffer.Read(v)
		if err != nil {
			return 0, err
		}
		_, err = fmt.Println(string(v))
		if err != nil {
			return 0, err
		}
	}
	return n, nil
}

/*
when we define type and assign methods to them, each of those types has what's called methods set.
when you working with types directly, the method set is all of the methods regardless of the receiver types
associated with that.

when implement interface with concrate value (such as struct), the method set should have a
value receiver, otherwise the method set will be incomplete.

when implement interface with a pointer value, such as *BufferedWriterCloser in
func (bwc *BufferedWriterCloser) Close() error{}
the method set for a pointer is the sum of the value receiver and pointer receiver

in other words:
when implement interface with a concrate value, interface must have value receivers
when implement interface with a pointer, interface will have both value receivers and pointer receivers, so
receiver type does not matter anymore


e.g.
implement with a pointer receiver:
func (bwc *BufferedWriterCloser) Close() error {}
then, when using interface, we must define a pointer receiver
e.g. var bwc WriterCloser = &BufferedWriterCloser{}

e.g.
implement with a value receiver:
func (bwc BufferedWriterCloser) Close() error {}
then, when using interface, we can either define a pointer receiver or value receiver:
var bwc WriterCloser = &BufferedWriterCloser{}
or
var bwc WriterCloser = BufferedWriterCloser{}
both work




Notice here we use pointer receiver *BufferedWriterCloser
*/
func (bwc *BufferedWriterCloser) Close() error {
	for bwc.buffer.Len() > 0 {
		data := bwc.buffer.Next(8)
		_, err := fmt.Println(string(data))
		if err != nil {
			return err
		}
	}
	return nil
}

func NewBufferedWriterClose() *BufferedWriterCloser {
	return &BufferedWriterCloser{
		buffer: bytes.NewBuffer([]byte{}),
	}
}
