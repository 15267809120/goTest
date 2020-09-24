package cv1

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func GetIp(c *gin.Context){
	//获取用户IP
	ip := c.ClientIP()
	fmt.Println(ip)
	fmt.Println("ip")
}