package main

import (
	"fmt"
	"runtime"
	"time"
)

func child() {
	for i := 'a'; i < 'z'; i++ {
		fmt.Printf("%v\n", i)
		time.Sleep(1 * time.Millisecond)
	}
}

func parent() {
	for i := 'a'; i < 'z'; i++ {
		fmt.Printf("%d\n", i)
		time.Sleep(1 * time.Millisecond)
	}
}

func main() {
	go func() {
		go child()
		for i := 'a'; i < 'z'; i++ {
			fmt.Printf("%d\n", i)
		}
	}()

	cpuNum := runtime.NumCPU()
	fmt.Println("cpuNum:", cpuNum)
	runtime.GOMAXPROCS(cpuNum / 2)

	const P = 1000000
	for i := 0; i < P; i++ {
		go time.Sleep(10 * time.Second)
	}
	fmt.Println("进程中存活的协程数，", runtime.NumGoroutine())
}
