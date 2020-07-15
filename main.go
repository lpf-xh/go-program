package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	srv := &http.Server{
		Addr:    ":8081",
		Handler: http.HandlerFunc(getHello),
	}

	go func() {
		err := srv.ListenAndServe()
		if err != nil {
			log.Println("server err:", err)
		}
	}()

	shutdown(context.Background(), srv)
}

func getHello(w http.ResponseWriter, r *http.Request) {
	time.Sleep(5 * time.Second)
	fmt.Fprintln(w, "hello")
}

func shutdown(ctx context.Context, srv *http.Server) {
	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	<-c

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("server shutdown failed, err:", err)
	}
	log.Println("server gracefully shutdown")
}

