/**
 * @Title  tuotoo/qrcode读取二维码
 * @description  读取二维码
 * @Author  沈来
 * @Update  2020/8/26 22:09
 **/
package main

import (
	"fmt"
	"os"

	"github.com/tuotoo/qrcode"
)

func main() {
	fi, err := os.Open("./img/wgjjx.png")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer fi.Close()

	qrmatrix, err := qrcode.Decode(fi)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(qrmatrix.Content)
}