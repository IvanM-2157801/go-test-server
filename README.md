# GO swapi server

## Go server spoedcursus
```go
// om een package toe te voegen => in tools.go, vervolgens:$
go mod tidy

// om de boilerplate code te genereren
go run github.com/99designs/gqlgen init

// resolver.go wordt gebruikt voor lokale of DB storage
// resolvers in schema.resolvers.go
// voor mapping van graphql naar go structs => models/<modelname>.go

// na nieuwe models toegevoegd te hebben in schema.graphqls
go run github.com/99designs/gqlgen generate

// om de server te starten
go run server.go
```

[Handige bron](https://gqlgen.com/getting-started/)
[GO cheatsheet](https://devhints.io/go)
