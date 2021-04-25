package main

import (
	"fmt"

	"github.com/faroukelkholy/bank/config"
	"github.com/faroukelkholy/bank/internal/server"
	"github.com/faroukelkholy/bank/internal/service/account"
	"github.com/faroukelkholy/bank/internal/service/customer"
	"github.com/faroukelkholy/bank/internal/service/transaction"
	"github.com/faroukelkholy/bank/internal/storage/postgres"
)

func main() {
	repo, err := postgres.New(&postgres.Options{
		Debug:    config.Parse().Debug,
		DBHost:   config.Parse().DBHost,
		DBPort:   config.Parse().DBPort,
		DBName:   config.Parse().DBName,
		DBUser:   config.Parse().DBUser,
		DBPass:   config.Parse().DBPass,
		DBSchema: config.Parse().DBSchema,
	})
	if err != nil {
		panic(err)
	}

	srv := server.New()
	srv.AddRoutesAS(account.New(repo))
	srv.AddRoutesCS(customer.New(repo))
	srv.AddRoutesTS(transaction.New(repo))

	if err = srv.Start(fmt.Sprintf(":%s", config.Parse().HTTPPort)); err != nil {
		panic(err)
	}
}
