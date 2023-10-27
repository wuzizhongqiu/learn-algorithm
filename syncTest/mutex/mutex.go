package mutex

import (
	"fmt"
	"sync"
	"syncTest/common"
)

func Mutex() {
	mu := sync.Mutex{}
	var wg sync.WaitGroup
	wg.Add(10)

	for i := 0; i < 10; i++ {
		go func() {
			mu.Lock()
			common.DoAdd()
			mu.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Printf("num=%d", common.Num)
}
