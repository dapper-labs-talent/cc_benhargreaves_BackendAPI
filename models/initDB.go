package models

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

var (
	dbhost     = os.Getenv("DB_HOST")
	dbport     = os.Getenv("DB_PORT")
	dbuser     = os.Getenv("DB_USER")
	dbpassword = os.Getenv("DB_PASSWORD")
	dbname     = os.Getenv("DB_DATABASE")
)

var DB *bun.DB

func InitDB() error {

	dbaddress := fmt.Sprintf("%s:%s", dbhost, dbport)

	pgconn := pgdriver.NewConnector(
		pgdriver.WithNetwork("tcp"),
		pgdriver.WithAddr(dbaddress),
		pgdriver.WithInsecure(true),
		pgdriver.WithUser(dbuser),
		pgdriver.WithPassword(dbpassword),
		pgdriver.WithDatabase(dbname),
	)

	sqldb := sql.OpenDB(pgconn)

	err := sqldb.Ping()
	if err != nil {
		return err
	}

	bunDB := bun.NewDB(sqldb, pgdialect.New())

	DB = bunDB
	fmt.Println("Successfully connected!")

	return nil
}
