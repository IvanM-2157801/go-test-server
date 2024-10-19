package main

import (
	"log"
	"net/http"
	"os"
	"test-server/database"
	"test-server/graph"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/joho/godotenv"
)

const defaultPort = "8080"

func goDotEnvVariable(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	return os.Getenv(key)
}

func main() {
	// setup DB
	database.ConnectDB(goDotEnvVariable("DB_USER"), goDotEnvVariable("DB_PASS"))
	database.CreateDB()
	database.MigrateDB()

	// setup graphql server
	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{
		Database: database.DBInstance,
	}}))

	// handle requests
	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", defaultPort)
	log.Fatal(http.ListenAndServe(":"+defaultPort, nil))
}
