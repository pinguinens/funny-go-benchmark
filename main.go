package main

import (
	"github.com/pinguinens/funny-go-benchmark/internal/config"
	"github.com/pinguinens/funny-go-benchmark/internal/runner"
	"github.com/pinguinens/funny-go-benchmark/internal/tester/hash"
	"github.com/pinguinens/funny-go-benchmark/pkg/log"
)

func main() {
	appCfg := config.New()
	logger := log.Logger{}
	logger.Info("== Funny Go Benchmark v.0.2 ==")

	rn := runner.New(appCfg)
	//test := memory.New(&logger, rn.Buffer(), appCfg.System.CoreCount)
	test := hash.New(&logger, rn.Buffer(), appCfg.System.CoreCount)
	rn.Run(test)
}
