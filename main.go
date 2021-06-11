package main

import (
	"github.com/gin-gonic/gin"
	"why/ginessential/common"
)


func main() {
	common.InitDB()

	r := gin.Default()
	r = CollectRoute(r)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

