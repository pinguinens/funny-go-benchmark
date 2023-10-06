package main

import (
	"sync"
	"time"
	"unsafe"

	"github.com/pinguinens/funny-go-benchmark/internal/config"

	"github.com/pinguinens/funny-go-benchmark/pkg/log"
	"github.com/pinguinens/funny-go-benchmark/pkg/units"
)

var (
	buffer chan [][units.Kilobyte]byte
)

func main() {
	appCfg := config.New()

	logger := log.Logger{}

	logger.Info("Funny Go Benchmark")
	logger.Infof("CPU count: %v", appCfg.System.CoreCount)

	bufferSize := units.Kilobyte * appCfg.Buffer.Multiplier
	bs, unit := units.FormatPrettyBytes(bufferSize * units.Kilobyte)
	logger.Infof("target buffer size: %.2f %v\n", bs, unit)

	timer := time.Now()

	buffer = make(chan [][units.Kilobyte]byte, appCfg.System.CoreCount)

	var arr [units.Kilobyte]byte
	wg := sync.WaitGroup{}
	wg.Add(appCfg.System.CoreCount)
	routineBufSize := make(chan int, appCfg.System.CoreCount)
	portionSize := bufferSize / appCfg.System.CoreCount
	for i := 0; i < appCfg.System.CoreCount; i++ {
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
