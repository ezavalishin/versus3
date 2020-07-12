package main

import (
	orm2 "github.com/ezavalishin/versus3/internal/orm"
	"github.com/ezavalishin/versus3/pkg/server"

	log "github.com/ezavalishin/versus3/internal/logger"
)

func main() {

	orm, err := orm2.Factory()

	if err != nil {
		log.Panic(err)
	}

	server.Run(orm)
}
