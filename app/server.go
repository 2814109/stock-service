package main

import (
	// "github.com/uptrace/bunrouter"

	// "context"
	// "fmt"
	"app/db"
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

	db := db.OpenDB()
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))
	// handler := cors.Default().Handler(srv)
	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	qiita_user_name := os.Getenv("QIITA_USER_NAME")
	log.Printf("user name is %s", qiita_user_name)
	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))

}
