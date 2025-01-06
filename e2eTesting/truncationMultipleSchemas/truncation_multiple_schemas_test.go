package commandsUpVersionZeroTest

import (
	"github.com/phentrox/db-pg-trunc/e2eTesting"
	"github.com/phentrox/db-pg-trunc/e2eTesting/dbTestingHelpers"
	"github.com/phentrox/db-pg-trunc/internal/config"
	"github.com/phentrox/db-pg-trunc/internal/postgresql"
	"github.com/phentrox/db-pg-trunc/internal/truncate"
	"github.com/phentrox/db-pg-trunc/internal/truncinit"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTruncationMultipleSchemas(t *testing.T) {
	// setup: connection
	postgresConfig := config.ReadConfig(e2eTesting.ConfigPathTwoDirsUp)
	postgresql.OpenPostgreSqlConnection(postgresConfig)

	// setup: create schemas
	_, itSchemaErr := postgresql.DB.Exec("CREATE SCHEMA it")
	assert.NoError(t, itSchemaErr)
	_, hrSchemaErr := postgresql.DB.Exec("CREATE SCHEMA hr")
	assert.NoError(t, hrSchemaErr)
	_, salesSchemaErr := postgresql.DB.Exec("CREATE SCHEMA sales")
	assert.NoError(t, salesSchemaErr)

	// setup: create shouldBeEmptyTables
	createItTables(t)
	createHrTables(t)
	createSalesTables(t)

	// setup: insert data
	insertItData(t)
	insertHrData(t)
	insertSalesData(t)

	// setup: init truncation (function)
	err := truncinit.TruncationInit()
	assert.NoError(t, err)

	// given
	schemas := []string{"it", "hr"}

	// when
	truncate.Truncate(schemas)

	// then all the tables inside the schemas from the config should be empty
	shouldBeEmptyTables := []string{"it.software", "it.hardware", "hr.department", "hr.employee"}
	for _, table := range shouldBeEmptyTables {
		tableIsEmpty, tableCheckErr := dbTestingHelpers.TableIsEmpty(table)
		assert.NoError(t, tableCheckErr)
		assert.True(t, tableIsEmpty)
	}

	shouldNotBeEmptyTables := []string{"sales.account", "sales.sale"}
	for _, table := range shouldNotBeEmptyTables {
		tableIsEmpty, tableCheckErr := dbTestingHelpers.TableIsEmpty(table)
		assert.NoError(t, tableCheckErr)
		assert.False(t, tableIsEmpty)
	}

	// cleanup
	cleanup(t)
}

func insertItData(t *testing.T) {
	insertSoftware := `INSERT INTO it.software
    VALUES
    (DEFAULT, 'erp'),
    (DEFAULT, 'word'),
    (DEFAULT, 'excel')
    `
	_, err := postgresql.DB.Exec(insertSoftware)
	assert.NoError(t, err)

	insertHardware := `INSERT INTO it.hardware
    VALUES
    (DEFAULT, 'computer'),
    (DEFAULT, 'monitor'),
    (DEFAULT, 'mouse')
    `
	_, err = postgresql.DB.Exec(insertHardware)
	assert.NoError(t, err)
}

func insertHrData(t *testing.T) {
	insertDepartments := `INSERT INTO hr.department
    VALUES
    (1, 'it'),
    (2, 'hr'),
    (3, 'production')
    `
	_, err := postgresql.DB.Exec(insertDepartments)
	assert.NoError(t, err)

	insertEmployees := `INSERT INTO hr.employee
    VALUES
    (DEFAULT, 'horst', 'maier', 1),
    (DEFAULT, 'maria', 'becker', 2),
    (DEFAULT, 'franz', 'bauer', 3)
    `
	_, err = postgresql.DB.Exec(insertEmployees)
	assert.NoError(t, err)
}

func insertSalesData(t *testing.T) {
	insertAccounts := `INSERT INTO sales.account
    VALUES
    (1, 'ebay'),
    (2, 'amazon'),
    (3, 'microsoft')
    `
	_, err := postgresql.DB.Exec(insertAccounts)
	assert.NoError(t, err)

	insertEmployees := `INSERT INTO sales.sale
    VALUES
    (DEFAULT, 1, 2000),
    (DEFAULT, 2, 3000),
    (DEFAULT, 2, 4000),
    (DEFAULT, 3, 1000)
    `
	_, err = postgresql.DB.Exec(insertEmployees)
	assert.NoError(t, err)
}

func createItTables(t *testing.T) {
	itSoftwareTableQuery := `CREATE TABLE it.software (
        id BIGSERIAL UNIQUE PRIMARY KEY,
        name TEXT NOT NULL
    )
    `
	_, err := postgresql.DB.Query(itSoftwareTableQuery)
	assert.NoError(t, err)

	itHardwareTableQuery := `CREATE TABLE it.hardware (
        id BIGSERIAL UNIQUE PRIMARY KEY,
        name TEXT NOT NULL
    )
    `
	_, err = postgresql.DB.Query(itHardwareTableQuery)
	assert.NoError(t, err)
}

func createHrTables(t *testing.T) {
	departmentTable := `CREATE TABLE hr.department (
        id BIGINT UNIQUE PRIMARY KEY,
        name TEXT NOT NULL
    )
    `
	_, err := postgresql.DB.Query(departmentTable)
	assert.NoError(t, err)

	employeeTable := `CREATE TABLE hr.employee (
        id BIGSERIAL UNIQUE PRIMARY KEY,
        first_name TEXT NOT NULL,
        last_name TEXT NOT NULL,
		department_id BIGINT NOT NULL references hr.department(id)
    )
    `
	_, err = postgresql.DB.Query(employeeTable)
	assert.NoError(t, err)

}

func createSalesTables(t *testing.T) {
	accountTable := `CREATE TABLE sales.account (
        id BIGINT UNIQUE PRIMARY KEY,
        name TEXT NOT NULL
    )
    `
	_, err := postgresql.DB.Query(accountTable)
	assert.NoError(t, err)

	saleTable := `CREATE TABLE sales.sale (
        id BIGSERIAL UNIQUE PRIMARY KEY,
        account_id BIGINT NOT NULL references sales.account(id),
        amount INT NOT NULL
    )
    `
	_, err = postgresql.DB.Query(saleTable)
	assert.NoError(t, err)
}

func cleanup(t *testing.T) {
	// cleanup schemas
	_, dropItSchemaErr := postgresql.DB.Exec("DROP SCHEMA it CASCADE ")
	assert.NoError(t, dropItSchemaErr)
	_, dropHrSchemaErr := postgresql.DB.Exec("DROP SCHEMA hr CASCADE ")
	assert.NoError(t, dropHrSchemaErr)
	_, dropSalesSchemaErr := postgresql.DB.Exec("DROP SCHEMA sales CASCADE ")
	assert.NoError(t, dropSalesSchemaErr)

	// cleanup function
	_, err := postgresql.DB.Query("DROP FUNCTION IF EXISTS truncate_tables_in_schemas")
	assert.NoError(t, err)
}
