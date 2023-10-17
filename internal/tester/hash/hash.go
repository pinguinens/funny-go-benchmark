package hash

import (
	"github.com/pinguinens/funny-go-benchmark/internal/buffer"
	"github.com/pinguinens/funny-go-benchmark/pkg/log"
	"github.com/pinguinens/funny-go-benchmark/pkg/timer"
	"hash/adler32"
	"hash/crc32"
	"hash/crc64"
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
		t.logger.Info("start Adler32")
		var (
			n   int
			r   []byte
			err error
		)
		ti.Start()

		for i := 0; i < 10240; i++ {
			hObj := adler32.New()
			n, err = hObj.Write(testload)
			if err != nil {
				t.logger.Info(err)
			}
			r = hObj.Sum(nil)
		}

		ti.Stop()
		t.logger.Infof("- %v:%x", n, r)
		t.logger.Infof("result Adler32: %v", ti.Result())
	}

	{
		t.logger.Info("start CRC32")
		var (
			n   int
			r   []byte
			err error
		)
		ti.Start()

		for i := 0; i < 10240; i++ {
			hObj := crc32.New(crc32.IEEETable)
			n, err = hObj.Write(testload)
			if err != nil {
				t.logger.Info(err)
			}
			r = hObj.Sum(nil)
		}

		ti.Stop()
		t.logger.Infof("- %v:%x", n, r)
		t.logger.Infof("result CRC32: %v", ti.Result())
	}

	{
		t.logger.Info("start CRC64")
		var (
			n   int
			r   []byte
			err error
		)
		ti.Start()

		for i := 0; i < 10240; i++ {
			hObj := crc64.New(crc64.MakeTable(crc64.ISO))
			n, err = hObj.Write(testload)
			if err != nil {
				t.logger.Info(err)
			}
			r = hObj.Sum(nil)
		}

		ti.Stop()
		t.logger.Infof("- %v:%x", n, r)
		t.logger.Infof("result CRC64: %v", ti.Result())
	}
}
