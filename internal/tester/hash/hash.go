package hash

import (
	"github.com/pinguinens/funny-go-benchmark/internal/buffer"
	"github.com/pinguinens/funny-go-benchmark/pkg/log"
	"github.com/pinguinens/funny-go-benchmark/pkg/timer"
	"hash/adler32"
	"hash/crc32"
	"hash/crc64"
	"hash/fnv"
	"sync"
	"time"
)

type Tester struct {
	logger *log.Logger
	buffer *buffer.Buffer
}

func New(logger *log.Logger, buffer *buffer.Buffer, routines int) *Tester {
	return &Tester{
		logger: logger,
		buffer: buffer,
	}
}

func (t *Tester) Run() {
	t.logger.Info("Hash Test...")

	uuid := []byte("628aace1-ad98-4f44-8a5c-cb73862c46c4")
	payload := []byte{123, 34, 116, 97, 115, 107, 34, 58, 34, 116, 101, 115, 116, 34, 125}
	testload := append(uuid, payload...)
	ti := timer.Timer{}

	{
		t.logger.Info("---\n Adler32")
		var (
			n       int
			r       []byte
			err     error
			counter uint64
		)
		ti.Start()

		wg := sync.WaitGroup{}
		wg.Add(1)

		cont := true
		go func() {
			go func() {
				for cont {
					hObj := adler32.New()
					n, err = hObj.Write(testload)
					if err != nil {
						t.logger.Info(err)
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

		ti.Stop()
		t.logger.Infof("- %v:%x", n, r)
		t.logger.Infof("time: %v", ti.Result())
		t.logger.Infof("count: %v", count)
	}

	{
		t.logger.Info("---\n CRC32")
		var (
			n       int
			r       []byte
			err     error
			counter uint64
		)
		ti.Start()

		wg := sync.WaitGroup{}
		wg.Add(1)

		cont := true
		go func() {
			go func() {
				for cont {
					hObj := crc32.New(crc32.IEEETable)
					n, err = hObj.Write(testload)
					if err != nil {
						t.logger.Info(err)
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

		ti.Stop()
		t.logger.Infof("- %v:%x", n, r)
		t.logger.Infof("time: %v", ti.Result())
		t.logger.Infof("count: %v", count)
	}

	{
		t.logger.Info("---\n CRC64")
		var (
			n       int
			r       []byte
			err     error
			counter uint64
		)
		ti.Start()

		wg := sync.WaitGroup{}
		wg.Add(1)

		cont := true
		go func() {
			go func() {
				for cont {
					hObj := crc64.New(crc64.MakeTable(crc64.ISO))
					n, err = hObj.Write(testload)
					if err != nil {
						t.logger.Info(err)
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

		ti.Stop()
		t.logger.Infof("- %v:%x", n, r)
		t.logger.Infof("time: %v", ti.Result())
		t.logger.Infof("count: %v", count)
	}

	{
		t.logger.Info("---\n FNV-1 32")
		var (
			n       int
			r       []byte
			err     error
			counter uint64
		)
		ti.Start()

		wg := sync.WaitGroup{}
		wg.Add(1)

		cont := true
		go func() {
			go func() {
				for cont {
					hObj := fnv.New32()
					n, err = hObj.Write(testload)
					if err != nil {
						t.logger.Info(err)
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

		ti.Stop()
		t.logger.Infof("- %v:%x", n, r)
		t.logger.Infof("time: %v", ti.Result())
		t.logger.Infof("count: %v", count)
	}

	{
		t.logger.Info("---\n FNV-1 32a")
		var (
			n       int
			r       []byte
			err     error
			counter uint64
		)
		ti.Start()

		wg := sync.WaitGroup{}
		wg.Add(1)

		cont := true
		go func() {
			go func() {
				for cont {
					hObj := fnv.New32a()
					n, err = hObj.Write(testload)
					if err != nil {
						t.logger.Info(err)
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

		ti.Stop()
		t.logger.Infof("- %v:%x", n, r)
		t.logger.Infof("time: %v", ti.Result())
		t.logger.Infof("count: %v", count)
	}

	{
		t.logger.Info("---\n FNV-1 64")
		var (
			n       int
			r       []byte
			err     error
			counter uint64
		)
		ti.Start()

		wg := sync.WaitGroup{}
		wg.Add(1)

		cont := true
		go func() {
			go func() {
				for cont {
					hObj := fnv.New64()
					n, err = hObj.Write(testload)
					if err != nil {
						t.logger.Info(err)
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

		ti.Stop()
		t.logger.Infof("- %v:%x", n, r)
		t.logger.Infof("time: %v", ti.Result())
		t.logger.Infof("count: %v", count)
	}

	{
		t.logger.Info("---\n FNV-1 64a")
		var (
			n       int
			r       []byte
			err     error
			counter uint64
		)
		ti.Start()

		wg := sync.WaitGroup{}
		wg.Add(1)

		cont := true
		go func() {
			go func() {
				for cont {
					hObj := fnv.New64a()
					n, err = hObj.Write(testload)
					if err != nil {
						t.logger.Info(err)
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

		ti.Stop()
		t.logger.Infof("- %v:%x", n, r)
		t.logger.Infof("time: %v", ti.Result())
		t.logger.Infof("count: %v", count)
	}

	{
		t.logger.Info("---\n FNV-1 128")
		var (
			n       int
			r       []byte
			err     error
			counter uint64
		)
		ti.Start()

		wg := sync.WaitGroup{}
		wg.Add(1)

		cont := true
		go func() {
			go func() {
				for cont {
					hObj := fnv.New128()
					n, err = hObj.Write(testload)
					if err != nil {
						t.logger.Info(err)
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

		ti.Stop()
		t.logger.Infof("- %v:%x", n, r)
		t.logger.Infof("time: %v", ti.Result())
		t.logger.Infof("count: %v", count)
	}

	{
		t.logger.Info("---\n FNV-1 128a")
		var (
			n       int
			r       []byte
			err     error
			counter uint64
		)
		ti.Start()

		wg := sync.WaitGroup{}
		wg.Add(1)

		cont := true
		go func() {
			go func() {
				for cont {
					hObj := fnv.New128a()
					n, err = hObj.Write(testload)
					if err != nil {
						t.logger.Info(err)
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

		ti.Stop()
		t.logger.Infof("- %v:%x", n, r)
		t.logger.Infof("time: %v", ti.Result())
		t.logger.Infof("count: %v", count)
	}
}
