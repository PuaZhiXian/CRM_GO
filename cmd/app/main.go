package main

import (
	"context"
	"crm-backend/cmd/db"
	api "crm-backend/gen"
	"crm-backend/internal/httpx"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
)

func main() {
	//connect database
	ctx := context.Background()
	db, err := db.Open(ctx, []string{})
	if err != nil {
		log.Panic(err)
	}
	mysqlDb, _ := db.DB()
	defer mysqlDb.Close()

	//handler setup
	handler := httpx.InitHandle(db)
	r := chi.NewMux()
	h := api.HandlerFromMux(handler, r)

	//server setup
	port := "8080"
	readTimeOut := time.Second * 30
	writeTimeOut := time.Second * 30
	server := httpx.NewServer(port, readTimeOut, writeTimeOut)
	server.Start(h)
	defer shutDownServer(server)

	log.Println("server running on port ", port)
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, os.Kill, syscall.SIGINT, syscall.SIGTERM)
	_ = <-sig
}

func shutDownServer(server *httpx.Server) {
	if err := server.ShutDown(context.Background()); err != nil {
		panic(err)
	}
}
