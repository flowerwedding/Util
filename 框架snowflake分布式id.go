/**
 * @Title  框架snowflake分布式id
 * @description  框架snowflake分布式id
 * @Author  沈来
 * @Update  2020/8/10 10:07
 **/
package main

import (
	"fmt"
	"github.com/bwmarrin/snowflake"
	"os"
)

func main() {
	n, err := snowflake.NewNode(1)
	if err != nil {
		println(err)
		os.Exit(1)
	}

	for i := 0; i < 3000; i++ {
		id := n.Generate()
		fmt.Println("id", id)
	}
}