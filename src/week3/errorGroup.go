package main

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"net/http"
	"os"
	"os/signal"
)

func myHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("myHandler is work")
}

func Httpserver(srv *http.Server) error {
	http.HandleFunc("/myHandler", myHandler)
	fmt.Println("http server start")
	err := srv.ListenAndServe()
	return err
}

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	group, errCtx := errgroup.WithContext(ctx)
	srv := &http.Server{Addr: ":9090"}
	//启动Httpserver
	group.Go(func() error {
		return Httpserver(srv)
	})

	group.Go(func() error {
		<-errCtx.Done() //阻塞。
		fmt.Println("http server stop")
		return srv.Shutdown(errCtx) // 关闭 http server
	})

	chanel := make(chan os.Signal, 1) //监听型号量
	signal.Notify(chanel)

	group.Go(func() error {
		for {
			select {
			case <-errCtx.Done():
				fmt.Println("go errCtx " + errCtx.Err().Error())
				return errCtx.Err()
			case <-chanel:
				fmt.Println("get os.Signa!")
				cancel()
			}
		}
		return nil
	})

	if err := group.Wait(); err != nil {
		fmt.Println("group error: ", err)
	}
	fmt.Println("all group done!")
}
