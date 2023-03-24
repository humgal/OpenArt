package db

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/mysql"
	"github.com/humgal/art-server/util"
)

var DB *sql.DB

func InitDB() {
	db, err := sql.Open("mysql", util.Config.Mysqlhost)

	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	DB = db
}

func CloseDB() error {
	return DB.Close()
}

func Migrate() {
	if err := DB.Ping(); err != nil {
		util.Logger.Fatal(err)
	}
	driver, _ := mysql.WithInstance(DB, &mysql.Config{})
	m, _ := migrate.NewWithDatabaseInstance(
		"file://db/migrations/mysql",
		"mysql",
		driver,
	)
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		util.Logger.Fatal(err)
	}

}
