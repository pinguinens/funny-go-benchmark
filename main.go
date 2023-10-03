package main

import (
	"fmt"
	"runtime"
	"sync"

	"github.com/DmitriyVTitov/size"
)

const (
	multiplier = 1024 * 2
	kiloByte   = 1024
)

var (
	buffer chan [][1024]byte
)

func main() {
	fmt.Println("Funny Go Benchmark")

	cores := runtime.NumCPU()
	fmt.Printf("CPU count: %v\n", cores)

	bufferSize := kiloByte * multiplier
	bs, meter := prettyBytesValue(uint64(bufferSize * kiloByte))
	fmt.Printf("target buffer size: %.2f %v\n", bs, meter)

	go func(limit int) {
		buffer = make(chan [][1024]byte, limit)
	}(cores)

	var arr [1024]byte
	wg := sync.WaitGroup{}
	wg.Add(cores)

	routineBufSize := make(chan int, cores)
	portionSize := bufferSize / cores
	for i := 0; i < cores; i++ {
		go func(ps int) {
			defer wg.Done()

			bs, meter := prettyBytesValue(uint64(ps * kiloByte))
			fmt.Printf("target routine buffer size: %.2f %v\n", bs, meter)

			rBuffer := make([][1024]byte, 0, portionSize)
			for i := 0; i < ps; i++ {
				rBuffer = append(rBuffer, arr)
			}

			routineBufSize <- size.Of(rBuffer)
			buffer <- rBuffer
		}(portionSize)
	}

	wgB := sync.WaitGroup{}
	wgB.Add(1)
	var resultBufSize uint64
	go func() {
		defer wgB.Done()

		for s := range routineBufSize {
			bs, meter := prettyBytesValue(uint64(s))
			fmt.Printf("routine buffer size: %.2f %v\n", bs, meter)
			resultBufSize += uint64(s)
		}

		bs, meter = prettyBytesValue(resultBufSize)
		fmt.Printf("total buffer size: %.2f %v\n", bs, meter)
	}()

	wg.Wait()
	close(routineBufSize)
	wgB.Wait()
}

func prettyBytesValue(v uint64) (float32, string) {
	temp := float32(v)
	m := "Bytes"

	if temp >= kiloByte {
		temp = temp / kiloByte
		m = "KB"
	} else {
		return temp, m
	}

	if temp >= kiloByte {
		temp = temp / kiloByte
		m = "MB"
	} else {
		return temp, m
	}

	if temp >= kiloByte {
		temp = temp / kiloByte
		m = "GB"
	} else {
		return temp, m
	}

	return temp, m
}