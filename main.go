package main

import "github.com/gin-gonic/gin"

func main() {
    r := gin.Default()

    r.GET("/local_conf", func(c *gin.Context) {
        c.File("local.conf")
    })

    r.Run(":8080")
}
