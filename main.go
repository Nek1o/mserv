package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	Route(r)
	r.Run()
}
