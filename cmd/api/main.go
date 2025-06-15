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

// more info contact me at @marifsulaksono
func main() {
	var ctx = context.Background()
	err := config.Load(ctx, true)
	if err != nil {
		log.Fatalf("Error load configuration with value isEnvFile = true: %v", err)
	}

	contract, err := contract.NewContract(ctx)
	if err != nil {
		log.Fatalf("Error setup contract / dependecy injection: %v", err)
	}

	contract.Common.AutoMigrate()

	e := api.NewHTTPServer(contract)
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGTERM, syscall.SIGINT)
	go func() {
		e.RunHTTPServer()
	}()

	<-sig

	log.Println("Shutting down....")
	// put all processes to be stopped before successful termination here
	contract.Common.Close()

	log.Printf("Server %s with UID: %s is gracefully terminated.", config.Config.App.Name, config.Config.App.UID)
}
