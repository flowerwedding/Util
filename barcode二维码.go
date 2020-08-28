/**
 * @Title  barcode二维码
 * @description  生成二维码
 * @Author  沈来
 * @Update  2020/8/26 21:56
 **/
package main

import (
	"image/png"
	"os"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
)

func main() {
	qrCode, _ := qr.Encode("https://github.com/flowerwedding", qr.M, qr.Auto)

	qrCode, _ = barcode.Scale(qrCode, 256, 256)

	file, _ := os.Create("./img/ba.png")//保存路径和名字

	defer file.Close()

	_ = png.Encode(file, qrCode)
}