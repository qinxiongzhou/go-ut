package main

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		//c.String(200, "pong")
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/square", func(c *gin.Context) {
		num,_:= strconv.Atoi(c.Query("num"))
		result := Square(num)
		c.JSON(200,gin.H{
			"result":result,
		})
	})
	return r
}

//计算平方
func Square(num int) int{
	result := num * num
	return result
}

func main() {
	//去除ACC日志的打印
	gin.SetMode(gin.ReleaseMode)
	//gin.DefaultWriter = ioutil.Discard

	r := setupRouter()
	r.Run(":8080")
}
