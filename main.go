package main

import (
	"fmt"
	"justtest/lib"
	"time"
)

func main() {
	start := time.Now()

	lib.Receive()

	end := time.Now()
	subS := end.Sub(start)
	fmt.Printf("------------ %.1f 秒 --------------\n", subS.Seconds()) // 保留一位小数
}

// // 生成 CPU 报告
// func cpuProfile() {
// 	f, err := os.OpenFile("cpu.prof", os.O_RDWR|os.O_CREATE, 0644)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer f.Close()

// 	log.Println("CPU Profile started")
// 	pprof.StartCPUProfile(f)
// 	defer pprof.StopCPUProfile()

// 	time.Sleep(3 * time.Second)
// 	fmt.Println("CPU Profile stopped")
// }

// // 生成堆内存报告
// func heapProfile() {
// 	f, err := os.OpenFile("heap.prof", os.O_RDWR|os.O_CREATE, 0644)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer f.Close()

// 	time.Sleep(3 * time.Second)

// 	pprof.WriteHeapProfile(f)
// 	fmt.Println("Heap Profile generated")
// }

// // 生成追踪报告
// func traceProfile() {
// 	f, err := os.OpenFile("trace.out", os.O_RDWR|os.O_CREATE, 0644)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer f.Close()

// 	log.Println("Trace started")
// 	trace.Start(f)
// 	defer trace.Stop()

// 	time.Sleep(3 * time.Second)
// 	fmt.Println("Trace stopped")
// }
