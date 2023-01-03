package main

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"time"
)

func main() {
	c := cron.New()

	// 1、调用cron对象的AddFunc()方法向管理器中添加定时任务。
	// AddFunc()接受两个参数，参数 1 以字符串形式指定触发时间规则，参数 2 是一个无参的函数，每次触发时调用。
	// @every 1s表示每秒触发一次，@every后加一个时间间隔，表示每隔多长时间触发一次。
	c.AddFunc("@every 1s", func() {
		fmt.Println("tick every 1 second")
	})

	// 2、除了直接将无参函数作为回调外，cron还支持Job接口：
	c.AddJob("@every 1s", GreetingJob{Name: "超哥"}) // 调用cron对象的AddJob()方法将GreetingJob对象添加到定时管理器中

	// 调用c.Start()启动定时循环。
	c.Start()                   // 启动一个新的 goroutine 做循环检测
	time.Sleep(time.Second * 5) // 防止主 goroutine 退出
}

// Job 接口
type Job interface {
	Run()
}

// GreetingJob 定义一个实现接口Job的结构：
type GreetingJob struct {
	Name string
}

func (g GreetingJob) Run() {
	fmt.Println("Hello", g.Name)
}
