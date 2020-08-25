/**
 * @Title  bigcache
 * @description  高性能缓存，分片技术
 * @Author  沈来
 * @Update  2020/8/18 15:40
 **/
package main

import (
	"github.com/allegro/bigcache/v2"
	"log"
	"time"
)

func main(){
	cache, err := bigcache.NewBigCache(bigcache.DefaultConfig(10 * time.Minute))
	if err != nil {
		log.Println(err)
		return
	}

	entry, err := cache.Get("my-unique-key")
	if err != nil && err != bigcache.ErrEntryNotFound {//没找到会报错
		log.Println(err)
		return
	}

	if entry == nil {
		entry = []byte("value")//从数据库中获取
		_ = cache.Set("my-unique-key", entry)
	}

	log.Println(string(entry))
}