package truncate

import "github.com/phentrox/db-pg-trunc/internal/postgresql"

func TruncateAllSchemasInList(schemaSqlArray string) error {
	_, err := postgresql.DB.Exec("SELECT  truncate_tables_in_schemas($1)", schemaSqlArray)
	if err != nil {
		return err
	}

	return nil
}
