package server

import (
	"github.com/ezavalishin/versus3/internal/handlers"
	"github.com/ezavalishin/versus3/pkg/utils"
	"github.com/gin-gonic/gin"
	"log"
)

var host, port string

func init() {
	host = utils.MustGet("SERVER_HOST")
	port = utils.MustGet("SERVER_PORT")
}

func Run() {
	r := gin.Default()

	r.GET("/ping", handlers.Ping())

	log.Println("Running @ http://" + host + ":" + port)
	log.Fatalln(r.Run(host + ":" + port))
}
