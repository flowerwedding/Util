/**
 * @Title  文件分享
 * @description  网页自动打开链接
 * @Author  沈来
 * @Update  2020/8/26 20:20
 **/
package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os/exec"
	"time"
)

func main() {
	address := GetIntranetIp()
	fmt.Println("本机ip地址列表：")
	for _, item := range address {
		fmt.Println(item)
	}
	http.Handle("/", http.FileServer(http.Dir("./")))//把当前文件目录作为共享目录
	go func() {
		time.Sleep(2000)
		loclstr := fmt.Sprintf("http://%s:8080", address[0])//本机浏览器自动打开这个地址
		cmd := exec.Command("cmd", "/C", "start "+loclstr)
		_ = cmd.Run()
	}()
	if err := http.ListenAndServe(":8080", nil); err != nil {//
		fmt.Println("err:", err)
	}
}

//获取内网的地址
func GetIntranetIp() (r []string) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		log.Fatal(err)
	}
	for _, address := range addrs {
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				r = append(r, ipnet.IP.String())
			}
		}
	}
	return
}