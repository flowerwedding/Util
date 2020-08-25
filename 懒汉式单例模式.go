/**
 * @Title  懒汉式单例模式
 * @description  懒汉式单例模式
 * @Author  沈来
 * @Update  2020/8/12 14:32
 **/
package main

import "sync"

type singleton struct {
	count int
}

var(
	instance *singleton
	mutex sync.Mutex
)

func New() *singleton {
	if instance == nil{//第一次检查
		//这里可能有多个 goroutine 同时到达
		mutex.Lock()
		//这里确保只有一个goroutine
		if instance == nil {//第二次检查
			instance = new(singleton)
		}
		mutex.Unlock()
	}
	return instance
}

func (s *singleton) Add() int {
	s.count ++
	return s.count
}

