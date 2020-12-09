package main

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main()  {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(100)
	})

	http.HandleFunc("/a", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, 123456789)
	})

	http.HandleFunc("/b", func(w http.ResponseWriter, r *http.Request) {
		log.Fatal("错误啦")
	})

	s := http.Server{Addr: ":8080"}

	eg := errgroup.Group{}
	serErr := make(chan error, 1)
	sinC := make(chan os.Signal, 1)

	eg.Go(func() error {
		go func() {
			serErr <- s.ListenAndServe()
		}()
		select {
		case err := <-serErr:
			close(sinC)
			close(serErr)
			return err
		}
	})

	eg.Go(func() error {
		signal.Notify(sinC, syscall.SIGINT|syscall.SIGTERM|syscall.SIGKILL)
		<-sinC
		return s.Shutdown(context.TODO())
	})

	log.Println(eg.Wait())
}