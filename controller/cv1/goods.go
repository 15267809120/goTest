package cv1

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func Search(c *gin.Context){
	name := c.Param("name")
	fmt.Println(name)
	fmt.Println("search")
}