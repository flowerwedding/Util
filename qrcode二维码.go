/**
 * @Title  qrcode二维码
 * @description  生成二维码
 * @Author  沈来
 * @Update  2020/8/26 21:43
 **/
package main

import (
	"fmt"
	"github.com/skip2/go-qrcode"
)

//这个二维码留白比较多
//filename是生成的二维码的存储路径和名字
func main() {
	err := qrcode.WriteFile("https://github.com/flowerwedding", qrcode.Medium, 256, "./img/qr.png")
	if err != nil {
		fmt.Println(err)
	}
}