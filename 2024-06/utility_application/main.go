package main

import (
	"bytes"
	"flag"
	"math/rand"
	"os"
	"runtime"
	"runtime/pprof"
	"sync"
	"time"
)

func main() {

	var isCPUPprof bool
	var isMemPprof bool

	// 传入 -cpu -mem 参数 生成文件
	flag.BoolVar(&isCPUPprof, "cpu", false, "turn cpu pprof on")
	flag.BoolVar(&isMemPprof, "mem", false, "turn mem pprof on")

	flag.Parse()

	if isCPUPprof {
		cpuProf, err := os.Create("cpu.prof")
		if err != nil {
			panic(err)
		}
		err = pprof.StartCPUProfile(cpuProf) // 开始cpu profile
		if err != nil {
			panic(err)
		}
		defer cpuProf.Close()        // 关闭文件
		defer pprof.StopCPUProfile() // 停止cpu profile
	}

	var wg sync.WaitGroup
	wg.Add(200)

	for range 200 {
		go cycleNum(30000, &wg)
	}

	writeBytes()

	wg.Wait()

	time.Sleep(1 * time.Second)

	if isMemPprof {
		memProf, err := os.Create("mem.prof")
		if err != nil {
			panic(err)
		}

		runtime.GC()

		defer memProf.Close()
		defer pprof.WriteHeapProfile(memProf) // 写入内存profile
	}
}

func cycleNum(num int, wg *sync.WaitGroup) {
	slice := make([]int, 0)
	for i := 0; i < num; i++ {
		for j := 0; j < num; j++ {
			j = i + j
			slice = append(slice, j)
		}
	}
	wg.Done()
}

func writeBytes() *bytes.Buffer {
	var buff bytes.Buffer

	for range 30000 {
		buff.Write([]byte{'0' + byte(rand.Intn(10))})
	}
	return &buff
}
