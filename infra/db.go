package infra

import (
	"database/sql"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/golang-migrate/migrate/v4"
	pg "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

func RunMigrateScripts(conf Config) (*sql.DB, error) {

	db, err := sql.Open("postgres", conf.DbConnString)
	if err != nil {
		return nil, err
	}

	driver, err := pg.WithInstance(db, &pg.Config{})
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

func InitDB(conf Config) *gorm.DB {

	db, err := gorm.Open(postgres.Open(conf.DbConnString), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	return db
}
