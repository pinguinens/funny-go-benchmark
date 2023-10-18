package hash

import (
	"crypto/md5"
	"crypto/sha512"
	"hash/crc64"
	"hash/fnv"
	"sync"
	"time"

	"github.com/pinguinens/funny-go-benchmark/pkg/log"
	"github.com/pinguinens/funny-go-benchmark/pkg/timer"

	"github.com/pinguinens/funny-go-benchmark/internal/buffer"
	"github.com/pinguinens/funny-go-benchmark/internal/tester/hash/payload"
	"github.com/pinguinens/funny-go-benchmark/internal/tester/hash/test"
)

type Tester struct {
	logger *log.Logger
	buffer *buffer.Buffer
	timer  *timer.Timer
}

func New(logger *log.Logger, buffer *buffer.Buffer, routines int) *Tester {
	return &Tester{
		logger: logger,
		buffer: buffer,
		timer:  &timer.Timer{},
	}
}

func (t *Tester) execute(data []byte) {
	t.logger.Info("---\n CRC32")
	var (
		n       int
		r       []byte
		err     error
		counter uint64
	)
	t.timer.Start()

	wg := sync.WaitGroup{}
	wg.Add(1)

	cont := true
	go func() {
		go func() {
			for cont {
				n, r, err = test.CRC32(data)
				if err != nil {
					t.logger.Info(err)
				}

				counter++
			}
		}()

		time.Sleep(1 * time.Second)
		wg.Done()
	}()

	wg.Wait()
	cont = false
	count := counter

	t.timer.Stop()
	t.logger.Infof("- %v:%x", n, r)
	t.logger.Infof("time: %v", t.timer.Result())
	t.logger.Infof("count: %v", count)
}

func (t *Tester) Run() {
	t.logger.Info("Hash Test...")

	testload := payload.New()
	ti := timer.Timer{}

	test.Adler32(t.logger, &ti, testload)
	t.execute(testload)

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

	{
		t.logger.Info("---\n MD5")
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
					hObj := md5.New()
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
		t.logger.Info("---\n SHA-512")
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
					hObj := sha512.New()
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
