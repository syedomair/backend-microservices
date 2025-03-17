package container

import (
	"fmt"
	"sync"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ConnectionPoolInterface interface {
	Get() (*grpc.ClientConn, error)
	Put(conn *grpc.ClientConn)
	Close()
}

type ConnectionPool struct {
	target      string
	maxSize     int
	connections chan *grpc.ClientConn
	mu          sync.Mutex
	active      int
}

func NewConnectionPool(target string, maxSize int) (*ConnectionPool, error) {
	if maxSize <= 0 {
		return nil, fmt.Errorf("maxSize must be greater than 0")
	}
	pool := &ConnectionPool{
		target:      target,
		maxSize:     maxSize,
		connections: make(chan *grpc.ClientConn, maxSize),
		active:      0,
	}
	return pool, nil
}

func (p *ConnectionPool) Get() (*grpc.ClientConn, error) {
	select {
	case conn := <-p.connections:
		return conn, nil
	default:
		p.mu.Lock()
		defer p.mu.Unlock()
		if p.active < p.maxSize {
			conn, err := grpc.NewClient(p.target, grpc.WithTransportCredentials(insecure.NewCredentials()))
			if err != nil {
				return nil, fmt.Errorf("failed to dial: %v", err)
			}
			p.active++
			return conn, nil
		}
		// If max size is reached, wait for a connection to be returned.
		return <-p.connections, nil
	}
}

func (p *ConnectionPool) Put(conn *grpc.ClientConn) {
	select {
	case p.connections <- conn:
	default:
		// Pool is full, close the connection.
		conn.Close()
		p.mu.Lock()
		p.active--
		p.mu.Unlock()
	}
}

func (p *ConnectionPool) Close() {
	close(p.connections)
	for conn := range p.connections {
		conn.Close()
		p.mu.Lock()
		p.active--
		p.mu.Unlock()
	}
}
