package handlers

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/ezavalishin/versus3/internal/gql"
	"github.com/ezavalishin/versus3/internal/gql/resolvers"
	"github.com/gin-gonic/gin"
)

func GraphqlHandler() gin.HandlerFunc {
	c := gql.Config{
		Resolvers: &resolvers.Resolver{},
	}

	h := handler.New(gql.NewExecutableSchema(c))
	h.AddTransport(transport.POST{})
	return func(context *gin.Context) {
		h.ServeHTTP(context.Writer, context.Request)
	}
}

func PlaygroundHandler(path string) gin.HandlerFunc {
	h := playground.Handler("Go GraphQL Server", path)

	return func(context *gin.Context) {
		h.ServeHTTP(context.Writer, context.Request)
	}
}
