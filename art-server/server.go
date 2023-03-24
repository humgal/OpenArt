package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
	"github.com/humgal/art-server/auth"
	"github.com/humgal/art-server/db"
	"github.com/humgal/art-server/graph"
	"github.com/humgal/art-server/util/redis"
)

const defaultPort = "8000"

func main() {

	ctx := context.Background()
	if redis.Rdb.Set(ctx, "key", "value", 0).Err() != nil {
		panic("set redis error")
	}
	db.InitDB()
	defer db.CloseDB()
	db.Migrate()

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	router := chi.NewRouter()
	router.Use(auth.Middleware())
	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))

	router.Handle("/", playground.Handler("GraphQL playground", "/graphql"))
	router.Handle("/graphql", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))

}
