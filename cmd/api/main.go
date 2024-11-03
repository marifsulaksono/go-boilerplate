package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/marifsulaksono/go-echo-boilerplate/internal/api"
	"github.com/marifsulaksono/go-echo-boilerplate/internal/config"
	"github.com/marifsulaksono/go-echo-boilerplate/internal/contract"
)

func main() {
	var ctx = context.Background()
	err := config.Load(ctx, true)
	if err != nil {
		log.Fatalf("Error load configuration with value isEnvFile = true: %v", err)
	}

	contact, err := contract.NewContract(ctx)
	if err != nil {
		log.Fatalf("Error setup contract / dependecy injection: %v", err)
	}
	contact.Common.AutoMigrate()

	e := api.NewHTTPServer(contact)
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGTERM, syscall.SIGINT)
	go func() {
		e.RunHTTPServer()
	}()

	<-sig

	log.Println("Shutting down....")
	// put all processes to be stopped before successful termination here
	log.Println("Server gracefully terminated.")
}
