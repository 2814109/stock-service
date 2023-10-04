package main

import (
	// "github.com/uptrace/bunrouter"

	// "context"
	// "fmt"
	"app/graph"
	"log"
	"net/http"
	"os"

	// "time"

	// "github.com/rs/cors"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	// "database/sql"
	// _ "github.com/lib/pq"
	// "github.com/uptrace/bun"
	// "github.com/uptrace/bun/dialect/pgdialect"
	// "github.com/uptrace/bun/driver/pgdriver"
)

const defaultPort = "8081"

func main() {

	// if _, err := db.NewCreateTable().Model((*Post)(nil)).Exec(context.Background()); err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println("create table")

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))
	// handler := cors.Default().Handler(srv)
	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
