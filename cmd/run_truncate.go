package main

import (
	"github.com/phentrox/db-pg-trunc/internal/config"
	"github.com/phentrox/db-pg-trunc/internal/postgresql"
	"github.com/phentrox/db-pg-trunc/internal/truncate"
)

func runTruncate() {
	configEntity := config.ReadConfig("pgtrunc.yaml")

	postgresql.OpenPostgreSqlConnection(configEntity)

	println("Truncating ...")
	truncate.Truncate(configEntity.Schemas)
	println("Truncation Successfull!")

	postgresql.ClosePostgreSqlConnection()
}
