package infra

import (
	"database/sql"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

func RunMigrateScripts() (*sql.DB, error) {

	db, err := sql.Open("postgres", "postgres://myusername:mypassword@localhost:5432/mydb?sslmode=disable")
	if err != nil {
		return nil, err
	}

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return db, err
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://infra/db/",
		"postgres", driver)
	if err != nil {
		return db, err
	}

	m.Up()

	return db, nil
}
