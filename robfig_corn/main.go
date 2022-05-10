package main

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"time"
)

// 返回一个支持至 秒 级别的 cron
func newWithSeconds() *cron.Cron {
	secondParser := cron.NewParser(cron.Second | cron.Minute |
		cron.Hour | cron.Dom | cron.Month | cron.DowOptional | cron.Descriptor)
	return cron.New(cron.WithParser(secondParser), cron.WithChain())
}

func main() {
	fmt.Println("starting go cron...")

	i := 0
	//这个只支持精确到分钟级别
	cron := cron.New()
	//支持精确到秒级别
	//cron := newWithSeconds()
	spec1 := "*/1 * * * *"
	//直接闭包函数
	cron.AddFunc(spec1, func() {
		i++
		fmt.Println("cron is running...", i)
	})

	spec2 := "*/1 * * * *"
	cron.AddFunc(spec1, Test1) //运行函数Test1
	cron.AddFunc(spec2, Test2) //运行函数Test2

	//启动计划任务
	cron.Start()

	//关闭着计划任务, 但是不能关闭已经在执行中的任务.
	defer cron.Stop()
	//time.Sleep(time.Minute) //一分钟后主线程退出
	select {}
}

func CronTest() {
	c := cron.New()
	c.AddFunc("*/1 * * * * ?", func() { fmt.Println("Every hour on the half hour1") })
	//c.AddFunc("0 30 * * * *", func() { fmt.Println("Every hour on the half hour") })
	//c.AddFunc("@hourly", func() { fmt.Println("Every hour") })
	//c.AddFunc("@every 1h30m", func() { fmt.Println("Every hour thirty") })
	//c.AddFunc("@every 5s", func() { fmt.Println("Every hour thirty") })
	c.Start()
	//..
	// Funcs are invoked in their own goroutine, asynchronously.
	//...
	// Funcs may also be added to a running Cron
	//c.AddFunc("@daily", func() { fmt.Println("Every day") })
	//..
	// Inspect the cron job entries' next and previous run times.
	//inspect(c.Entries())
	//..
	//c.Stop() // Stop the scheduler (does not stop any jobs already running).
	select {}
}

func CronTest2() {
	fmt.Println("starting go cron...")

	i := 0
	cron := cron.New()
	spec1 := "*/1 * * * * ?"
	//直接闭包函数
	cron.AddFunc(spec1, func() {
		i++
		fmt.Println("cron is running...", i)
	})

	spec2 := "*/1 * * * * ?"
	cron.AddFunc(spec1, Test1) //运行函数Test1
	cron.AddFunc(spec2, Test2) //运行函数Test2

	//启动计划任务
	cron.Start()

	//关闭着计划任务, 但是不能关闭已经在执行中的任务.
	defer cron.Stop()

	select {}

}

func Test1() {
	time.Sleep(10 * time.Second)
	fmt.Println("Test1...")
}
func Test2() {
	fmt.Println("Test2...")
}
