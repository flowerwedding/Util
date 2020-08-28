/**
 * @Title  文件下载io.Copy
 * @description  #
 * @Author  沈来
 * @Update  2020/8/27 19:21
 **/
package main

import (
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"net/url"
)

func download3(c *gin.Context){
	res, err :=http.Get("http://test.com/a.pdf")
	if err !=nil {
		panic(err)
	}

	filename :=url.QueryEscape("test.pdf")// 防止中文乱码
	c.Writer.Header().Add("Content-Type", "application/octet-stream")
	c.Writer.Header().Add("Content-Disposition", "attachment; filename=\""+filename+"\"")
	_, _ = io.Copy(c.Writer, res.Body)
}