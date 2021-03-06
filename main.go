package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)
//定义全局中间件
func CalcTime(c *gin.Context) {
	start := time.Now()
	c.Next()
	since := time.Since(start)
	fmt.Println("花费的时间为：", since)
}

func IndexHandler(c *gin.Context){
	time.Sleep(5 * time.Second)
}

func HomeHandler(c *gin.Context){
	time.Sleep(3 * time.Second)
}

func main() {
	r := gin.Default()
	//实现全局中间件
	r.Use(CalcTime)
	routerGroup := r.Group("/begin")
	{	
		routerGroup.GET("/index",IndexHandler)
		routerGroup.GET("/home",HomeHandler)
	}
	r.Run(":8080")

}
