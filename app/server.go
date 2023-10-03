package main

import (
	"app/graph"
	// "github.com/uptrace/bunrouter"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"

	"database/sql"
	_ "github.com/lib/pq"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

const defaultPort = "8081"

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

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))
	// router := bunrouter.New()

	// router.GET("/", func(w http.ResponseWriter, req bunrouter.Request) error {
	// 	// req embeds *http.Request and has all the same fields and methods
	// 	fmt.Println(req.Method, req.Route(), req.Params().Map())
	// 	return w.playground.Handler("GraphQL playground", "/query")
	// })
	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
