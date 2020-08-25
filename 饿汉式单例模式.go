/**
 * @Title  饿汉式单例模式
 * @description  饿汉式单例模式
 * @Author  沈来
 * @Update  2020/8/12 14:32
 **/
package main

import "fmt"

type singleton struct {
	count int
}

var Instance = new(singleton)

func (s *singleton) Add() int {
	s.count ++
	return s.count
}

func main() {
	c := Instance.Add()
	fmt.Println(c)
}