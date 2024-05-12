package pool

import (
	"fmt"
	"sync"
)

type Connection struct {
	ID int
}

type Pool struct {
	mu      sync.Mutex
	conns   chan *Connection
	factory func() (*Connection, error)
	closed  bool
}

func NewPool(maxSize int, factory func() (*Connection, error)) (*Pool, error) {
	if maxSize <= 0 {
		return nil, fmt.Errorf("maxSize must be greater than 0")
	}
	pool := &Pool{
		conns:   make(chan *Connection, maxSize),
		factory: factory,
	}
	for i := 0; i < maxSize; i++ {
		conn, err := factory()
		if err != nil {
			pool.Close()
			return nil, fmt.Errorf("error creating connection: %v", err)
		}
		pool.conns <- conn
	}
	return pool, nil
}

func (p *Pool) Get() (*Connection, error) {
	// p.mu.Lock()
	// defer p.mu.Unlock()

	if p.closed {
		return nil, fmt.Errorf("pool is closed")
	}
	conn := <-p.conns
	fmt.Println("Reusing existing connection", conn.ID)
	return conn, nil
}

func (p *Pool) Release(conn *Connection) error {
	p.mu.Lock()
	defer p.mu.Unlock()
	if p.closed {
		return fmt.Errorf("pool is closed")
	}
	p.conns <- conn
	fmt.Println("connection discarded", conn.ID)
	return nil
}

func (p *Pool) Close() {
	p.mu.Lock()
	defer p.mu.Unlock()
	if p.closed {
		return
	}
	p.closed = true
	close(p.conns)
	for conn := range p.conns {
		fmt.Println("Closing connection", conn.ID)
	}
}
