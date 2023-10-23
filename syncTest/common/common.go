package common

import (
	"fmt"
	"time"
)

var Num int

func DoAdd() {
	Num++
	fmt.Printf("num is %d\n", Num)
	time.Sleep(50 * time.Millisecond)
}
