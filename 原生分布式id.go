/**
 * @Title  id
 * @description  snowflake分布式id
 * @Author  沈来
 * @Update  2020/8/10 9:43
 **/
package main

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

const (
	twePoch        = int64(1483228800000)             //开始时间截 (2017-01-01)
	workerIDBits   = uint(10)                         //机器id所占的位数
	sequenceBits   = uint(12)                         //序列所占的位数
	workerIDMax    = int64(-1 ^ (-1 << workerIDBits)) //支持的最大机器id数量
	sequenceMask   = int64(-1 ^ (-1 << sequenceBits)) //
	workerIDShift  = sequenceBits                     //机器id左移位数
	timestampShift = sequenceBits + workerIDBits      //时间戳左移位数
)

// A Snowflake struct holds the basic information needed for a snowflake generator worker
type Snowflake struct {
	sync.Mutex
	timestamp int64
	workerid  int64
	sequence  int64
}

// NewNode returns a new snowflake worker that can be used to generate snowflake IDs
func NewSnowflake(workerid int64) (*Snowflake, error) {

	if workerid < 0 || workerid > workerIDMax {
		return nil, errors.New("workerid must be between 0 and 1023")
	}

	return &Snowflake{
		timestamp: 0,
		workerid:  workerid,
		sequence:  0,
	}, nil
}

// Generate creates and returns a unique snowflake ID
func (s *Snowflake) Generate() int64 {

	s.Lock()

	now := time.Now().UnixNano() / 1000000

	if s.timestamp == now {
		s.sequence = (s.sequence + 1) & sequenceMask

		if s.sequence == 0 {
			for now <= s.timestamp {
				now = time.Now().UnixNano() / 1000000
			}
		}
	} else {
		s.sequence = 0
	}

	s.timestamp = now

	r := int64((now-twePoch)<<timestampShift | (s.workerid << workerIDShift) | (s.sequence))

	s.Unlock()
	return r
}

func main() {
	// 测试脚本

	// 生成节点实例
	worker, err := NewSnowflake(2)

	if err != nil {
		fmt.Println(err)
		return
	}

	ch := make(chan int64)
	count := 3000
	// 并发 count 个 goroutine 进行 snowflake ID 生成
	for i := 0; i < count; i++ {
		go func() {
			id := worker.Generate()
			ch <- id
		}()
	}

	defer close(ch)

	m := make(map[int64]int)
	for i := 0; i < count; i++  {
		id := <- ch
		// 如果 map 中存在为 id 的 key, 说明生成的 snowflake ID 有重复
		_, ok := m[id]
		if ok {
			fmt.Println("ID is not unique!")
		}
		// 将 id 作为 key 存入 map
		m[id] = i
		fmt.Println(id)
	}
	// 成功生成 snowflake ID
	fmt.Println("All", count, "snowflake ID Get succeeded!")
}