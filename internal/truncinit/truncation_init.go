package truncinit

import "github.com/phentrox/db-pg-trunc/internal/postgresql"

func TruncationInit() error {
	query := `
        CREATE OR REPLACE FUNCTION truncate_tables_in_schemas(schemas text[]) RETURNS void AS
		$$
		DECLARE
			schema_name text;
			tbl_name text;
		BEGIN
			-- Loop over the specified schemas
			FOR schema_name IN SELECT unnest(schemas)
				LOOP
					-- Loop over all tables in the current schema
					FOR tbl_name IN
						SELECT t.table_name
						FROM information_schema.tables t
						WHERE t.table_schema = schema_name
						  AND t.table_type = 'BASE TABLE'
						LOOP
							-- Generate and execute the TRUNCATE statement with correct format specifiers
							EXECUTE format('TRUNCATE TABLE %I.%I CASCADE;', schema_name, tbl_name);
						END LOOP;
				END LOOP;
		END
		$$ LANGUAGE plpgsql;
    `
	_, err := postgresql.DB.Exec(query)
	if err != nil {
		return err
	}

	return nil
}
