package main

import (
	"context"
	"fmt"
	"github.com/oklog/run"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {

	// 编排开始
	var g run.Group
	ctxAll, cancelAll := context.WithCancel(context.Background())
	{

		// 处理信号退出的handler
		term := make(chan os.Signal, 1)
		signal.Notify(term, os.Interrupt, syscall.SIGTERM)
		cancelC := make(chan struct{})
		g.Add(
			func() error {
				select {
				case <-term:
					fmt.Println("msg", "Receive SIGTERM ,exiting gracefully....")
					cancelAll()
					return nil
				case <-cancelC:
					fmt.Println("msg", "other cancel exiting")
					return nil
				}
			},
			func(err error) {
				close(cancelC)
			},
		)
	}
	{
		// logjob 结果的metrics http server
		g.Add(func() error {
			errChan := make(chan error, 1)
			go func() {
				for {
					time.Sleep(time.Second * 1)
					fmt.Println("Go routine 1 is sleeping...")
				}
				fmt.Println("Go routine 3 is closed")
			}()
			select {
			case err := <-errChan:
				fmt.Println("msg", "logjob.metrics.web.server.error", "err", err)
				return err
			case <-ctxAll.Done():
				fmt.Println("msg", "receive_quit_signal_web_server_exit")
				return nil
			}

		}, func(err error) {
			cancelAll()
		},
		)
	}
	g.Run()
}
