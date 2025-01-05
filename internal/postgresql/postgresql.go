package postgresql

import (
	"database/sql"
	"fmt"
	"github.com/phentrox/db-pg-trunc/internal/config"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func OpenPostgreSqlConnection(config config.Config) {
	connString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.Host, config.Port, config.User, config.Password, config.Database)
	var err error
	DB, err = sql.Open("postgres", connString)
	if err != nil {
		panic("postgres opening connection panic: " + err.Error())
	}
}

func ClosePostgreSqlConnection() {
	err := DB.Close()
	if err != nil {
		panic(err)
	}
}
