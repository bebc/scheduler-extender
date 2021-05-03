package main

import (
	"github.com/gin-gonic/gin"
	"scheduler-extender/pkg"
)

func main()  {
	r := gin.Default()
	r.POST("/extender/filter",pkg.Filter)
	r.POST("/extender/prioritize",pkg.Prioritize)
	r.Run(":8080")
}
