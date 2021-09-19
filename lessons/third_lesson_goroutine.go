package main

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

var signalChan = make(chan os.Signal)

func init() {
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
}

func main() {
	rootCtx, cancel := context.WithCancel(context.Background())

	eg, ctx := errgroup.WithContext(rootCtx)

	eg.Go(func() error {
		<-ctx.Done()
		fmt.Println("done!")
		//这里不确定如何关闭，所以采用的exit方法
		os.Exit(1)
		return nil
	})

	http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("test success"))
	})

	eg.Go(func() error {
		return http.ListenAndServe("127.0.0.1:8080", nil)
	})

	eg.Go(func() error {

		return http.ListenAndServe("127.0.0.1:8081", nil)
	})

	eg.Go(func() error {
		for {
			select {
			case <-ctx.Done():
				return ctx.Err()
			case s := <-signalChan:
				switch s {
				case syscall.SIGINT, syscall.SIGTERM:
					fmt.Println("received signal")
					cancel()
				default:
					fmt.Println("undefined signal")
				}
			}
		}
	})
	if err := eg.Wait(); err != nil && err != context.Canceled {
		fmt.Printf("err: %v", err)
	}
}
