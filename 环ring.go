/**
 * @Title  环ring
 * @description  container/ring包
 * @Author  沈来
 * @Update  2020/8/14 16:42
 **/
package main

import (
	"container/ring"
	"fmt"
)

func main() {
	r := ring.New(5)
	n := r.Len()
	for i := 1; i <= n; i++ {
		r.Value = i
		r = r.Next()
	}

	r.Do(func (value interface{}){
		fmt.Println(value.(int))
	})
}