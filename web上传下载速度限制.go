/**
 * @Title  web上传下载速度限制
 * @description  速度限制器ratelimit
 * @Author  沈来
 * @Update  2020/8/27 21:59
 **/
package main

import (
	"bytes"
	"fmt"
	"io"
	"time"

	"gitee.com/kzquu/wego/util/ratelimit"
)

type writer struct {
}

func (w *writer) Write(p []byte) (n int, err error) {
	time.Sleep(time.Second / 5) // 模拟网络传输消耗的时间
	return len(p), nil
}

func main() {
	// 源数据  1000 KB
	// 实验用的源数据不建议太大，因为io.Copy内部一次copy 32KB，源数据太大导致copy次数过多影响实验效果
	// 就像本来网速支持每秒10KB，你用ratelimit限制100KB每秒是完全没效果的
	src := bytes.NewReader(make([]byte, 1000*1024))
	// 模拟客户端
	dst := &writer{}

	// new 一个速度限制器，每秒传输 100 Kb
	bucket := ratelimit.New(100 * 1024)

	start := time.Now()

	// 大概需要 10s 传输完
	//n, err := io.Copy(dst, ratelimit.Reader(src, bucket))
	n, err := io.Copy(ratelimit.Writer(dst, bucket), src)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Copied %d bytes in %s\n", n, time.Since(start))

	// 输出：
	// Copied 1024000 bytes in 10.2005835s
	// 多出来的 0.2s 其实是最后一次 write 消耗的时间
}