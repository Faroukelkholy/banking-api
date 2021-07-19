package main

import (
	"context"
	"fmt"
	"time"

	"github.com/faroukelkholy/bank/config"
	"github.com/faroukelkholy/bank/internal/server"
	"github.com/faroukelkholy/bank/internal/service/account"
	"github.com/faroukelkholy/bank/internal/service/customer"
	"github.com/faroukelkholy/bank/internal/service/transaction"
	"github.com/faroukelkholy/bank/internal/storage/postgres"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cfg := config.Parse()

	repo, err := postgres.New(ctx, &postgres.Options{
		Debug:    cfg.Debug,
		DBHost:   cfg.DBHost,
		DBPort:   cfg.DBPort,
		DBName:   cfg.DBName,
		DBUser:   cfg.DBUser,
		DBPass:   cfg.DBPass,
		DBSchema: cfg.DBSchema,
	})
	if err != nil {
		fmt.Println(err)
	}

	srv := server.New()
	srv.AddRoutesAS(account.New(repo))
	srv.AddRoutesCS(customer.New(repo))
	srv.AddRoutesTS(transaction.New(repo))

	if err = srv.Start(fmt.Sprintf(":%s", cfg.HTTPPort)); err != nil {
		fmt.Println(err)
	}
}
