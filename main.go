package main

import (
	"fhw.com/controller/cv1"
	"github.com/gin-gonic/gin"
)






func main(){
	router := gin.Default()
	v1 := router.Group("v1")
	{
		v1.POST("user/login", cv1.Login)
		v1.GET("goods/search/name/:name", cv1.Search)
		v1.GET("/ip", cv1.GetIp)
	}
	//在8080端口上监听
	router.Run(":8080")
}
