package main

import (
	"github.com/phentrox/db-pg-trunc/internal/config"
	"github.com/phentrox/db-pg-trunc/internal/postgresql"
	"github.com/phentrox/db-pg-trunc/internal/truncinit"
	"log"
)

func runInit() {
	println("Reading Config ...")
	configEntity := config.ReadConfig("pgtrunc.yaml")

	println("Connecting to Database ...")
	postgresql.OpenPostgreSqlConnection(configEntity)

	println("Running init ...")
	err := truncinit.TruncationInit()
	if err != nil {
		log.Fatal(err)
	}
	println("Init Succeeded!")
}
