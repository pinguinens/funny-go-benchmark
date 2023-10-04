package main

import (
	"flag"
	"runtime"
	"sync"
	"time"
	"unsafe"

	"github.com/pinguinens/funny-go-benchmark/pkg/log"
	"github.com/pinguinens/funny-go-benchmark/pkg/units"
)

var (
	buffer chan [][units.Kilobyte]byte
)

func main() {
	var (
		multiplier int
		cores      int
	)
	flag.IntVar(&multiplier, "b", 1, "buffer memory size in MB")
	flag.IntVar(&cores, "c", runtime.NumCPU(), "CPU cores")
	flag.Parse()

	logger := log.Logger{}

	logger.Info("Funny Go Benchmark")
	logger.Infof("CPU count: %v", cores)

	bufferSize := units.Kilobyte * multiplier
	bs, unit := units.FormatPrettyBytes(bufferSize * units.Kilobyte)
	logger.Infof("target buffer size: %.2f %v\n", bs, unit)

	timer := time.Now()

	buffer = make(chan [][units.Kilobyte]byte, cores)

	var arr [units.Kilobyte]byte
	wg := sync.WaitGroup{}
	wg.Add(cores)
	routineBufSize := make(chan int, cores)
	portionSize := bufferSize / cores
	for i := 0; i < cores; i++ {
		go func(ps, id int) {
			defer wg.Done()

			bs, unit := units.FormatPrettyBytes(ps * units.Kilobyte)
			logger.Infof("[%v] target routine buffer size: %.2f %v\n", id, bs, unit)

			var rSize int
			rBuffer := make([][units.Kilobyte]byte, 0, portionSize)
			for i := 0; i < ps; i++ {
				for _, in := range arr {
					rSize += int(unsafe.Sizeof(in))
				}
				rBuffer = append(rBuffer, arr)
			}

			routineBufSize <- rSize
			buffer <- rBuffer
		}(portionSize, i)
	}

	wgB := sync.WaitGroup{}
	wgB.Add(1)
	var resultBufSize int
	go func() {
		defer wgB.Done()

		for s := range routineBufSize {
			bs, unit := units.FormatPrettyBytes(s)
			logger.Infof("routine buffer size: %.2f %v\n", bs, unit)
			resultBufSize += s
		}

		bs, unit = units.FormatPrettyBytes(resultBufSize)
		logger.Infof("total buffer size: %.2f %v\n", bs, unit)
	}()

	wg.Wait()
	close(routineBufSize)
	wgB.Wait()

	logger.Infof("test duration: %v\n", time.Now().Sub(timer))
}
