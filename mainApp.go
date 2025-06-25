package main

import (
	"context"
	"crm-backend/db"
	"crm-backend/db/httpx"
	"crm-backend/handlers"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	ctx := context.Background()
	db, err := db.Open(ctx, []string{})
	if err != nil {
		log.Panic(err)
	}
	defer db.Close()

	handler := handlers.InitHandle(db)

	port := "8080"
	readTimeOut := time.Second * 30
	writeTimeOut := time.Second * 30
	server := httpx.NewServer(port, readTimeOut, writeTimeOut)
	server.Start([]httpx.Router{handler})

	defer func() {
		if err := server.ShutDown(context.Background()); err != nil {
			panic(err)
		}
	}()

	log.Println("server running on port ", port)
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, os.Kill, syscall.SIGINT, syscall.SIGTERM)
	_ = <-sig
}
