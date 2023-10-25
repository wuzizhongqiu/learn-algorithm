package goroutinePool

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

type Task struct {
	f func() error // 需要消费的任务（由用户实现任务内容）
}

// 创建新任务
func NewTask(funcArg func() error) *Task {
	return &Task{
		f: funcArg,
	}
}

// 协程池的数据结构
type Pool struct {
	RunningWorkers int64      // 正在工作的协程
	Capacity       int64      // 协程池的容量
	JobCh          chan *Task // 用于 Worker 拉取任务
	sync.Mutex
}

// 创建协程池（协程容量, 管道大小）
func NewPool(capacity int64, taskNum int) *Pool {
	return &Pool{
		Capacity: capacity,
		JobCh:    make(chan *Task, taskNum),
	}
}

// 一些 api 接口

// 获取协程池的容量大小
func (p *Pool) GetCap() int64 {
	return p.Capacity
}

// 工作协程+1
func (p *Pool) addRunning() {
	atomic.AddInt64(&p.RunningWorkers, 1)
}

// 工作斜插袋-1
func (p *Pool) decRunning() {
	atomic.AddInt64(&p.RunningWorkers, -1)
}

// 获取正在工作的协程数量
func (p *Pool) GetRunningWorkers() int64 {
	return atomic.LoadInt64(&p.RunningWorkers)
}

// 启用一个 Worker 消费任务
func (p *Pool) Run() {
	p.addRunning()
	go func() { // 起一个协程
		defer func() {
			p.decRunning()
		}()
		for task := range p.JobCh { // 从管道中取任务
			task.f() // 消费任务
		}
	}()
}

// 往协程池里塞任务
func (p *Pool) AddTask(task *Task) {
	p.Lock()
	defer p.Unlock()

	if p.GetRunningWorkers() < p.GetCap() { // 如果协程池满了，就不再创建协程了
		p.Run() // 启动一个 Worker
	}

	p.JobCh <- task // 将任务推入队列，等待消费
}

// 测试协程池的 demo 代码
func TestPool() {
	pool := NewPool(3, 10) // 创建协程池

	for i := 0; i < 20; i++ {
		pool.AddTask(NewTask(func() error { // 创建任务
			fmt.Println("I am Task")
			return nil
		}))
	}

	time.Sleep(1e9)
}
