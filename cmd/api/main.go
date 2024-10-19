package main

import (
	"context"
	"log"

	"github.com/marifsulaksono/go-echo-boilerplate/internal/api"
	"github.com/marifsulaksono/go-echo-boilerplate/internal/config"
	"github.com/marifsulaksono/go-echo-boilerplate/internal/contract"
)

func main() {
	var ctx = context.Background()
	err := config.Load(ctx, false)
	if err != nil {
		log.Fatalf("Error load configuration: %v", err)
	}

	contact, err := contract.NewContract(ctx)
	if err != nil {
		log.Fatalf("Error setup contract / dependecy injection: %v", err)
	}

	e := api.NewHTTPServer(contact)
	e.RunHTTPServer()
}
