package testpackage

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func HelloWorld() {
	fmt.Println("Hello world")

	router := gin.Default()
	router.Run(":8080")
}
