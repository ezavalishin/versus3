package server

import (
	"github.com/ezavalishin/versus3/internal/handlers"
	"github.com/ezavalishin/versus3/pkg/utils"
	"github.com/gin-gonic/gin"
	"log"
)

var host, port, gqlPath, gqlPgPath string
var isPgEnabled bool

func init() {
	host = utils.MustGet("SERVER_HOST")
	port = utils.MustGet("SERVER_PORT")

	gqlPath = utils.MustGet("SERVER_GRAPHQL_PATH")
	gqlPgPath = utils.MustGet("SERVER_GRAPHQL_PLAYGROUND_PATH")
	isPgEnabled = utils.MustGetBool("SERVER_GRAPHQL_PLAYGROUND_IS_ENABLED")
}

func Run() {
	endpoint := "http://" + host + ":" + port

	r := gin.Default()

	r.GET("/ping", handlers.Ping())

	if isPgEnabled {
		r.GET(gqlPgPath, handlers.PlaygroundHandler(gqlPath))
		log.Println("GraphQL Playground @ " + endpoint + gqlPgPath)
	}

	r.POST(gqlPath, handlers.GraphqlHandler())
	log.Println("GraphQL @ " + endpoint + gqlPath)

	log.Println("Running @ http://" + host + ":" + port)
	log.Fatalln(r.Run(host + ":" + port))
}
