/**
 * @Title  断点续传
 * @description  #
 * @Author  沈来
 * @Update  2020/8/27 23:45
 **/
package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	/*
	   断点续传
	   C:\Users\12905\Desktop\新建文件夹\goproject\src\owen\main\test.jpg
	   复制到当前工程下
	*/
	srcFile:="./img/wg.pdf"//源文件目录//参数1，源文件来源
	defFile:=srcFile[strings.LastIndex(srcFile,"/")+1:]//源文件后缀名
	fmt.Println(defFile)

	//源文件
	file1,err:=os.Open(srcFile)//打开文件建立链接
	Modelerr(err)//错误显示

	//目标文件
	goalsrc:=srcFile+"wg.pdf"//参数2，目标文件名
	file2,err:=os.OpenFile(goalsrc,os.O_CREATE|os.O_WRONLY,os.ModePerm)
	Modelerr(err)

	//临时文件
	tmpFile:=defFile+"temp.txt"
	file3,err:=os.OpenFile(tmpFile,os.O_CREATE|os.O_RDWR,os.ModePerm)
	Modelerr(err)

	_, _ = file3.Seek(0, io.SeekStart)

	defer file1.Close()
	defer file2.Close()

	bs:=make([]byte,100,100)
	n1,err:=file3.Read(bs)
	//Modelerr(err)
	countStr:=string(bs[:n1])
	count,err:=strconv.ParseInt(countStr,10,64)
	//Modelerr(err)
	fmt.Println(count)


	//step2:设置读，写的位置
	_, _ = file1.Seek(count, io.SeekStart)

	_, _ = file2.Seek(count, io.SeekStart)

	data:=make([]byte,1024,1024)//参数3？每次续传的量

	n2:=-1//读取的数据量
	n3:=-1//写出的数据量

	total:=int(count)//读取的总量

	//step3
	for{
		n2,err=file1.Read(data)
		if err==io.EOF||n2==0{
			fmt.Println("文件复制完毕。。。")
			file3.Close()
			_ = os.Remove(tmpFile)
			break
		}
		n3,err=file2.Write(data[:n2])
		total+=n3

		//将复制的总量，存储到临时文件中
		_, _ = file3.Seek(0, io.SeekStart)
		_, _ = file3.WriteString(strconv.Itoa(total))

		fmt.Println("total 总量：",total)
	}
}

func Modelerr(err error){
	if err!=nil{
		panic(err)
	}

}