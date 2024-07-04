package infra

import (
	"database/sql"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

func RunMigrateScripts() error {

	db, err := sql.Open("postgres", "postgres://myusername:mypassword@localhost:5432/mydb?sslmode=disable")
	if err != nil {
		return err
	}

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return err
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://infra/db/",
		"postgres", driver)
	if err != nil {
		return err
	}

	m.Up()
	return nil
}
