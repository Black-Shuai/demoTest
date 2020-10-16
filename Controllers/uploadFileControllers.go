package Controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

func Upload(ctx *gin.Context) {
	//获取文件头
	file, err := ctx.FormFile("file")
	fmt.Println(err)
	if err != nil {
		ctx.String(http.StatusBadRequest, "请求失败")
		return
	}
	//获取文件名
	fileName := file.Filename
	//保存文件到服务器本地
	//存储与项目文件夹下的file文件夹下
	filePath :="file/"+fileName
	//SaveUploadedFile(文件头，保存路径)
	if err := ctx.SaveUploadedFile(file, filePath); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Code":0,
			"Message":"上传文件失败",
			"error":err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"Code":1,
		"Message":"上传文件成功",
		"urlPath":"http://localhost:8088/api/file/getimage/"+fileName,
	})
}
func GetImage(c *gin.Context){
	imageName := c.Param("imageName")
	path:="file/"
	file, _ := ioutil.ReadFile(path+imageName)
	c.Writer.WriteString(string(file))
}