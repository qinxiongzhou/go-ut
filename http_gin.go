package http_gin_main

import (
	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		//c.String(200, "pong")
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	return r
}

func main() {
	//去除ACC日志的打印
	gin.SetMode(gin.ReleaseMode)
	//gin.DefaultWriter = ioutil.Discard

	r := setupRouter()
	r.Run(":8080")
}
