package tcp

import (
	"awesomeProject/go-tcp/interface/tcp"
	"awesomeProject/go-tcp/lib/logger"
	"context"
	"fmt"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

type Config struct {
	Address string
}

func ListenAndServeWithSignal(cfg *Config, handler tcp.Handler) error {

	closeChan := make(chan struct{})
	sigCh := make(chan os.Signal)
	signal.Notify(sigCh, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	go func() {
		sig := <-sigCh
		switch sig {
		case syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			closeChan <- struct{}{}
		}
	}()
	listener, err := net.Listen("tcp", cfg.Address)
	if err != nil {
		return err
	}
	//cfg.Address = listener.Addr().String()
	logger.Info(fmt.Sprintf("bind: %s, start listening...", cfg.Address))
	ListenAndServe(listener, handler, closeChan)
	return nil

}
func ListenAndServe(listener net.Listener, handler tcp.Handler,
	closeChan <-chan struct{}) {

	// listen signal
	errCh := make(chan error, 1)
	defer close(errCh)
	go func() {
		select {
		case <-closeChan:
			logger.Info("get exit signal")
		case er := <-errCh:
			logger.Info(fmt.Sprintf("accept error: %s", er.Error()))
		}
		logger.Info("shutting down...")
		_ = listener.Close() // listener.Accept() will return err immediately
		_ = handler.Close()  // close connections
	}()

	//关闭 tcp 连接 及 handler
	defer func() {
		// 错误 忽略
		_ = listener.Close()
		_ = handler.Close()
	}()
	var waitDone sync.WaitGroup
	ctx := context.Background()
	for true {
		conn, err := listener.Accept()
		if err != nil {
			break
		}
		waitDone.Add(1)
		logger.Info("accepted link")
		go func() {
			// == java  finnaly{}
			defer func() {
				waitDone.Done()
			}()
			handler.Handler(ctx, conn)
		}()
	}
	waitDone.Wait()
}
