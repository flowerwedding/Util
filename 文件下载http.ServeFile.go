/**
 * @Title  三行-文件下载
 * @description  #
 * @Author  沈来
 * @Update  2020/8/27 17:04
 **/
package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

//它的核心是调用了http包    http.ServeFile(c.Writer, c.Request, filepath)
//ServeFile  又调用了func serveFile(w ResponseWriter, r *Request, fs FileSystem, name string, redirect bool) 方法
func FileDownload(c *gin.Context){
	c.Writer.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=%s", "文件名-微观经济分析"))//对文件进行重命名
	c.Writer.Header().Add("Content-Type", "image/png")//value是文件转换对应的格式
	c.File("./img/wg.pdf")//文件在当前目录
}

func main(){
	r := gin.Default()

	r.GET("/download",FileDownload)

	_ = r.Run(":8000")
}