package main

import "fmt"

type Subject struct {
	Context   string
	Observers []Observer
}

type Observer interface {
	Notify(s Subject)
}

type Reader struct {
	Name string
}

func (r Reader) Notify(s Subject) {
	fmt.Printf("Reader %v has notified, s.Context: %v\n", r.Name, s.Context)
}

func (s Subject) Update() {
	for _, observer := range s.Observers {
		observer.Notify(s)
	}
}

func ObserverDemo() {
	s := new(Subject)
	s.Context = "aaa"
	r1 := new(Reader)
	r1.Name = "r1"
	r2 := new(Reader)
	r2.Name = "r2"
	r3 := new(Reader)
	r3.Name = "r3"
	s.Observers = append(s.Observers, r1)
	s.Observers = append(s.Observers, r2)
	s.Observers = append(s.Observers, r3)
	s.Update()
	s.Context = "bbb"
	s.Update()
}
