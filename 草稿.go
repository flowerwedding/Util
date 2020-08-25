/**
 * @Title  草稿
 * @description  就是草稿
 * @Author  沈来
 * @Update  2020/8/10 14:48
 **/
package main

import (
	"github.com/gin-gonic/gin"
	"myUtil/captcha/controller"
)
func main() {
	r := gin.Default()
	r.GET("/getCaptcha",controller.GetCaptcha)
	r.GET("/verifyCaptcha", controller.VerifyCaptcha)
	r.GET("/show/:source", controller.GetCaptchaPng)
	_ = r.Run(":8080")
}