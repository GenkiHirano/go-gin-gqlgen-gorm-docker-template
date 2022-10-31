package main

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/GenkiHirano/go-gin-gqlgen-gorm-docker-template/graph"
	"github.com/GenkiHirano/go-gin-gqlgen-gorm-docker-template/graph/generated"
	"github.com/GenkiHirano/go-gin-gqlgen-gorm-docker-template/infra"
	"github.com/gin-gonic/gin"
)

func graphqlHandler() gin.HandlerFunc {
	h := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func playgroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL", "/query")

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func main() {
	r := gin.Default()
	r.POST("/query", graphqlHandler())
	r.GET("/", playgroundHandler())

	db := infra.OpenDB()
	infra.SeedTodo(db)
	defer infra.CloseDB(db)

	r.Run()
}
