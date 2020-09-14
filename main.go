package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rampo0/samplegomod/testpackage"
)

func main() {
	testpackage.HelloWorld()

	router := gin.Default()
	router.Run(":8080")
}
