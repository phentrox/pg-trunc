package main

import (
	"github.com/phentrox/db-pg-trunc/internal/config"
	"github.com/phentrox/db-pg-trunc/internal/postgresql"
	"github.com/phentrox/db-pg-trunc/internal/truncinit"
	"log"
)

func runInit() {
	configEntity := config.ReadConfig("pgtrunc.yaml")

	postgresql.OpenPostgreSqlConnection(configEntity)

	println("Running init ...")
	err := truncinit.TruncationInit()
	if err != nil {
		log.Fatal(err)
	}
	println("Init Succeeded!")
}
