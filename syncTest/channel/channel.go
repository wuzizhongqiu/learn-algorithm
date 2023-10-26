package channel

import (
	"fmt"
	"sync"
	"syncTest/common"
)

func Channel() {
	chanLock := make(chan struct{}, 1) // 想要实现互斥，channel 的容量要设置 1
	var wg sync.WaitGroup
	wg.Add(10)

	for i := 0; i < 10; i++ {
		go func() {
			chanLock <- struct{}{}
			common.DoAdd()
			<-chanLock
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Printf("num=%d", common.Num)
}
