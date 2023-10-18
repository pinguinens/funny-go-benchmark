package test

import (
	"hash/adler32"
	"hash/crc32"
	"sync"
	"time"

	"github.com/pinguinens/funny-go-benchmark/pkg/log"
	"github.com/pinguinens/funny-go-benchmark/pkg/timer"
)

func Adler32(l *log.Logger, t *timer.Timer, b []byte) {
	l.Info("---\n Adler32")
	var (
		n       int
		r       []byte
		err     error
		counter uint64
	)
	t.Start()

	wg := sync.WaitGroup{}
	wg.Add(1)

	cont := true
	go func() {
		go func() {
			for cont {
				hObj := adler32.New()
				n, err = hObj.Write(b)
				if err != nil {
					l.Info(err)
				}
				r = hObj.Sum(nil)

				counter++
			}
		}()

		time.Sleep(1 * time.Second)
		wg.Done()
	}()

	wg.Wait()
	cont = false
	count := counter

	t.Stop()
	l.Infof("- %v:%x", n, r)
	l.Infof("time: %v", t.Result())
	l.Infof("count: %v", count)
}

func CRC32(b []byte) (int, []byte, error) {
	hObj := crc32.New(crc32.IEEETable)
	n, err := hObj.Write(b)
	if err != nil {
		return 0, nil, err
	}
	r := hObj.Sum(nil)

	return n, r, nil
}
