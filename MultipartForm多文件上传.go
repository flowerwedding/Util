/**
 * @Title  MultipartForm多文件上传
 * @description  多文件上传
 * @Author  沈来
 * @Update  2020/8/21 16:23
 **/
package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	router := gin.Default()
	// 给表单限制上传大小 (默认 32 MiB)
	// router.MaxMultipartMemory = 8 << 20  // 8 MiB
	router.POST("/upload", func(c *gin.Context) {
		// 多文件
		form, _ := c.MultipartForm()
		files := form.File["file"]

		for _, file := range files {
			log.Println(file.Filename)

			// 上传文件到指定的路径
			// c.SaveUploadedFile(file, dst)
		}
		c.String(http.StatusOK, fmt.Sprintf("%d files uploaded!", len(files)))
	})
	_ = router.Run(":8080")
}