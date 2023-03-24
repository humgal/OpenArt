package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
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
	res, err := db.DB.Exec("select * from item")
	if err == nil {
		println(res.RowsAffected())
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/graphql"))
	http.Handle("/graphql", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))

}
