package main

import (
	"known-anchors/router"
	"known-anchors/service"
)

func main() {
	sc := service.NewServiceContext()
	router := router.NewRouter(sc)
	router.Run(":8080")
}
