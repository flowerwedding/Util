/**
 * @Title  简单替换或正则替换
 * @description  敏感词处理
 * @Author  沈来
 * @Update  2020/8/13 15:09
 **/
package main

import (
	"fmt"
	"strings"
)

func main() {
	keywords := []string{"坏蛋","坏人","发票","傻子","傻大个","傻人"}
	content := "不要发票，你就是一个傻子，只会发呆"
	for _, keyword := range keywords {
		content = strings.ReplaceAll(content,keyword,"**")
	}
	fmt.Println(content)

	replacer := strings.NewReplacer("坏蛋","**","坏人","**","发票","**","傻子","**","傻大个","**","傻人","**")
	fmt.Println(replacer.Replace("不要发票，你就是一个傻子，只会发呆"))
}