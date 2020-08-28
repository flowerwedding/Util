/**
 * @Title  文件下载ioutil.ReadAll
 * @description  #
 * @Author  沈来
 * @Update  2020/8/27 19:20
 **/
package main

import (
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"net/url"
)

func download2(c *gin.Context){
    res, err :=http.Get("http://test.com/a.pdf")
    if err !=nil {
        panic(err)
    }

    content, err := ioutil.ReadAll(res.Body)
    if err != nil {
    	panic(err)
    }

    filename :=url.QueryEscape("test.pdf")// 防止中文乱码
    c.Writer.Header().Add("Content-Type", "application/octet-stream")
    c.Writer.Header().Add("Content-Disposition", "attachment; filename=\""+filename+"\"")
	_, _ = c.Writer.Write(content)
}