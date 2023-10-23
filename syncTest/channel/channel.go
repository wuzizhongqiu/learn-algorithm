package channel

import "sync"

func Channel() {
	chanLock := make(chan struct{}, 1) // 想要实现互斥，channel 的容量要设置 1
	var wg sync.WaitGroup
	wg.Add(10)
}
