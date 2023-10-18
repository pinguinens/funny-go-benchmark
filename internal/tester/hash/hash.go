package hash

import (
	"sync"
	"time"

	"github.com/pinguinens/funny-go-benchmark/pkg/log"
	"github.com/pinguinens/funny-go-benchmark/pkg/timer"

	"github.com/pinguinens/funny-go-benchmark/internal/buffer"
	"github.com/pinguinens/funny-go-benchmark/internal/tester/hash/payload"
	"github.com/pinguinens/funny-go-benchmark/internal/tester/hash/test"
)

type Test interface {
	Name() string
	Exec(b []byte) (int, []byte, error)
}

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

func (t *Tester) Run() {
	t.logger.Infoln("Hash Test...")
	t.logger.Infoln("hash | time | hash rate | - byte count:hash string")

	testload := payload.New()

	t.execute(&test.Adler32{}, testload)
	t.execute(&test.CRC32{}, testload)
	t.execute(&test.CRC64{}, testload)
	t.execute(&test.FNV32{}, testload)
	t.execute(&test.FNV32a{}, testload)
	t.execute(&test.FNV64{}, testload)
	t.execute(&test.FNV64a{}, testload)
	t.execute(&test.FNV128{}, testload)
	t.execute(&test.FNV128a{}, testload)
	t.execute(&test.MD5{}, testload)
	t.execute(&test.SHA512{}, testload)
}

func (t *Tester) execute(ts Test, data []byte) {
	t.logger.Infof("%v | ", ts.Name())
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
				n, r, err = ts.Exec(data)
				if err != nil {
					t.logger.Infoln(err)
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
	t.logger.Infof("%v | %v | - %v:%x\n", t.timer.Result(), count, n, r)
}
