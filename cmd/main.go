package main

import (
	"fmt"

	"github.com/faroukelkholy/bank/config"
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
	fmt.Println(repo)
}
