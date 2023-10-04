package main

import (
	// "app/graph"
	// "github.com/uptrace/bunrouter"
	"context"
	"fmt"
	"log"
	// "net/http"
	// "os"
	"time"

	// "github.com/99designs/gqlgen/graphql/handler"
	// "github.com/99designs/gqlgen/graphql/playground"

	"database/sql"
	_ "github.com/lib/pq"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

type Post struct {
	ID        int64 `bun:",pk,autoincrement"`
	Content   string
	CreatedAt time.Time `bun:",nullzero,notnull,default:current_timestamp"`
	UpdatedAt time.Time `bun:",nullzero,notnull,default:current_timestamp"`
}

type Sample struct {
	ID        int64 `bun:",pk,autoincrement"`
	Content   string
	CreatedAt time.Time `bun:",nullzero,notnull,default:current_timestamp"`
	UpdatedAt time.Time `bun:",nullzero,notnull,default:current_timestamp"`
}

func main() {
	dsn := "postgres://postgres:postgres@stock-postgres:5433/postgres?sslmode=disable"
	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))

	db := bun.NewDB(sqldb, pgdialect.New())

	var v string
	if err := db.NewSelect().ColumnExpr("version()").Scan(context.Background(), &v); err != nil {
		log.Fatal(err)
	}
	fmt.Println(v)
	log.Printf("success %v", db)

	// if _, err := db.NewCreateTable().Model((*Sample)(nil)).Exec(context.Background()); err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("create table of %v", "Sample")

	res, err := db.NewDropTable().Model((*Post)(nil)).Exec(context.Background())

	if err != nil {
		log.Fatal(err)
	}
	log.Printf("success %v", res)

	if _, err := db.NewCreateTable().Model((*Post)(nil)).Exec(context.Background()); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("create table of %v", "Post")

	post := &Post{Content: "post content"}
	insertRes, err := db.NewInsert().Model(post).Exec(context.Background())
	if err != nil {
		log.Fatal(err)

	}
	log.Printf("success %v", insertRes)

}
