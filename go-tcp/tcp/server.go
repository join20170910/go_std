package tcp

import (
	"awesomeProject/go-tcp/interface/tcp"
	"awesomeProject/go-tcp/lib/logger"
	"context"
	"net"
)

type Config struct {
	Address string
}

func ListenAndServeWithSignal(cfg *Config,
	handler tcp.Handler) error {
	closeChan := make(chan struct{})
	listener, err := net.Listen("tcp", cfg.Address)
	if err != nil {
		return err
	}
	logger.Info("start listen ")
	ListenAndServe(listener, handler, closeChan)
	return nil
}
func ListenAndServe(listener net.Listener, handler tcp.Handler,
	closeChan <-chan struct{}) {
	ctx := context.Background()
	for true {
		conn, err := listener.Accept()
		if err != nil {
			break
		}
		logger.Info("accepted link")
		go func() {
			handler.Handler(ctx, conn)
		}()
	}
}
