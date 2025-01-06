package dbTestingHelpers

import "github.com/phentrox/db-pg-trunc/internal/postgresql"

func TableIsEmpty(table string) (bool, error) {
	query := "SELECT EXISTS (SELECT * FROM " + table + ") AS table_has_rows;"

	var tableHasRows bool
	row := postgresql.DB.QueryRow(query)
	err := row.Scan(&tableHasRows)
	if err != nil {
		return false, err
	}

	return !tableHasRows, err

}
