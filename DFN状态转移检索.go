/**
 * @Title  DFN状态转移检索
 * @description  状态转移检索敏感词
 * @Author  沈来
 * @Update  2020/8/13 15:16
 **/
package main

import (
    "fmt"
    "strings"
)
import "github.com/antlinker/go-dirtyfilter"
import "github.com/antlinker/go-dirtyfilter/store"

var (
	filterText = `文@@件2333333。。。`
)

func main() {
	memStore, err:= store.NewMemoryStore(store.MemoryConfig{
        DataSource: []string{"文件"},
    })
    if err != nil {
         panic(err)
    }
    filterManage:= filter.NewDirtyManager(memStore)
    results, err:= filterManage.Filter().Filter(filterText, '*', '@')
    if err != nil {
        panic(err)
    }
    fmt.Println(results)
    for _, result := range results {
        filterText = strings.ReplaceAll(filterText,result,"**")
    }
    fmt.Println(filterText)
}