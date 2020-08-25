/**
 * @Title  github发表评论
 * @description  #
 * @Author  沈来
 * @Update  2020/8/23 22:58
 **/
package main

import (
	"fmt"
	"github.com/afocus/captcha"
	"image/color"
	"image/png"
	"net/http"
)

func main() {
	cap := captcha.New()
	//通过句柄调用 字体文件
	if err := cap.SetFont("你字体文件的路径"); err != nil {
		panic(err.Error())
	}
	//设置图片大小
	cap.SetSize(91,50)
	//设置感染强度
	cap.SetDisturbance(captcha.NORMAL)
	cap.SetFrontColor(color.RGBA{255, 255, 255, 255})
	cap.SetBkgColor(color.RGBA{255, 0, 0, 255}, color.RGBA{0, 0, 255, 255}, color.RGBA{0, 153, 0, 255})
	http.HandleFunc("/r", func(w http.ResponseWriter, r *http.Request) {
		img,str := cap.Create(4,captcha.ALL)
		_ = png.Encode(w, img)
		fmt.Println(str)
	})
	http.HandleFunc("/c", func(w http.ResponseWriter, r *http.Request) {
		str := r.URL.RawQuery
		img := cap.CreateCustom(str)
		png.Encode(w,img)
	})
	_ = http.ListenAndServe(":9000", nil)
}