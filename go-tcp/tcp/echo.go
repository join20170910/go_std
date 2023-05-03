package tcp

import (
	"awesomeProject/go-tcp/lib/logger"
	"awesomeProject/go-tcp/lib/sync/atomic"
	"awesomeProject/go-tcp/lib/sync/wait"
	"bufio"
	"context"
	"io"
	"net"
	"sync"
	"time"
)

type EchoClient struct {
	Conn    net.Conn
	Waiting wait.Wait
}

func MakeHandler() *EchoHandler {
	return &EchoHandler{}
}

type EchoHandler struct {
	activeConn sync.Map
	closing    atomic.Boolean
}

func (e *EchoClient) Close() error {
	e.Waiting.WaitWithTimeout(10 * time.Second)
	err := e.Conn.Close()
	return err
}

func (handler *EchoHandler) Handler(ctx context.Context, conn net.Conn) {
	if handler.closing.Get() {
		_ = conn.Close()
	}
	client := &EchoClient{
		Conn: conn,
	}
	handler.activeConn.Store(client, struct {
	}{})
	reader := bufio.NewReader(conn)
	for true {
		msg, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				logger.Info("Connection close")
				handler.activeConn.Delete(client)
			} else {
				logger.Warn(err)
			}
			return
		}
		//处理业务数据开始
		client.Waiting.Add(1)
		b := []byte(msg)
		_, _ = conn.Write(b)
		client.Waiting.Done()
		//处理业务 结束

	}

}

func (handler *EchoHandler) Close() error {
	logger.Info("handler shutting down")
	handler.closing.Set(true)
	handler.activeConn.Range(func(key, value any) bool {
		client := key.(*EchoClient)
		_ = client.Conn.Close()
		return true
	})
	return nil
}
