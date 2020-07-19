package server

import (
	"context"
	"github.com/ezavalishin/versus3/internal/auth"
	"github.com/ezavalishin/versus3/internal/handlers"
	log "github.com/ezavalishin/versus3/internal/logger"
	"github.com/ezavalishin/versus3/internal/orm"
	"github.com/ezavalishin/versus3/pkg/utils"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
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

func AuthMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {
		ctx := context.WithValue(c.Request.Context(), "GinContextKey", c)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}

func Run(orm *orm.ORM) {
	log.Info("GORM_CONNECTION_DSN: ", utils.MustGet("GORM_CONNECTION_DSN"))

	endpoint := "http://" + host + ":" + port

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowMethods:     []string{"GET", "POST"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Vk-Params", "Origin", "Host", "Connection", "Accept-Encoding", "Accept-Language", "X-Requested-With"},
		AllowCredentials: true,
		AllowOrigins:     []string{"*"},
		MaxAge:           12 * time.Hour,
	}))

	r.GET("/ping", handlers.Ping())

	if isPgEnabled {
		r.GET(gqlPgPath, handlers.PlaygroundHandler(gqlPath))
		log.Info("GraphQL Playground @ " + endpoint + gqlPgPath)
	}

	authorized := r.Group("/")

	authorized.Use(auth.Middleware(orm))
	{
		authorized.POST(gqlPath, handlers.GraphqlHandler(orm))
	}

	log.Info("GraphQL @ " + endpoint + gqlPath)

	log.Info("Running @ http://" + host + ":" + port)
	log.Info(r.Run(host + ":" + port))
}
