/**
 * @Title  Seek
 * @description  #
 * @Author  沈来
 * @Update  2020/8/27 23:53
 **/
package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	/*
		    Seek(offset int64, whence int) (int64, error)，设置指针光标的位置
		        第一个参数：偏移量
		        第二个参数：如何设置
		            0：seekstart，表示相对于文件开始
		            1：seekcurrent，表示相对于当前位置的偏移量
		            2：seekend，表示相对于末尾


		            // Seek whence values.
		const (
		    SeekStart   = 0 // seek relative to the origin of the file
		    SeekCurrent = 1 // seek relative to the current offset
		    SeekEnd     = 2 // seek relative to the end
		)
	*/
	fileName := "./img/wg.pdf"
	file,err := os.OpenFile(fileName,os.O_RDWR,os.ModePerm)
	if err != nil{
		log.Fatal(err)
	}
	defer file.Close()
	//读写
	bs :=[]byte{0}
	_, _ = file.Read(bs)
	fmt.Println(string(bs))

	//相对于文件头偏移4个字节
	_, _ = file.Seek(4, io.SeekStart)
	_, _ = file.Read(bs)
	fmt.Println(string(bs))

	//相对于文件末尾
	_, _ = file.Seek(2, 0) //SeekStart
	_, _ = file.Read(bs)
	fmt.Println(string(bs))

	//相对于当前位置
	_, _ = file.Seek(3, io.SeekCurrent)
	_, _ = file.Read(bs)
	fmt.Println(string(bs))

	_, _ = file.Seek(0, io.SeekEnd)
	_, _ = file.WriteString("ABC")
}