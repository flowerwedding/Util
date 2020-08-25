/**
 * @Title  main
 * @description  #
 * @Author  沈来
 * @Update  2020/8/16 15:17
 **/
package main

import (
	"fmt"
	"strconv"
)

func main(){
	for i := 1;i <= 7; i++{
		p := strconv.Itoa(i)
		go sending("what  is   "+p,"1.message.2")
	}
	for i := 8;i <= 10; i++{
		p := strconv.Itoa(i)
		go sending("what  is   "+p,"entering.3.4")
	}
	for i := 11;i <= 13; i++{
		p := strconv.Itoa(i)
		go sending("what  is   "+p,"5.6.leaving")
	}
	for {
		fmt.Println(1)
	}
}