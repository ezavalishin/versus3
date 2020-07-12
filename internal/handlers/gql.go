package handlers

import (
	"fmt"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/ezavalishin/versus3/internal/gql"
	"github.com/ezavalishin/versus3/internal/gql/resolvers"
	"github.com/ezavalishin/versus3/internal/orm"
	"github.com/gin-gonic/gin"
)

func GraphqlHandler(orm *orm.ORM) gin.HandlerFunc {
	c := gql.Config{
		Resolvers: &resolvers.Resolver{
			ORM: orm,
		},
	}

	h := handler.New(gql.NewExecutableSchema(c))
	h.AddTransport(transport.POST{})
	h.AddTransport(transport.GET{})
	h.AddTransport(transport.Options{})
	h.AddTransport(transport.MultipartForm{})

	h.SetQueryCache(lru.New(1000))

	h.Use(extension.Introspection{})
	h.Use(extension.AutomaticPersistedQuery{
		Cache: lru.New(100),
	})

	return func(context *gin.Context) {
		h.ServeHTTP(context.Writer, context.Request)
	}
}

func PlaygroundHandler(path string) gin.HandlerFunc {
	h := playground.Handler("Go GraphQL Server", path)

	return func(context *gin.Context) {
		fmt.Printf("%+v\n", context.Request.Header.Get("VkParams"))
		h.ServeHTTP(context.Writer, context.Request)
	}
}
