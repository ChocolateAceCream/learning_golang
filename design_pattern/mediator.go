package main

import "fmt"

type Train interface {
	Arrive()
	Depart()
}

type Train1 struct {
	Manager *Manager
}

func (t *Train1) Arrive() {
	fmt.Println("train1 want to arrive")
	if t.Manager.FreeSlot {
		t.Manager.FreeSlot = false
		fmt.Println("train1 has arrived")
		return
	}
	fmt.Println("train1 has to wait")
	t.Manager.Queue = append(t.Manager.Queue, t)
}

func (t *Train1) Depart() {
	fmt.Println("train1 has Depart")
	t.Manager.FreeSlot = true
	t.Manager.NotifyDepart()
}

type Train2 struct {
	Manager *Manager
}

func (t *Train2) Arrive() {
	fmt.Println("train2 want to arrive")
	if t.Manager.FreeSlot {
		t.Manager.FreeSlot = false
		fmt.Println("train2 has arrived")
		return
	}
	fmt.Println("train2 has to wait")
	t.Manager.Queue = append(t.Manager.Queue, t)
}

func (t *Train2) Depart() {
	fmt.Println("train2 has Depart")
	t.Manager.FreeSlot = true
}

type Manager struct {
	Queue    []Train
	FreeSlot bool
}

func (m *Manager) NotifyDepart() {
	if len(m.Queue) > 0 {
		m.FreeSlot = true
		m.Queue[0].Arrive()
		m.Queue = m.Queue[1:]
	}
}

func MediatorDemo() {
	manager := &Manager{
		FreeSlot: true,
	}
	train1 := &Train1{
		Manager: manager,
	}

	train2 := &Train2{
		Manager: manager,
	}
	train1.Arrive()
	train2.Arrive()
	train1.Depart()
}
