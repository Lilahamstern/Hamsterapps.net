package handler

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
	"github.com/lilahamstern/hamsterapps.net/server/graph"
	"github.com/lilahamstern/hamsterapps.net/server/graph/generated"
)

func GraphqlHandler() gin.HandlerFunc {
	h := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func PlaygroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL Hamsterapps", "/query")

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}
