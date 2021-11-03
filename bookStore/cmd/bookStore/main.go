package main

import (
	_ "bookStore/internal/store"
	"bookStore/store/factory"
	"bookstore/server"
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	s, err := factory.New("mem")
	if err != nil {
		panic(err)
	}

	service := server.NewBookStoreServer(":8080", s)

	errChan, err = service.ListenAndServe()
	if err != nil {
		log.Println("Web server start failed: ", err)
		return
	}
	log.Println("Web server starts successfully!")

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)

	select {
	case err = <-errChan:
		log.Println("Web server run failed: ", err)
		return
	case <-c:
		log.Println("Bookstore program exists")
		ctx, cf := context.WithTimeout(context.Background(), time.Second)
		defer cf()
		err = service.Shutdown(ctx)
	}

	if err != nil {
		log.Println("Bookstore program exits error: ", err)
		return
	}
	log.Println("Bookstore program exits successfully!")
}
