package main

import (
	"github.com/erickrodrigs/go-webapi-template/src/application"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	application.Run(r)
	r.Run()
}
