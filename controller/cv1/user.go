package cv1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)
type User struct{
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}

func Login(c *gin.Context){
	user := User{}
	if  c.ShouldBind(&user) == nil {
		log.Println(user.Username)
		log.Println(user.Password)
	}
	fmt.Println(user.Username)
	fmt.Println(user.Password)
	fmt.Println("login")
	c.String(200, "Success")
}