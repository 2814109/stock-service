package db

import (
	"database/sql"

	_ "github.com/lib/pq"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

func OpenDB() *bun.DB {
	dsn := "postgres://postgres:postgres@stock-postgres:5433/postgres?sslmode=disable"
	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))

	db := bun.NewDB(sqldb, pgdialect.New())
	// if err := db.Ping(); err != nil {
	// 	log.Fatal(err)
	// }
	// defer db.Close()
	return db
}
