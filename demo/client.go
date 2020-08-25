/**
 * @Title  client
 * @description  tcp聊天室的客户端
 * @Author  沈来
 * @Update  2020/8/10 10:34
 **/
package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main(){
	conn, err := net.Dial("tcp",":9000")
	if err != nil {
		panic(err)
	}

	done := make(chan struct{})
	go func() {
		_, _ = io.Copy(os.Stdout, conn)
		log.Println("done")
		done <- struct{}{}
	}()

	mustCopy(conn, os.Stdin)
	conn.Close()
	<-done
}

func mustCopy(dst io.Writer, src io.Reader){
	if _, err := io.Copy(dst, src); err != nil{
		log.Fatal(err)
	}
}