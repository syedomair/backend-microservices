package container

import (
	"context"
	"net"
	"testing"

	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
)

const bufSize = 1024 * 1024

var lis *bufconn.Listener

func bufDialer(context.Context, string) (net.Conn, error) {
	return lis.Dial()
}

func startBufconnServer() {
	lis = bufconn.Listen(bufSize)
	s := grpc.NewServer()
	go func() {
		if err := s.Serve(lis); err != nil {
			panic(err)
		}
	}()
}

func TestNewConnectionPool(t *testing.T) {
	pool, err := NewConnectionPool("bufnet", 2)
	assert.NoError(t, err)
	assert.NotNil(t, pool)
}

func TestConnectionPool_GetAndPut(t *testing.T) {
	startBufconnServer()
	pool, _ := NewConnectionPool("bufnet", 2)

	conn1, err := pool.Get()
	assert.NoError(t, err)
	assert.NotNil(t, conn1)

	pool.Put(conn1)
	conn2, err := pool.Get()
	assert.NoError(t, err)
	assert.Equal(t, conn1, conn2, "Expected the same connection to be reused")
	pool.Put(conn2)
}

func TestConnectionPool_MaxSize(t *testing.T) {
	startBufconnServer()
	pool, _ := NewConnectionPool("bufnet", 2)

	conn1, _ := pool.Get()
	conn2, _ := pool.Get()

	done := make(chan bool)
	go func() {
		conn3, _ := pool.Get()
		assert.NotNil(t, conn3)
		done <- true
	}()

	pool.Put(conn1)
	<-done

	pool.Put(conn2)
}

func TestConnectionPool_Close(t *testing.T) {
	startBufconnServer()
	pool, _ := NewConnectionPool("bufnet", 2)
	conn1, _ := pool.Get()
	conn2, _ := pool.Get()
	pool.Put(conn1)
	pool.Put(conn2)
	pool.Close()
	assert.Equal(t, 0, pool.active)
}
