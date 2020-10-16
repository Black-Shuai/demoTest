package Controllers

import (
	"demoTest/Models"
	"demoTest/Services"
	"demoTest/Services/ServiceModels"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)
type ResultResponse struct {
	Code int
	Message string
	Data interface{}
}
var user Models.User
//var userList []Models.User
//查找登录用户
func Login(c *gin.Context) {
	//数据解析，将json格式数据解析为结构体
	err := c.ShouldBindJSON(&user)
	var _ ServiceModels.ResultResponse
	 res:=Services.FindUser(user)
	 if err==nil{
		 fmt.Println(user)
	 	c.JSON(http.StatusOK,res)

	 }
	 return
}