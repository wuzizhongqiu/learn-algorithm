package signal

import (
	"context"
	"fmt"
	"golang.org/x/sync/semaphore"
	"sync"
	"syncTest/common"
)

const (
	Limit  = 1 // 同时运行的 goroutine 上限
	Weight = 1 // 每个 goroutine 获取信号量资源的权重
)

func Signal() {
	s := semaphore.NewWeighted(Limit)
	var w sync.WaitGroup
	w.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			s.Acquire(context.Background(), Weight) // 获取信号量
			common.DoAdd()
			s.Release(Weight) // 释放信号量
			w.Done()
		}()
	}
	w.Wait()
	fmt.Printf("num = %d\n", common.Num)
}
